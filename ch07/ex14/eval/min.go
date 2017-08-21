package eval

import (
	"fmt"
	"math"
)

func (v Var) Min(env Env) float64 {
	return env[v]
}

func (l literal) Min(env Env) float64 {
	return float64(l)
}

func (u unary) Min(env Env) float64 {
	var exp float64
	switch u.op {
	case '+':
		exp = +u.x.Eval(env)
	case '-':
		exp = -u.x.Eval(env)
	default:
		panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
	}
	return math.Min(u.x.Min(env), exp)
}

func (b binary) Min(env Env) float64 {
	arg := math.Min(b.x.Min(env), b.y.Min(env))
	var exp float64
	switch b.op {
	case '+':
		exp = b.x.Eval(env) + b.y.Eval(env)
	case '-':
		exp = b.x.Eval(env) - b.y.Eval(env)
	case '*':
		exp = b.x.Eval(env) * b.y.Eval(env)
	case '/':
		exp = b.x.Eval(env) / b.y.Eval(env)
	default:
		panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
	}
	return math.Min(arg, exp)
}

func (c call) Min(env Env) float64 {
	var arg float64
	var exp float64
	switch c.fn {
	case "pow":
		arg = math.Min(c.args[0].Min(env), c.args[1].Min(env))
		exp = math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		arg = c.args[0].Min(env)
		exp = math.Sin(c.args[0].Eval(env))
	case "sqrt":
		arg = c.args[0].Min(env)
		exp = math.Sqrt(c.args[0].Eval(env))
	default:
		panic(fmt.Sprintf("unsupported function call: %s", c.fn))
	}
	return math.Min(arg, exp)
}
