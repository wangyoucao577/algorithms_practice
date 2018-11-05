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

func (b *BTree) allocateNode() *treeNode {
	return &treeNode{0, []int{}, false, []*treeNode{}, nil}
}

func (b BTree) validateNode(x *treeNode) error {
	if x == nil {
		return nil
	}

	if x.n != len(x.keys) {
		return fmt.Errorf("expect x.n == len(x.keys), but got %v != %v", x.n, len(x.keys))
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

	if !x.isLeaf {
		if x.n+1 != len(x.childs) {
			return fmt.Errorf("expect x.n+1 == len(x.childs), but got %v != %v", x.n+1, len(x.childs))
		}

		for i := 0; i < x.n; i++ {
			k1 := x.childs[i].keys[x.childs[i].n-1]
			k2 := x.childs[i+1].keys[0]

			if k1 > x.keys[i] || x.keys[i] > k2 {
				return fmt.Errorf("expect k1 <= x.key[%v] <= k2, but got %v %v %v", i, k1, x.keys[i], k2)
			}
		}
	}

	return nil
}

func (b BTree) isFullNode(x *treeNode) (bool, error) {
	if x == nil {
		return false, fmt.Errorf("node is nil")
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

func (b *BTree) splitChild(x *treeNode, i int) {
	z := b.allocateNode()
	y := x.childs[i] // y is the full node

	z.isLeaf = y.isLeaf
	z.n = b.t - 1
	z.keys = make([]int, z.n, z.n)
	for j := 0; j < b.t-1; j++ {
		z.keys[j] = y.keys[b.t+j]
	}
	if !y.isLeaf {
		z.childs = make([]*treeNode, z.n+1, z.n+1)
		for j := 0; j < b.t; j++ {
			z.childs[j] = y.childs[b.t+j]
		}
	}

	x.keys = append(x.keys, 0)
	x.childs = append(x.childs, nil)
	for j := x.n; j > i; j-- {
		x.keys[j] = x.keys[j-1]
		x.childs[j+1] = x.childs[j]
	}
	x.keys[i] = y.keys[b.t-1]
	x.childs[i+1] = z
	x.n++

	y.n = b.t - 1
	y.keys = y.keys[:b.t-1]
	y.childs = y.childs[:b.t-1]

	diskWrite(y)
	diskWrite(z)
	diskWrite(x)
}

func (b *BTree) insertNonFull(x *treeNode, key int) {

	i := x.n - 1
	if x.isLeaf {

		x.keys = append(x.keys, 0)
		for ; i >= 0 && key < x.keys[i]; i-- {
			x.keys[i+1] = x.keys[i]
		}
		x.keys[i+1] = key
		x.n++
		diskWrite(x)
	} else {

		for ; i >= 0 && key < x.keys[i]; i-- {
		}
		i = i + 1
		diskRead(x.childs[i])

		if ok, _ := b.isFullNode(x.childs[i]); ok {
			b.splitChild(x, i)
			if key > x.keys[i] {
				i = i + 1
			}
		}
		b.insertNonFull(x.childs[i], key)
	}
}
