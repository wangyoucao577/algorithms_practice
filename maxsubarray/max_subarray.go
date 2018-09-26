// Package maxsubarray try to solve the maximum subarray problem
package maxsubarray

// findMaxSubarray find a maximum subarray (maximum sum of the subarray) from the input array,
// return the the subarray
func findMaxSubarray(in []int) []int {
	if in == nil || len(in) == 0 {
		return in
	}

	l, r, _ := findMaxSubarrayImpl(in, 0, len(in)-1)
	out := in[l : r+1]
	return out
}

// findMaxSubarrayImpl return left-index, right-index and sum value of the subarray
func findMaxSubarrayImpl(in []int, low, high int) (int, int, int) {

	if low == high {
		return low, high, in[low]
	}

	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := findMaxSubarrayImpl(in, low, mid)
	rightLow, rightHigh, rightSum := findMaxSubarrayImpl(in, mid+1, high)
	crossingLow, crossingHigh, crossingSum := findMaxCrossingSubarray(in, low, mid, high)

	if leftSum >= rightSum && leftSum >= crossingSum {
		return leftLow, leftHigh, leftSum
	}
	//else if
	if rightSum >= leftSum && rightSum >= crossingSum {
		return rightLow, rightHigh, rightSum
	}
	//else
	return crossingLow, crossingHigh, crossingSum
}

// findMaxCrossingSubarray return left-index, right-index and sum value of the crossing subarray
func findMaxCrossingSubarray(in []int, low, mid, high int) (int, int, int) {

	leftLow := mid
	leftSum := in[mid]

	sum := leftSum
	for i := mid - 1; i >= low; i-- {
		sum += in[i]
		if sum > leftSum {
			leftLow = i
			leftSum = sum
		}
	}

	rightHigh := mid + 1
	rightSum := in[mid+1]

	sum = rightSum
	for i := mid + 2; i <= high; i++ {
		sum += in[i]
		if sum > rightSum {
			rightHigh = i
			rightSum = sum
		}
	}

	return leftLow, rightHigh, leftSum + rightSum
}
