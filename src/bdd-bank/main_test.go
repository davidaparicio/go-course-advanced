package main

import (
	"testing"

	"github.com/davidaparicio/go-course-advanced/src/bdd-godog/bank"
)

//dummy test that executes main()
//https://stackoverflow.com/a/51277293
/*func TestRunMain(t *testing.T) {
	main()
}*/

func TestDeposit(t *testing.T) {
	testcases := []struct {
		start, deposit, want int
	}{
		{10, 0, 10},
		{100, 50, 150},
		{-100, 50, -50},
	}
	for _, tc := range testcases {
		acc := bank.NewAccount("Test")
		acc.Deposit(tc.start)
		acc.Deposit(tc.deposit)
		res := acc.Balance()
		if res != tc.want {
			t.Errorf("TestDeposit(start: %d, deposit: %d): got %d, want %d", tc.start, tc.deposit, res, tc.want)
		}
	}
}

func FuzzDeposit(f *testing.F) {
	acc := bank.NewAccount("Fuzz Buzz")
	acc.Deposit(844383084)                                                       //Random initialization
	testcases := []int{10, 100, -100, -9223372036854775808, 9223372036854775807} //https://gosamples.dev/int-min-max/
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, deposit int) {
		start := acc.Balance()
		acc.Deposit(deposit)
		end := acc.Balance()
		//Log added to understand better
		//t.Logf("Account: start=%d, deposit=%d, end=%d", start, deposit, end)
		if end != start+deposit {
			t.Errorf("Deposit produced invalid result %d != %d + %d", end, start, deposit)
		}
	})
}

func benchmarkDeposit(b *testing.B, deposits []int) {
	acc := bank.NewAccount("Bench Mark")
	for i := 0; i < b.N; i++ {
		for _, deposit := range deposits {
			acc.Deposit(deposit)
		}
	}
}

func BenchmarkDeposit_Small(b *testing.B) {
	d := []int{0, 10, -10}
	benchmarkDeposit(b, d)
}

func BenchmarkReverse_Big(b *testing.B) {
	d := []int{132283115, -596768679, 224391349, -316524437, 217443011, 435025920, 357738921, 91156801, -448356816, 844383084}
	// Timestamp: 2022-07-13 14:52:17 UTC
	// https://www.random.org/integers/?num=10&min=-1000000000&max=1000000000&col=5&base=10&format=html&rnd=new
	benchmarkDeposit(b, d)
}
