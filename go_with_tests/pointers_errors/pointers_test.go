package bitcoin

import "testing"

type Wallet struct {
	balance Bitcoin
}

func TestWallet(t *testing.T) {
	t.Run("Wallet Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(30))
		got := wallet.Balance()
		want := Bitcoin(30)
		if got != want {
			t.Errorf("Expected: %s Recieved: %s", got, want)
		}
	})
	t.Run("Wallet Withdrawl", func(t *testing.T) {
		// balance must be added
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdrawl(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("Expected: %s Recieved: %s", got, want)
		}
	})
}
