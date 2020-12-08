package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit and check balance", func(t *testing.T) {
		//declare wallet of type Wallet struct
		wallet := Wallet{}

		//call wallet.Deposit method
		wallet.Deposit(10)

		validateBalance(t, wallet, 10)
	})
	t.Run("Withdraw and check balance", func(t *testing.T) {
		//declare Wallet struct with an initial balance
		wallet := Wallet{balance: 100}

		wallet.Withdraw(10)

		validateBalance(t, wallet, 90)
	})
}

func validateBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("Got %d, Want %d", got, want)
	}
}
