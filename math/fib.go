package math

import "math/big"

var memoize map[int]*big.Int

func init() {
	memoize = make(map[int]*big.Int)
}

// Fib prints the nth digit of the fibonacci sequence it will return 1 for anything < 0 as well ..
// it's calculated recursively and use big.Int since int64 will quickly overflow
func Fib(n int) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}

	// base case
	if n < 2 {
		memoize[n] = big.NewInt(1)
	}
	// Check if we stored it before if so return with no calculation
	if val, ok := memoize[n]; ok {
		return val
	}
	// initialize map then add previous 2 fib values
	memoize[n] = big.NewInt(0)
	memoize[n].Add(memoize[n], Fib(n-1))
	memoize[n].Add(memoize[n], Fib(n-2))

	return memoize[n]
}
