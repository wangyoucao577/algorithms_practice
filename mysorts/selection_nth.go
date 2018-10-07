package mysorts

// RandomizedSelectNth select n-th (n start by 1) order statistic from the input array,
// return the index of the n-th element if succeed, -1 if failed.
// based on randomizedPartition() of quicksort
func RandomizedSelectNth(in myInterface, n int) int {
	if n <= 0 || n > in.Len() {
		return -1
	}

	//return selectImplByRescurse(in, n-1, 0, in.Len()-1)
	return randomizedSelectImplByLoop(in, n-1, 0, in.Len()-1)
}

func randomizedSelectImplByRescurse(in myInterface, i, p, r int) int {
	q := randomizedPartition(in, p, r)
	if q == i {
		return q
	}

	if q > i {
		return randomizedSelectImplByRescurse(in, i, p, q-1)
	}
	return randomizedSelectImplByRescurse(in, i, q+1, r)
}

func randomizedSelectImplByLoop(in myInterface, i, p, r int) int {

	for {
		q := randomizedPartition(in, p, r)
		if q == i {
			return q
		}

		if q > i {
			r = q - 1
		} else {
			p = q + 1
		}
	}
}

// SelectNth select n-th (n start by 1) order statistic from the input array,
// return the index of the n-th element if succeed, -1 if failed.
func SelectNth(in []int, n int) int {
	if n <= 0 || n > len(in) {
		return -1
	}

	return selectNthImpl(subArray(in), n-1, 0, len(in)-1)
}

func selectNthImpl(in subArray, i, p, r int) int {

	medianV := selectNthMedianImpl(in, p, r)

	// partition by medianV
	q := partitionByX(in, medianV, p, r)

	if q == i {
		return q
	}

	if q > i {
		return selectNthImpl(in, i, p, q-1)
	}
	return selectNthImpl(in, i, q+1, r)
}

// partition by x
func partitionByX(in subArray, x, p, r int) int {

	k := p - 1
	xIndex := -1

	for j := p; j <= r; j++ {
		if in[j] <= x {
			k++

			if in[j] == x {
				xIndex = k //at least 1 element value == x
			}

			if k != j {
				in.Swap(k, j)
			}

		}
	}

	in.Swap(xIndex, k)
	return k
}

func selectNthMedianImpl(in subArray, p, r int) int {

	piece := 5
	sparseArray := subArray{}

	for j := p; j <= r; j += piece {
		median := piece / 2
		jr := j + piece
		if jr > r+1 {
			jr = r + 1
			median = (jr - j) / 2
		}

		// sort each array
		s := (subArray)(in[j:jr])
		InsertionSort(s)

		sparseArray = append(sparseArray, s[median])
	}

	if sparseArray.Len() > 1 {
		return selectNthMedianImpl(sparseArray, 0, sparseArray.Len()-1)
	}
	return sparseArray[0]
}

type subArray []int

func (s subArray) Len() int           { return len(s) }
func (s subArray) Less(i, j int) bool { return s[i] < s[j] }
func (s subArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
