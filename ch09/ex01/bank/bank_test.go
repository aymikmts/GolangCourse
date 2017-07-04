package bank

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	var withdraw int
	var want bool

	// balance:0  withdraw:0
	withdraw = 0
	want = true
	ok := Withdraw(withdraw)
	if ok != want {
		t.Errorf("balance:%d, withdraw:%d. ret is %v, but want %v.",
			Balance(), withdraw, ok, want)
	}

	// balance:0  withdraw:100
	withdraw = 100
	want = false
	ok = Withdraw(withdraw)
	if ok != want {
		t.Errorf("balance:%d, withdraw:%d. ret is %v, but want %v.",
			Balance(), withdraw, ok, want)
	}

	// balance:100  withdraw:0
	Deposit(100)
	withdraw = 0
	want = true
	ok = Withdraw(withdraw)
	if ok != want {
		t.Errorf("balance:%d, withdraw:%d. ret is %v, but want %v.",
			Balance(), withdraw, ok, want)
	}

	// balance:100  withdraw:10
	withdraw = 10
	want = true
	ok = Withdraw(withdraw)
	if ok != want {
		t.Errorf("balance:%d, withdraw:%d. ret is %v, but want %v.",
			Balance(), withdraw, ok, want)
	}

	// balance:100  withdraw:100
	Deposit(10)
	withdraw = 100
	want = true
	ok = Withdraw(withdraw)
	if ok != want {
		t.Errorf("balance:%d, withdraw:%d. ret is %v, but want %v.",
			Balance(), withdraw, ok, want)
	}

	// balance:100  withdraw:200
	Deposit(100)
	withdraw = 200
	want = false
	ok = Withdraw(withdraw)
	if ok != want {
		t.Errorf("balance:%d, withdraw:%d. ret is %v, but want %v.",
			Balance(), withdraw, ok, want)
	}
}
