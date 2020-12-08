package pointers

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("memory address of deposit balance: %p\n", &w.balance)
	w.balance += amount

}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
