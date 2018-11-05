package btree

import (
	"fmt"
)

type treeNode struct {
	n      int         // how many keys in this node
	keys   []int       // keys stored in this node, max n keys
	isLeaf bool        // whether this node is a leaf
	childs []*treeNode // child pointers, max n+1 child nodes

	payload interface{} // satellite information
}

func (b BTree) validateNode(x *treeNode) error {
	if x == nil {
		return nil
	}

	if x != b.root {
		if x.n <= b.t-1 || x.n >= 2*b.t-1 {
			return fmt.Errorf("treeNode %v keys count %v out of range [%v,%v]", x, x.n, b.t-1, 2*b.t-1)
		}
	}

	for i := 1; i < x.n; i++ {
		if x.keys[i-1] > x.keys[i] {
			return fmt.Errorf("expect keys[%v] <= keys[%v], but got %v > %v", i-1, i, x.keys[i-1], x.keys[i])
		}
	}

	return nil
}

func (b BTree) isFullNode(x *treeNode) (bool, error) {
	if x == nil {
		return false, fmt.Errorf("node is nil")
	}
	if x == b.root {
		return false, fmt.Errorf("node is root")
	}

	if x.n == 2*b.t-1 {
		return true, nil
	}
	return false, nil
}

func (b BTree) searchNode(x *treeNode, key int) (*treeNode, int) {
	if x == nil {
		return nil, 0
	}

	i := 0
	for i < x.n && key > x.keys[i] {
		i++
	}

	if i < x.n && x.keys[i] == key {
		return x, i
	}
	if x.isLeaf { // else if
		return nil, 0
	}

	diskRead(x.childs[i])
	return b.searchNode(x.childs[i], key)
}
