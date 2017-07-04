package bank

var deposits = make(chan int)     // 入金額を送信する
var balances = make(chan int)     // 残高を受信する
var withdraw = make(chan int)     // 出金額を送信する
var withDrawRet = make(chan bool) // 出金可否を送受信する

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

// Withdrawは、指定された金額を引き出し、
// 取引が成功したか、残高不足で失敗したかを返します。
func Withdraw(amount int) bool {
	withdraw <- amount
	return <-withDrawRet
}

func teller() {
	var balance int // balanceはtellerゴルーチンに閉じ込められている
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraw:
			if balance < amount {
				withDrawRet <- false
			} else {
				balance -= amount
				withDrawRet <- true
			}
		}
	}
}

func init() {
	go teller() // モニターゴルーチンを開始する
}
