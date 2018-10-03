package mysorts

// InsertionSort implements Insertion Sort in-place
func InsertionSort(in myInterface) {
	if in.Len() <= 1 {
		return
	}

	// sort in-place
	for i := 1; i < in.Len(); i++ {

		// insert in[i] into sorted in[0] ~ in[i-1]
		// compare from i-1 to 0
		j := i - 1
		for ; j >= 0; j-- {
			if in.Less(j+1, j) { //Less()
				in.Swap(j, j+1) //Swap()
			} else {
				break
			}
		}
	}
}
