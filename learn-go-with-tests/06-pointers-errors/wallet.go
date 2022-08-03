package pointers_errors

import "fmt"

type Bitcoin int

type Wallet struct {
	// private outside `pointers_errors` package
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
