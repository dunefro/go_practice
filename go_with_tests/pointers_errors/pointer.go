package bitcoin

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (w *Wallet) Deposit(money Bitcoin) {
	w.balance += money
}

func (w *Wallet) Withdrawl(money Bitcoin) error {
	if w.balance < money {
		return ErrInsufficientFunds
	}
	w.balance -= money
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC.", b)
}
