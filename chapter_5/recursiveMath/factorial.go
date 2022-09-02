////////////////////////////////////////////////////////////////////////////////
//	factorial.go  -  Ago-31-2022  -  aldebap
//
//	Implementation of a recursive factorial function
////////////////////////////////////////////////////////////////////////////////

package recursiveMath

import "fmt"

func Factorial(num uint8) uint64 {
	fmt.Printf("[debug] %d!\n", num)

	if num == 0 {
		return 1
	}

	return uint64(num) * Factorial(num-1)
}
