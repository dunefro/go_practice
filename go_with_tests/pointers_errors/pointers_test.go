package bitcoin

import "testing"

type Wallet struct {
	balance Bitcoin
}

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("Expected: %s Recieved: %s", got, want)
		}
	}
	t.Run("Wallet Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(30))
		assertBalance(t, wallet, Bitcoin(30))
	})
	t.Run("Wallet Withdrawl", func(t *testing.T) {
		// balance must be added
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdrawl(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
}
