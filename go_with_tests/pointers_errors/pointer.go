package bitcoin

type Bitcoin int

func (w *Wallet) Deposit(money Bitcoin) {
	w.balance += money
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
