package mysorts

type intSlice []int

func (s intSlice) Len() int           { return len(s) }
func (s intSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s intSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s intSlice) deepCopy() intSlice {
	newCopy := intSlice{}
	newCopy = append(newCopy, s...)
	return newCopy
}

func isSorted(in intSlice, ascending bool) bool {

	for i := 0; i < len(in)-1; i++ {
		if ascending {
			if in[i] > in[i+1] {
				return false
			}
		} else {
			if in[i] < in[i+1] {
				return false
			}
		}
	}
	return true
}
