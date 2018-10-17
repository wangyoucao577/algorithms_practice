package mysorts

import "github.com/wangyoucao577/algorithms_practice/binarysearchtree"

// TreeSort execute sort based by binary search tree, not in-place sort
func TreeSort(in []int) {
	if len(in) <= 1 {
		return
	}

	var tree binarysearchtree.BinarySearchTree
	for _, v := range in {
		tree.Insert(v, nil)
	}

	in = in[:0]

	tree.InorderTreeWalk(func(key int, payload interface{}) {
		in = append(in, key)
	})
}
