package mysorts

import (
	"math"
)

/************************* In-place implementation *************************/

// MergeSortInPlace implements Merge Sort In-place
func MergeSortInPlace(in myInterface) {

	mergeSortImpl(in, 0, in.Len()-1)
}

// mergeSortImpl split the array to two parts, then recurses call itself for them.
//  after that merge results of the two parts.
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

/************************* In-place implementation *************************/

/************************* With aux array implementation *************************/

// MergeSort implements Merge Sort with aux array
func MergeSort(in []int) {
	mergeSortImplAuxArrayBased(in, 0, len(in)-1)
}

// mergeSortImplAuxArrayBased split the array to two parts, then recurses call itself for them.
//  after that merge results of the two parts.
// Same process flow as mergeSortImpl
func mergeSortImplAuxArrayBased(in []int, p, r int) {

	if p < r {
		q := (p + r) / 2
		mergeSortImplAuxArrayBased(in, p, q)
		mergeSortImplAuxArrayBased(in, q+1, r)
		mergeAuxArrayBased(in, p, q, r)
	}
}

func mergeAuxArrayBased(in []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q // r - (q+1) + 1

	// create aux arrays
	auxArray1 := make([]int, n1+1, n1+1)
	auxArray2 := make([]int, n2+1, n2+1)
	for i := 0; i < n1; i++ {
		auxArray1[i] = in[p+i]
	}
	auxArray1[n1] = math.MaxInt32 // set last value to infinity, only for aux
	for i := 0; i < n2; i++ {
		auxArray2[i] = in[q+1+i]
	}
	auxArray2[n2] = math.MaxInt32 // set last value to infinity, only for aux

	// merge two aux arrays into original array
	i, j := 0, 0
	for k := p; k <= r; k++ {
		if auxArray1[i] <= auxArray2[j] {
			in[k] = auxArray1[i]
			i++
		} else {
			in[k] = auxArray2[j]
			j++
		}
	}

}

/************************* With aux array implementation *************************/
