package mysorts

// insertionSort implements Insertion Sort in-place
func insertionSort(in []int) {
	if len(in) <= 1 {
		return
	}

	// sort in-place
	for i := 1; i < len(in); i++ {

		// insert in[i] into sorted in[0] ~ in[i-1]
		// compare from i-1 to 0
		j := i - 1
		for ; j >= 0; j-- {
			if in[j+1] < in[j] { //Less()
				in[j], in[j+1] = in[j+1], in[j] //swap()
			} else {
				break
			}
		}
	}
}
