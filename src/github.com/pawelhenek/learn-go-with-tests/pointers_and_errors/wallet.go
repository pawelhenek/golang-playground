package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (wallet *Wallet) Withdraw(amount Bitcoin) (Bitcoin, error) {

	if amount > wallet.balance {
		return Bitcoin(-1), ErrInsufficientFunds
	}

	wallet.balance -= amount
	return Bitcoin(wallet.balance), nil
}

func main() {
	var inPocketAtStart = Bitcoin(10)
	var wallet = Wallet{inPocketAtStart}

	var toWithdrawn = Bitcoin(6)
	var afterWithdrawn, _ = wallet.Withdraw(toWithdrawn)

	fmt.Println(fmt.Sprintf("So we had %d and %d was withdrawn so we got %d now", inPocketAtStart, toWithdrawn, afterWithdrawn))

}
