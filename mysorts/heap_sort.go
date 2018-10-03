package mysorts

// heapSort implements Sort by MaxHeap in-place
func heapSort(in myInterface) {
	if in.Len() <= 1 {
		return
	}

	bh := binaryHeap{}
	bh.buildMaxHeap(in)

	for i := in.Len(); i >= 2; i-- { // 1 started index here
		in.Swap(0, i-1) // move biggest element to the end
		bh.heapSize--
		bh.maxHeapity(in, 1) // then find the biggest and move it to root again
	}
}

// binaryHeap assoicate binary heap related methods
// NOTE: an outside slice will be operated based on these methods.
type binaryHeap struct {
	heapSize int
}

func (m binaryHeap) leftChild(i int) int {
	return i * 2
}
func (m binaryHeap) rightChild(i int) int {
	return i*2 + 1
}
func (m binaryHeap) parent(i int) int {
	return i / 2
}

// maxHeapity to keep max heap properties
// i is node index started by 1, only will be translated to 0 started index until operate array
func (m binaryHeap) maxHeapity(in myInterface, i int) {
	// i will be the parent in the maxHeapity procedure
	left := m.leftChild(i)
	right := m.rightChild(i)

	largest := i

	// find largest from i,left,right
	if left <= m.heapSize && in.Less(largest-1, left-1) { // translate to 0 started index to operate array
		largest = left
	}
	if right <= m.heapSize && in.Less(largest-1, right-1) { // translate to 0 started index to operate array
		largest = right
	}

	if largest != i {
		in.Swap(largest-1, i-1) // translate to 0 started index to operate array
		m.maxHeapity(in, largest)
	}
}

func (m *binaryHeap) buildMaxHeap(in myInterface) {
	m.heapSize = in.Len()

	// start at in.Len()/2 because all leaf nodes don't need to do the `maxheapity`
	for i := in.Len() / 2; i >= 1; i-- { // 1 started index here
		m.maxHeapity(in, i)
	}
}
