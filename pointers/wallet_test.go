package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	//declare wallet of type Wallet struct
	wallet := Wallet{}

	//call wallet.Deposit method
	wallet.Deposit(10)

	//declare got variable and apply value from wallet.Balance() method
	got := wallet.Balance()
	fmt.Printf("memory address of test balance: %p\n", &wallet.balance)

	want := 10

	if got != want {
		t.Errorf("Got %d, Want %d", got, want)
	}
}
