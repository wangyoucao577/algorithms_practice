package mysorts

import "math/rand"

// QuickSort implements Quick Sort in-place
func QuickSort(in myInterface) {

	quickSortImpl(in, 0, in.Len()-1)
}

func quickSortImpl(in myInterface, p, r int) {
	if p < r {
		//q := partition(in, p, r)
		q := randomizedPartition(in, p, r)
		quickSortImpl(in, p, q-1)
		quickSortImpl(in, q+1, r)
	}
}

func partition(in myInterface, p, r int) int {

	x := r // pick up last element as the pivot element
	i := p - 1

	for j := p; j <= r-1; j++ {
		if !in.Less(x, j) {
			i++

			if i != j {
				in.Swap(i, j)
			}
		}
	}

	if i+1 != x {
		in.Swap(i+1, x)
	}

	return i + 1
}

func randomizedPartition(in myInterface, p, r int) int {

	// pick up random x in [p,r]
	x := rand.Intn(r + 1 - p)
	x += p

	if x != r {
		in.Swap(x, r) // let x being the pivot element
	}

	return partition(in, p, r)
}
