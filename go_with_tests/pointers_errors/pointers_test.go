package bitcoin

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("Expected: %s Recieved: %s", got, want)
		}
	}
	assertError := func(t testing.TB, errGotString error, errWantString error) {
		t.Helper()
		if errGotString == nil {
			t.Fatal("Wanted an error but didn't get one")
		}
		if errGotString.Error() != errWantString.Error() {
			t.Errorf("Expected Error: %q, Recieved Error: %q", errWantString, errWantString)
		}
	}
	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("Wasn't expecting any error here but got %q", err.Error())
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
		err := wallet.Withdrawl(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})
	t.Run("Withdrawl more than amount", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdrawl(Bitcoin(30))
		// Here we are passing the starting balance only
		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds)
	})
}
