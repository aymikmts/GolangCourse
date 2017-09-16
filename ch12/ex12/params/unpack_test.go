package params

import "testing"

func TestCheckValue(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		isValid bool
	}{
		// 電子メールアドレス
		{"em", "test.0123456@test.ne.jp", true},
		{"em", "@test.ne.jp", false},
		{"em", "_test3@test.ne.jp", false},
		{"em", ".test@test.ne.jp", false},

		// クレジットカード番号
		{"cn", "12345678901234", true},
		{"cn", "123456789012345", true},
		{"cn", "1234567890123456", true},
		{"cn", "12345", false},
		{"cn", "12345", false},

		// 郵便コード
		{"pn", "1234567", true},
		{"pn", "123456", false},
	}

	for _, test := range tests {
		isValid := true
		err := checkValue(test.name, test.value)
		if err != nil {
			t.Log(err)
			isValid = false
		}

		if isValid != test.isValid {
			t.Errorf("checkValue(%#v, %#v) got isValid:%v, want isValid:%v",
				test.name, test.value, isValid, test.isValid)
		}
	}
}
