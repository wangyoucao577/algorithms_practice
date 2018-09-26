package maxsubarray

// FindMaxSubarrayByViolent find a maximum subarray (maximum sum of the subarray) from the input array,
// by voilent solution
// return the the subarray
func FindMaxSubarrayByViolent(in []int) []int {
	if in == nil || len(in) == 0 {
		return in
	}

	maxLow, maxHigh := 0, 0
	maxSum := in[0]

	for i := 0; i < len(in); i++ {

		sum := in[i]
		for j := i + 1; j < len(in); j++ {
			sum += in[j]

			if sum > maxSum {
				maxLow = i
				maxHigh = j
				maxSum = sum
			}
		}
	}

	return in[maxLow : maxHigh+1]
}
