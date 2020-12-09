package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit and Check Balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		fmt.Printf("memory address of balance from Test: %p\n", &wallet.balance)
		assertBalance(t, wallet, 10)

	})

	t.Run("Withdraw and Check Balance", func(t *testing.T) {
		wallet := Wallet{balance: 90}

		err := wallet.Withdraw(10)
		assertBalance(t, wallet, 80)
		assertNoError(t, err)
	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(20)
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Got %s, Want %s", got, want)
	}
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("Wanted error but got none")
	}
	if err.Error() != want.Error() {
		t.Errorf("Got %s, Want %s", err.Error(), want)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Error("Got error expected none")
	}
}
