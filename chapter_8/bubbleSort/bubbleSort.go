////////////////////////////////////////////////////////////////////////////////
//	bubbleSort.go  -  Sep-2-2022  -  aldebap
//
//	Implementation of a "bubble sort" algorithm
////////////////////////////////////////////////////////////////////////////////

package bubbleSort

func Sort(input []string) {
	for i := len(input) - 1; i >= 0; i-- {

		for j := 1; j <= i; j++ {

			//	if an element smaller than input[i] was found, swap them
			if input[j-1] > input[j] {
				aux := input[j-1]
				input[j-1] = input[j]
				input[j] = aux
			}
		}
	}
}
