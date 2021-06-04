package bitcoin

import "testing"

type Wallet struct {
	balance int
}

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	want := 10
	if got != want {
		t.Errorf("Expected: %d Recieved: %d", got, want)
	}
}
