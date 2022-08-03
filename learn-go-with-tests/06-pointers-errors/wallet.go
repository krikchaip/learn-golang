package pointers_errors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

// ?? implement `fmt.Stringer` interface on `Bitcoin`
// ?? this will be called when used with the `%s` format string
// ?? ref: https://pkg.go.dev/fmt#Stringer
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// ** private outside `pointers_errors` package
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of wallet in Deposit is %p \n", w)
	w.balance += amount // alternative: (*w).balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address of wallet in Balance is %p \n", w)
	return w.balance // alternative: (*w).balance
}

// ?? `error` is a built-in interface for handling errors
func (w *Wallet) Withdraw(amount Bitcoin) error {
	fmt.Printf("address of wallet in Withdraw is %p \n", w)

	if amount > w.balance {
		// ** like `new Error(msg)` in Javascript
		return ErrInsufficientFunds
	}

	w.balance -= amount // alternative: (*w).balance

	return nil
}
