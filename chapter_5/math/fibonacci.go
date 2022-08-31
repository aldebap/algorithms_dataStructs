////////////////////////////////////////////////////////////////////////////////
//	fibonacci.go  -  Ago-31-2022  -  aldebap
//
//	Implementation of a recursive Fibonacci function
////////////////////////////////////////////////////////////////////////////////

package math

func Fibonacci(num uint8) uint64 {
	if num == 0 {
		return 1
	}
	if num == 1 {
		return 1
	}

	return Fibonacci(num-1) + Fibonacci(num-2)
}
