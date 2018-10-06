package mysorts

func findMaxElement(in []int) int {
	k := in[0]
	for _, v := range in {
		if k < v {
			k = v
		}
	}

	return k
}
