package mysorts

import "github.com/wangyoucao577/algorithms_practice/rbtree"

// RedBlackTreeSort execute sort based by red-black tree, not in-place sort
func RedBlackTreeSort(in []int) {
	if len(in) <= 1 {
		return
	}

	tree := rbtree.NewRBTree()
	for _, v := range in {
		tree.Insert(v, nil)
	}

	in = in[:0]

	tree.InorderTreeWalk(func(key int, payload interface{}) {
		in = append(in, key)
	})
}
