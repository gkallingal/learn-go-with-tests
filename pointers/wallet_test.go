package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit and check balance", func(t *testing.T) {
		//declare wallet of type Wallet struct
		wallet := Wallet{}

		//call wallet.Deposit method
		wallet.Deposit(10)

		//declare got variable and apply value from wallet.Balance() method
		got := wallet.Balance()
		fmt.Printf("memory address of test balance: %p\n", &wallet.balance)

		want := Bitcoin(10)

		if got != want {
			t.Errorf("Got %s, Want %s", got, want)
		}
	})
	t.Run("Withdraw and check balance", func(t *testing.T) {
		//declare Wallet struct with an initial balance
		wallet := Wallet{balance: 100}

		wallet.Withdraw(10)

		got := wallet.Balance()

		want := Bitcoin(90)

		if got != want {
			t.Errorf("Got %d, Want %d", got, want)
		}
	})
}
