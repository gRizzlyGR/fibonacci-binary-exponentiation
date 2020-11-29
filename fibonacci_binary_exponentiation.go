package fibonaccibinaryexponentiation

import (
	"math/big"
)

// | 1 1 |
// | 1 0 |
var fibMatrix = [][]big.Int{{*big.NewInt(1), *big.NewInt(1)}, {*big.NewInt(1), *big.NewInt(0)}}

// | 1 0 |
// | 0 1 |
var identity = [][]big.Int{{*big.NewInt(1), *big.NewInt(0)}, {*big.NewInt(0), *big.NewInt(1)}}

// Fibonacci returns the fibonacci number at the n-th position
func Fibonacci(n int64) int64 {
	bigMat := pow(fibMatrix, big.NewInt(int64(n)))

	// | 1 1 |^n == | f(n+1)  f(n)  |  <-- we take this
	// | 1 0 |      |  f(n)  f(n-1) |
	bigN := bigMat[0][1]

	return bigN.Int64()
}

// pow computes the product using the binary exponentation, where we can compute
// any product in the form of a^n, applying a certain logic to the exponent n:
// we recursively divide n by 2, taking one half each time and ignoring the
// other one, since they are equal, so we can save resources. We must check if n
// is even or odd: 
// - If even, we can multiply the partial result for itself
// - If odd, we must multiply the result for itself, and then with the base number
// This approach works for numbers, but also for matrixes, in particular we can
// consider the matrix rapresentation of Fibonacci where:
// | 1 1 |^n == | f(n+1)  f(n)  |
// | 1 0 |      |  f(n)  f(n-1) |
// with f(x) returning the Fibonacci number for a certain value x
// We can compute Fibonacci faster than the classical approach.
// Complexity is O(log n)
func pow(mat [][]big.Int, n *big.Int) [][]big.Int {
	// If exponent is 0, return indentity matrix
	if n.Cmp(big.NewInt(0)) == 0 {
		return identity
	}

	// If exponenty is 1, return itself
	if n.Cmp(big.NewInt(1)) == 0 {
		return mat
	}

	// Recursively compute the power matrix dividing n by 2
	r := pow(fibMatrix, new(big.Int).Div(n, big.NewInt(2)))

	// If n is even, multiply the resulting matrix for itself:
	mod := new(big.Int).Mod(n, big.NewInt(2))
	if mod.Cmp(big.NewInt(0)) == 0 {
		return matrixMultiply(r, r)
	}

	// If n is odd, we must consider a spurious matrix that we must multiply
	tmp := matrixMultiply(fibMatrix, r)
	return matrixMultiply(tmp, r)
}

// matrixMultiply computes the row-per-column product of a 2x2 matrix
func matrixMultiply(a [][]big.Int, b [][]big.Int) [][]big.Int {
	c := make([][]big.Int, 2)

	for i := 0; i < 2; i++ {
		c[i] = make([]big.Int, 2)
	}

	//a[0][0] * b[0][0] + a[0][1] * b[1][0]
	c[0][0] = *new(big.Int).Add(
		new(big.Int).Mul(&a[0][0], &b[0][0]),
		new(big.Int).Mul(&a[0][1], &b[1][0]),
	)

	// a[0][0] * b[0][1] + a[0][1] * b[1][1]
	c[0][1] = *new(big.Int).Add(
		new(big.Int).Mul(&a[0][0], &b[0][1]),
		new(big.Int).Mul(&a[0][1], &b[1][1]),
	)

	// a[1][0] * b[0][0] + a[1][1] * b[1][0]
	c[1][0] = *new(big.Int).Add(
		new(big.Int).Mul(&a[1][0], &b[0][0]),
		new(big.Int).Mul(&a[1][1], &b[1][0]),
	)

	// a[1][0] * b[0][1] + a[1][1] * b[1][1]
	c[1][1] = *new(big.Int).Add(
		new(big.Int).Mul(&a[1][0], &b[0][1]),
		new(big.Int).Mul(&a[1][1], &b[1][1]),
	)

	return c
}
