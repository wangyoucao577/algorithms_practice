package mysorts

// insertionSort implements Insertion Sort
func insertionSort(in []int) []int {
	if len(in) <= 1 {
		return in
	}

	// sort in-place
	for i := 1; i < len(in); i++ {
		curr := in[i]

		// insert in[i] into sorted in[0] ~ in[i-1]
		// compare from i-1 to 0
		j := i - 1
		for ; j >= 0; j-- {
			if curr < in[j] { //Less()
				in[j+1] = in[j] //Swap()
			} else {
				break
			}
		}
		in[j+1] = curr //Swap()
	}

	return in
}
