package eval

import (
	"fmt"
	"math"
)

// Exprは算術式
type Expr interface {
	// Evalは、環境env内でこのExprの値を返します。
	Eval(env Env) float64

	// Checkは、このExpr内のエラーを報告し、セットにそのVarを追加します。
	Check(vars map[Var]bool) error
}

// Varは変数を特定します。
type Var string

// literalは数値定数。
type literal float64

// unaryは単項演算子を表します。
type unary struct {
	op rune // '+' or '-'
	x  Expr
}

// binaryは二項演算子を表します。
type binary struct {
	op   rune // '+', '-', '*', '/'
	x, y Expr
}

// callは関数呼び出し式を表します。
type call struct {
	fn   string // "pow", "sin", "sqrt"
	args []Expr
}

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
