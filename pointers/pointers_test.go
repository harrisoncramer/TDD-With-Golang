package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdrawal", func(t *testing.T) {
		wallet := Wallet{value: 10}

		err := wallet.Withdrawal(9)
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(1))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{value: startingBalance}

		err := wallet.Withdrawal(Bitcoin(11))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

/* Function definitions are hoisted in Golang */
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if wallet.value != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	/* We use Fatal here to stop the test from continuing */
	if got == nil {
		t.Fatal("wanted err but didn't get one")
	}

	if got.Error() != want.Error() {
		t.Errorf("wanted error msg %q, got error msg %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("err with message %s", err.Error())
	}
}
