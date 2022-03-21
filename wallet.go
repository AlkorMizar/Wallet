package wallet

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(bit Bitcoin) (e error) {
	return nil
}

func (w *Wallet) Withdraw(bit Bitcoin) (e error) {
	return nil
}

func (w Wallet) Balance() (res Bitcoin) {
	return -1
}
