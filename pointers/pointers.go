package pointers

import (
	"errors"
	"fmt"
)

/* Variables declared with var are package scoped */
var ErrInsufficientFunds = errors.New("Insufficient Funds")

type Bitcoin int

/*
Bitcoin now implements the "Stringer"
interface which is defined in the fmt package,
and allows us to use the %s formatter in prints
*/

/*
Also notice that declaring a method
on a type declaration is possible just
like declaring a method on a struct
*/
func (b Bitcoin) String() string {
	return fmt.Sprintf("%dBTC", b)
}

type Stringer interface {
	String() string
}

type Wallet struct {
	value Bitcoin
}

func (w Wallet) Balance() Bitcoin {
	return w.value
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.value += amount
}

func (w *Wallet) Withdrawal(amount Bitcoin) error {
	if amount > w.Balance() {
		return ErrInsufficientFunds
	}
	w.value -= amount
	return nil
}
