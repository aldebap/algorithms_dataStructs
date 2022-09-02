////////////////////////////////////////////////////////////////////////////////
//	insertionSort.go  -  Sep-2-2022  -  aldebap
//
//	Implementation of a "insertion sort" algorithm
////////////////////////////////////////////////////////////////////////////////

package insertionSort

func Sort(input []string) {
	for i := 1; i < len(input); i++ {
		aux := input[i]
		j := i

		//	if input[i] if smaller than elements before it, swap them
		for {
			if j == 0 || input[j-1] <= aux {
				break
			}
			input[j] = input[j-1]
			j--
		}
		input[j] = aux
	}
}
