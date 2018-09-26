package maxsubarray

// FindMaxSubarrayByKadane find a maximum subarray (maximum sum of the subarray) from the input array,
// by Kadane Algorithm
// return the the subarray
func FindMaxSubarrayByKadane(in []int) []int {
	if in == nil || len(in) <= 1 {
		return in
	}

	maxSoFarLow, maxSoFarHigh := 0, 0
	maxSoFarSum := in[0]

	maxEndingHereLow := 0
	maxEndingHereSum := in[0]

	for i := 1; i < len(in); i++ {

		if maxEndingHereSum+in[i] > in[i] {
			maxEndingHereSum += in[i]
		} else {
			maxEndingHereSum = in[i]
			maxEndingHereLow = i
		}

		if maxEndingHereSum > maxSoFarSum {
			maxSoFarSum = maxEndingHereSum
			maxSoFarLow = maxEndingHereLow
			maxSoFarHigh = i
		}
	}

	return in[maxSoFarLow : maxSoFarHigh+1]
}
