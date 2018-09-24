package mysorts

// insertionSort implements Merge Sort In-place
func mergeSort(in myInterface) {

	mergeSortImpl(in, 0, in.Len()-1)
}

func mergeSortImpl(in myInterface, p, r int) {

	if p < r {
		q := (p + r) / 2
		mergeSortImpl(in, p, q)
		mergeSortImpl(in, q+1, r)
		merge(in, p, q, r)
	}
}

//merge sorted sub-array in[p~q] and in[q+1~r] in-place, 0 <= p <= q < r= < in.Len()-1
func merge(in myInterface, p, q, r int) {

	i := p
	j := q + 1
	for k := p; k < r; k++ {
		if j > r {
			break
		}

		if !in.Less(i, j) {
			for m := j; m > i; m-- {
				in.Swap(m, m-1)
			}
			j++
		}
		i++

	}
}
