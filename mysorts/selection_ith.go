package mysorts

// Select select i-th order statistic from the input array,
// return the index of the i-th element if succeed, -1 if failed.
// based on randomizedPartition() of quicksort
func Select(in myInterface, i int) int {
	if i < 0 || i >= in.Len() {
		return -1
	}

	//return selectImplByRescurse(in, i, 0, in.Len()-1)
	return selectImplByLoop(in, i, 0, in.Len()-1)
}

func selectImplByRescurse(in myInterface, i, p, r int) int {
	q := randomizedPartition(in, p, r)
	if q == i {
		return q
	}

	if q > i {
		return selectImplByRescurse(in, i, p, q-1)
	}
	return selectImplByRescurse(in, i, q+1, r)
}

func selectImplByLoop(in myInterface, i, p, r int) int {

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
