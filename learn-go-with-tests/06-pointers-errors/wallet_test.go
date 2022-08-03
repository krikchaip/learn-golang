package pointers_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(10)

	fmt.Printf("address of wallet in test is %p \n", &wallet)

	// ** this will call got.String() and want.String() internally
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
