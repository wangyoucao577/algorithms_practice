package mysorts

// CountingSort implements Counting Sort with 2 aux array
// make sure all elements of `in`  >= 0
func CountingSort(in []int) {
	if len(in) <= 1 {
		return
	}

	// find max value from in
	k := findMaxElement(in)

	counting := make([]int, k+1, k+1) // all elements will be initialized by 0
	out := make([]int, len(in), len(in))

	for _, v := range in {
		counting[v]++
	}
	for i := 1; i < k+1; i++ {
		counting[i] += counting[i-1]
	}

	for j := len(in) - 1; j >= 0; j-- {
		v := in[j]
		out[counting[v]-1] = v
		counting[v]--
	}

	// copy out back to in
	in = in[:0]
	in = append(in, out...)
}
