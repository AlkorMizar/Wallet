package wallet

import (
	"math"
	"testing"
)

//general test for accessing wallet balance
func TestBalance(t *testing.T) {

	var wallet Wallet
	wallet.balance = 10

	got := wallet.Balance()

	if math.Abs(float64(wallet.balance-got)) > 0.00000000000001 {
		t.Fatalf("Erorr expected in test. Expected:\n%f\nGot:\n%f", wallet.balance, got)
	} else {
		t.Log("Test went right")
	}
}

//test for Deposite func with correct input
func TestDepositCorrectInput(t *testing.T) {
	tests := map[string]struct {
		input Bitcoin
		want  Bitcoin
	}{
		"simple": {input: 10, want: 10},
		"zero":   {input: 0, want: 0},
		"float":  {input: 12.15, want: 12.15},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var wallet Wallet

			got := wallet.Deposit(tc.input)
			if got != nil {
				t.Fatalf("Erorr expected in test %s. Expected:\nnil\nGot:\n%s", name, got)
			} else if math.Abs(float64(wallet.Balance()-tc.want)) > 0.00000000001 {
				t.Fatalf("Erorr expected in test %s. Expected:\n%f\nGot:\n%f", name, tc.want, wallet.Balance())
			} else {
				t.Logf("Test %s went right", name)
			}
		})
	}
}

//test for Deposite func with incorrect input
func TestDepositIncorrectInput(t *testing.T) {
	var wallet Wallet

	got := wallet.Deposit(-1.256)
	if got == nil {
		t.Fatalf("Expcted error")
	} else {
		t.Log("Test went right")
	}
}

//structure for input variables
type input struct {
	deopsitInp Bitcoin
	withdraw   Bitcoin
}

//test for Withdraw func with correct input
func TestWithdrawCorrectInput(t *testing.T) {
	tests := map[string]struct {
		input input
		want  Bitcoin
	}{
		"simple":            {input: input{20, 10}, want: 10},
		"zero result":       {input: input{100, 100}, want: 0},
		"float":             {input: input{123.567, 8.9}, want: 114.667},
		"float zero result": {input: input{5.6789, 5.6789}, want: 0},
		"zero withdraw":     {input: input{12.3, 0}, want: 12.3},
		"big precision":     {input: input{1234.123456789, 12.815647935111}, want: 1221.307808853889},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var wallet Wallet

			got := wallet.Deposit(tc.input.deopsitInp)
			if got != nil {
				t.Fatalf("Erorr expected in test %s. Expected:\nnil\nGot:\n%s", name, got)
			}

			got = wallet.Withdraw(tc.input.withdraw)
			if got != nil {
				t.Fatalf("Erorr expected in test %s. Expected:\nnil\nGot:\n%s", name, got)
			} else if math.Abs(float64(wallet.Balance()-tc.want)) > 0.00000000001 {
				t.Fatalf("Erorr expected in test %s. Expected:\n%.15f\nGot:\n%.15f", name, tc.want, wallet.balance)
			} else {
				t.Logf("Test %s went right", name)
			}
		})
	}
}

//test for Withdraw func with incorrect input
func TestWithdrawIncorrectInput(t *testing.T) {
	tests := map[string]struct {
		input input
	}{
		"withdraw more than exist":       {input: input{10, 20}},
		"float withdraw more than exist": {input: input{1.23567, 8.9}},
		"negtive withdraw":               {input: input{11.3, -16.5}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var wallet Wallet

			got := wallet.Deposit(tc.input.deopsitInp)
			if got != nil {
				t.Fatalf("Erorr expected in test %s. Expected:\nnil\nGot:\n%s", name, got)
			}

			got = wallet.Withdraw(tc.input.withdraw)
			if got == nil {
				t.Fatalf("Erorr expected in test %s. Expected:\nerror\nGot:nill", name)
			} else {
				t.Log("Test went right")
			}
		})
	}
}
