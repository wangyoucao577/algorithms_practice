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
