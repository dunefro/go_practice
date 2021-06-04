package bitcoin

import "fmt"

type Bitcoin int

func (w *Wallet) Deposit(money Bitcoin) {
	w.balance += money
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC.", b)
}
