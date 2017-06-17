package eval

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		want string
	}{
		{"x", "x"},
		{"+x", "(+x)"},
		{"x+y", "(x + y)"},
		{"pow(x)", "pow(x)"},
		{"sin(x+y)", "sin((x + y))"},
		{"pow(x,3)+pow(y,3)", "(pow(x, 3) + pow(y, 3))"},
		{"5/9*(F-32)", "((5 / 9) * (F - 32))"},
	}
	for _, test := range tests {
		// 入力をパース
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}

		// 構文ツリーをチェック
		got := fmt.Sprintf("%v", expr)
		if got != test.want {
			t.Errorf("eval:%q, want %q\n",
				got, test.want)
		}

		// 出力したツリーを再度Parse
		expr2, err := Parse(got)
		if err != nil {
			t.Error(err)
			continue
		}
		got2 := fmt.Sprintf("%v", expr2)
		if got2 != got {
			t.Errorf("result of re-Parse: %v, it is not equal: %v.\n", got2, got)
		}
	}
}
