package fibonaccibinaryexponentiation

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tables := []struct {
		n    int64
		want int64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{42, 267914296},
	}

	for _, table := range tables {
		fib := Fibonacci(table.n)
		if fib != table.want {
			t.Errorf("Fibonacci(%v): got: %v, want: %v", table.n, fib, table.want)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(int64(i))
	}
}
