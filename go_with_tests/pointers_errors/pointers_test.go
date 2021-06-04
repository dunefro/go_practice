package bitcoin

import "testing"

type Wallet struct {
	balance Bitcoin
}

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	want := Bitcoin(10)
	if got != want {
		t.Errorf("Expected: %s Recieved: %s", got, want)
	}
}
