package params

import "testing"
import "reflect"

func TestCheckValue(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		isErr bool
	}{
		// 電子メールアドレス
		{"em", "test@test.ne.jp", true},
		{"em", []string{"test1@test.ne.jp", "test2@test.ne.jp"}, true},
		{"em", "test.@test.ne.jp", false},
		{"em", "test..123@test.ne.jp", false},
		{"em", ".test@test.ne.jp", false},
		{"em", []string{"test.@test.ne.jp", "test.@test.ne.jp"}, false},

		// クレジットカード番号
		// テスト用番号: https://www.webcreatorbox.com/tech/creditcard-test-numbers
		{"cn", 4012888888881881, true},
		{"cn", []int{4012888888881881, 5105105105105100}, true},
		{"cn", 12345, false},
		{"cn", 12345, false},
		{"cn", []int{5105105105105100, 12345}, false},

		// 郵便コード
		{"pn", 1234567, true},
		{"pn", []int{1234567, 7654321}, true},
		{"pn", 123456, false},
		{"pn", []int{1234567, 123456}, false},

		// tagがない場合
		{"other", "test", true},
	}

	for _, test := range tests {
		isErr := false
		err := checkValue(test.name, reflect.ValueOf(test.value))
		if err != nil {
			isErr = true
		}

		if isErr != test.isErr {
			t.Errorf("checkValue(%#v, %#v) got isErr:%v, want isErr:%v",
				test.name, test.value, isErr, test.isErr)
		}
	}
}
