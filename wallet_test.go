package wallet_test

import (
	"errors"
	"sync"
	"testing"

	wallet "github.com/AlkorMizar/Wallet"
)

// test for Deposite func
func TestDeposit(t *testing.T) {
	tests := map[string]struct {
		input   wallet.Bitcoin
		wantErr error
		wantVal wallet.Bitcoin
	}{
		"simple": {
			input:   10,
			wantErr: nil,
			wantVal: 10,
		},
		"zero": {
			input:   0,
			wantErr: nil,
			wantVal: 0,
		},
		"float": {
			input:   12.15,
			wantErr: nil,
			wantVal: 12.15,
		},
		"negative": {
			input:   -1.56,
			wantErr: wallet.ErrIncorrectInput,
			wantVal: 0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			w := wallet.InitWallet()

			got := w.Deposit(tc.input)
			if !errors.Is(got, tc.wantErr) {
				t.Fatalf("erorr in test %s, expected: %v, got: %v", name, tc.wantErr, got)
			}

			if w.Balance() != tc.wantVal {
				t.Fatalf("erorr in test %s, expected: %.15f ,got: %.15f", name, tc.wantVal, w.Balance())
			}
		})
	}
}

// test for Withdraw func
func TestWithdraw(t *testing.T) {
	tests := map[string]struct {
		inputDepos  wallet.Bitcoin
		inputWithdr wallet.Bitcoin
		wantErr     error
		wantVal     wallet.Bitcoin
	}{
		"simple": {
			inputDepos:  20,
			inputWithdr: 10,
			wantErr:     nil,
			wantVal:     10},
		"zero result": {
			inputDepos:  100,
			inputWithdr: 100,
			wantErr:     nil,
			wantVal:     0,
		},
		"float": {
			inputDepos:  123.56,
			inputWithdr: 8.9,
			wantErr:     nil,
			wantVal:     114.66,
		},
		"float zero result": {
			inputDepos:  5.6789,
			inputWithdr: 5.6789,
			wantErr:     nil,
			wantVal:     0,
		},
		"zero withdraw": {
			inputDepos:  12.3,
			inputWithdr: 0,
			wantErr:     nil,
			wantVal:     12.3,
		},
		"big precision": {
			inputDepos:  1234.123456789,
			inputWithdr: 12.815647935111,
			wantErr:     nil,
			wantVal:     1221.307808853889128,
		},

		"negtive withdraw": {
			inputDepos:  11.3,
			inputWithdr: -16.5,
			wantErr:     wallet.ErrIncorrectInput,
			wantVal:     11.3,
		},
		"float withdraw more than exist": {
			inputDepos:  1.23567,
			inputWithdr: 8.9,
			wantErr:     wallet.ErrNotEnoughOnBalance,
			wantVal:     1.23567,
		},
		"withdraw more than exist": {
			inputDepos:  10,
			inputWithdr: 20,
			wantErr:     wallet.ErrNotEnoughOnBalance,
			wantVal:     10,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			w := wallet.InitWallet()

			got := w.Deposit(tc.inputDepos)
			if got != nil {
				t.Fatalf("error during deposite : name: %s, error: %v", name, got)
			}

			got = w.Withdraw(tc.inputWithdr)

			if !errors.Is(got, tc.wantErr) {
				t.Fatalf("erorr in test %s, expected: %v, got: %v", name, tc.wantErr, got)
			}

			if w.Balance() != tc.wantVal {
				t.Fatalf("erorr in test %s want: %.15f, got: %.15f", name, tc.wantVal, w.Balance())
			}
		})
	}
}

func TestCocurrency(t *testing.T) {
	iterations := 10000
	witdr := wallet.Bitcoin(5.6478)
	want := wallet.Bitcoin(10)

	w := wallet.InitWallet()
	_ = w.Deposit(witdr*wallet.Bitcoin(iterations) + want)

	var wg sync.WaitGroup

	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		go func() {
			defer wg.Done()

			_ = w.Withdraw(witdr)
		}()
	}

	wg.Wait()

	if w.Balance() != want {
		t.Fatalf("erorr in test want: %f, got: %f", want, w.Balance())
	}
}
