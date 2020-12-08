package pointers

import "testing"

func TestWallet(t *testing.T) {
	//declare wallet of type Wallet struct
	wallet := Wallet{}

	//call wallet.Deposit method
	wallet.Deposit(10)

	//declare got variable and apply value from wallet.Balance() method
	got := wallet.Balance()

	want := 10

	if got != want {
		t.Errorf("Got %d, Want %d", got, want)
	}
}
