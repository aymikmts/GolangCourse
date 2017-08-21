package eval

// Expr は算術式
type Expr interface {
	// Evalは、環境env内でこのExprの値を返します。
	Eval(env Env) float64

	// Checkは、このExpr内のエラーを報告し、セットにそのVarを追加します。
	Check(vars map[Var]bool) error

	// Stringは、構文ツリーをきれいに表示します。
	String() string

	// Minは、オペランドの最小値を計算します。
	Min(env Env) float64
}

// Var は変数を特定します。例: x
type Var string

// literal は数値定数。例: 3.141
type literal float64

// unary は単項演算式を表します。例: -x
type unary struct {
	op rune // '+' or '-'
	x  Expr
}

// binary は二項演算式を表します。例: x+y
type binary struct {
	op   rune // '+', '-', '*', '/'のどれか
	x, y Expr
}

// call は関数呼び出し式を表します。例: sin(x)
type call struct {
	fn   string // "pow", "sin", "sqrt"のどれか
	args []Expr
}

// Env は変数名を値へと対応付けるmapです。
type Env map[Var]float64
