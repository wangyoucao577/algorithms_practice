package mysorts

// BucketSort implement bucket sort by aux array,
// please be aware that it expect input array with Uniform distribution to get liner O(n)
func BucketSort(in []int) {
	if len(in) <= 1 {
		return
	}

	bucketsCount := len(in)

	// find max value from in
	k := in[0]
	for _, v := range in {
		if k < v {
			k = v
		}
	}
	piece := k / bucketsCount
	if piece == 0 {
		piece = 1 // to avoid divide by zero
	}

	// initialize buckets
	buckets := make([]bucket, bucketsCount, bucketsCount)

	// split elements into buckets
	for _, v := range in {
		index := v / piece
		if index >= bucketsCount {
			index = bucketsCount - 1 //set to last bucket if overflow
		}
		buckets[index] = append(buckets[index], v)
	}

	// sort each bucket
	for i := range buckets {
		InsertionSort(buckets[i])
	}

	// copy back to in
	in = in[:0]
	for _, v := range buckets {
		in = append(in, v...)
	}
}

type bucket []int

func (b bucket) Len() int           { return len(b) }
func (b bucket) Less(i, j int) bool { return b[i] < b[j] }
func (b bucket) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
