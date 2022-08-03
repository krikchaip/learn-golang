package pointers_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		fmt.Printf("address of wallet in t.deposit is %p \n", &wallet)

		assertBalance(t, wallet.Balance(), Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		fmt.Printf("address of wallet in t.withdraw is %p \n", &wallet)

		assertNoError(t, err)
		assertBalance(t, wallet.Balance(), Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)

		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet.Balance(), startingBalance)
	})
}

func assertBalance(t testing.TB, got, want Bitcoin) {
	t.Helper()

	// ** this will call got.String() and want.String() internally
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()

	if err == nil {
		// ** stop the current test and exit process
		t.Fatal("wanted an error but didn't get one")
		return
	}

	// ?? `error` is comparable 🤔
	if err != want {
		t.Errorf("got %q want %q", err, want)
	}
}
