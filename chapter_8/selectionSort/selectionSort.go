////////////////////////////////////////////////////////////////////////////////
//	selectionSort.go  -  Sep-1-2022  -  aldebap
//
//	Implementation of a "selection sort" algorithm
////////////////////////////////////////////////////////////////////////////////

package selectionSort

func Sort(input []string) {
	for i := 0; i < len(input)-1; i++ {
		small := i

		for j := i + 1; j < len(input); j++ {
			if input[j] < input[small] {
				small = j
			}
		}

		//	if an element smaller than input[i] was found, swap them
		if small != i {
			aux := input[i]
			input[i] = input[small]
			input[small] = aux
		}
	}
}
