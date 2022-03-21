package wallet

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

// Deposite on this wallet.
//bit >= 0
func (w *Wallet) Deposit(bit Bitcoin) (e error) {
	return nil
}

// Deposite on this wallet.
//bit >= 0
func (w *Wallet) Withdraw(bit Bitcoin) (e error) {
	return nil
}

//Get balance of this wallet
func (w Wallet) Balance() (res Bitcoin) {
	return -1
}
