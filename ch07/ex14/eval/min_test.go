package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestMin(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want float64
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, math.Pi},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, 1},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, 3},
		{"5 / 9 * (F - 32)", Env{"F": -40}, -72},
		{"5 / 9 * (F - 32)", Env{"F": 32}, 0},
		{"5 / 9 * (F - 32)", Env{"F": 212}, 0.5555555555555556},
	}
	var prevExpr string
	for _, test := range tests {
		// 変更されているときだけexprを表示する。
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // パースエラー
			continue
		}
		got := expr.Min(test.env)
		fmt.Printf("\t%v => %.6g\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Min() in %v = %.6g, want %.6g\n",
				test.expr, test.env, got, test.want)
		}
	}
}
