package wallet_test

import (
	"math"
	"testing"

	wallet "github.com/AlkorMizar/Wallet"
)

// structure for input variables
type input struct {
	deopsitInp wallet.Bitcoin
	withdraw   wallet.Bitcoin
}

// structure for expected output
type output struct {
	err error
	res wallet.Bitcoin
}

// test for Deposite func
func TestDeposit(t *testing.T) {
	tests := map[string]struct {
		input wallet.Bitcoin
		want  output
	}{
		"simple": {input: 10, want: output{nil, 10}},
		"zero":   {input: 0, want: output{nil, 0}},
		"float":  {input: 12.15, want: output{nil, 12.15}},

		"negative": {input: -1.56, want: output{wallet.ErrNegDeposite, 0}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var w wallet.Wallet

			got := w.Deposit(tc.input)
			if got != tc.want.err {
				t.Fatalf("Erorr expected in test %s. Expected:\n%s\nGot:\n%s", name, tc.want.err, got)
			}

			if math.Abs(float64(w.Balance()-tc.want.res)) > 0.00000000001 {
				t.Fatalf("Erorr expected in test %s. Expected:\n%f\nGot:\n%f", name, tc.want.res, w.Balance())
			} else {
				t.Logf("Test %s went right", name)
			}
		})
	}
}

// test for Withdraw func
func TestWithdraw(t *testing.T) {
	tests := map[string]struct {
		input input
		want  output
	}{
		"simple":            {input: input{20, 10}, want: output{nil, 10}},
		"zero result":       {input: input{100, 100}, want: output{nil, 0}},
		"float":             {input: input{123.567, 8.9}, want: output{nil, 114.667}},
		"float zero result": {input: input{5.6789, 5.6789}, want: output{nil, 0}},
		"zero withdraw":     {input: input{12.3, 0}, want: output{nil, 12.3}},
		"big precision":     {input: input{1234.123456789, 12.815647935111}, want: output{nil, 1221.307808853889}},

		"negtive withdraw":               {input: input{11.3, -16.5}, want: output{wallet.ErrNegWithdraw, 11.3}},
		"float withdraw more than exist": {input: input{1.23567, 8.9}, want: output{wallet.ErrNotEnoughOnBalance, 1.23567}},
		"withdraw more than exist":       {input: input{10, 20}, want: output{wallet.ErrNotEnoughOnBalance, 10}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var w wallet.Wallet

			got := w.Deposit(tc.input.deopsitInp)
			if got != nil {
				t.Fatalf("Erorr expected in test %s. Deposite error:\n%s", name, got)
			}

			got = w.Withdraw(tc.input.withdraw)

			if got != tc.want.err {
				t.Fatalf("Erorr expected in test %s. Expected:\n%s\nGot:\n%s", name, tc.want.err, got)
			}

			if math.Abs(float64(w.Balance()-tc.want.res)) > 0.00000000001 {
				t.Fatalf("Erorr expected in test %s. Expected:\n%.15f\nGot:\n%.15f", name, tc.want.res, w.Balance())
			} else {
				t.Logf("Test %s went right", name)
			}
		})
	}
}
