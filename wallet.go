package wallet

import (
	"errors"
)

var (
	ErrIncorrectInput     = errors.New("incorrect input")
	ErrNotEnoughOnBalance = errors.New("balance less then withdraw amount")
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

// Balance get balance of this wallet
func (w *Wallet) Balance() (res Bitcoin) {
	return w.balance
}

// Deposite on this wallet.
// bit >= 0
func (w *Wallet) Deposit(bit Bitcoin) error {
	if bit < 0 {
		return ErrIncorrectInput
	}

	w.balance += bit

	return nil
}

// Withdraw from this wallet.
// bit >= 0 and can't be bigger than wallets balance
func (w *Wallet) Withdraw(bit Bitcoin) error {
	if bit < 0 {
		return ErrIncorrectInput
	}

	if w.balance-bit < 0 {
		return ErrNotEnoughOnBalance
	}

	w.balance -= bit

	return nil
}

func InitWallet() *Wallet {
	return &Wallet{}
}
