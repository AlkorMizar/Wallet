package wallet

import (
	"errors"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

// Deposite on this wallet.
//bit >= 0
func (w *Wallet) Deposit(bit Bitcoin) (e error) {
	if bit < 0 {
		return errors.New("can't process negative deposite")
	}
	w.balance += bit
	return nil
}

// Deposite on this wallet.
//bit >= 0 and can't be bigger than wallets balance
func (w *Wallet) Withdraw(bit Bitcoin) (e error) {
	if bit < 0 {
		return errors.New("can't process negative withdraw")
	}

	if w.balance-bit < 0 {
		return errors.New("balance less then withdraw amount")
	}

	w.balance -= bit
	return nil
}

//Get balance of this wallet
func (w Wallet) Balance() (res Bitcoin) {
	return w.balance
}
