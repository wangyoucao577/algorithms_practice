package mysorts

import (
	"math"
)

// RadixSort implement a decimal based radix sort
func RadixSort(in []int) {
	if len(in) <= 1 {
		return
	}

	// find max value from in
	k := in[0]
	for _, v := range in {
		if k < v {
			k = v
		}
	}

	s := 10 // decimal
	var d int
	newK := k
	for newK > 0 {
		d++
		newK /= s
	}

	internalCountingSort := func(in []int, p int) { // MUST be stable sort

		counting := make([]int, s, s) // all elements will be initialized by 0
		out := make([]int, len(in), len(in))

		for _, v := range in {
			pv := (v / p) % s
			counting[pv]++
		}
		for i := 1; i < s; i++ {
			counting[i] += counting[i-1]
		}

		for j := len(in) - 1; j >= 0; j-- {
			v := in[j]
			pv := (v / p) % s
			out[counting[pv]-1] = v
			counting[pv]--
		}

		// copy out back to in
		in = in[:0]
		in = append(in, out...)
	}

	for i := 0; i < d; i++ {
		p := (int)(math.Pow10(i))
		internalCountingSort(in, p)
	}

}
