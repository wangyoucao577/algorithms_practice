package mysorts

// HeapSort implements Sort by MaxHeap in-place
func HeapSort(in myInterface) {
	if in.Len() <= 1 {
		return
	}

	bh := binaryHeap{}
	bh.maxHeapity = bh.maxHeapityLoopBased //switch loop or rescurse based maxheapity here

	bh.buildMaxHeap(in) // first build max heap

	for i := in.Len(); i >= 2; i-- { // 1 started index here
		in.Swap(0, i-1) // move biggest element to the end
		bh.heapSize--

		// then find the biggest and move it to root again
		//bh.maxHeapityRecursed(in, 1)
		//bh.maxHeapityLoopBased(in, 1)
		bh.maxHeapity(in, 1)
	}
}

// binaryHeap assoicate binary heap related methods
// NOTE: an outside slice will be operated based on these methods.
type binaryHeap struct {
	heapSize int

	maxHeapity func(myInterface, int) //maxHeapity func
}

/********************** Be aware that `i` in below functions are all started by 1, not 0 *****************/

func (m binaryHeap) leftChild(i int) int {
	return i * 2
}
func (m binaryHeap) rightChild(i int) int {
	return i*2 + 1
}
func (m binaryHeap) parent(i int) int {
	return i / 2
}

// maxHeapityRecurseBased to keep max heap properties recursed
// i is node index started by 1, only will be translated to 0 started index until operate array
func (m *binaryHeap) maxHeapityRecurseBased(in myInterface, i int) {
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
		m.maxHeapityRecurseBased(in, largest)
	}
}

// maxHeapityLoopBased to keep max heap properties by loop
// i is node index started by 1, only will be translated to 0 started index until operate array
func (m *binaryHeap) maxHeapityLoopBased(in myInterface, i int) {

	for {

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

		if largest == i {
			break
		}

		in.Swap(largest-1, i-1) // translate to 0 started index to operate array
		i = largest
	}
}

func (m *binaryHeap) buildMaxHeap(in myInterface) {
	m.heapSize = in.Len()

	// start at in.Len()/2 because all leaf nodes don't need to do the `maxheapity`
	for i := in.Len() / 2; i >= 1; i-- { // 1 started index here
		//m.maxHeapityRecursed(in, i)
		//m.maxHeapityLoopBased(in, i)
		m.maxHeapity(in, i)
	}
}
