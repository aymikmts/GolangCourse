package eval

// Exprは算術式
type Expr interface {
	// Evalは、環境env内でこのExprの値を返します。
	Eval(env Env) float64

	// Checkは、このExpr内のエラーを報告し、セットにそのVarを追加します。
	Check(vars map[Var]bool) error

	// Ex13で追加
	// 構文ツリーを表示する
	String() string
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
