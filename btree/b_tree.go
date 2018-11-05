// Package btree defined the b-tree strucutre and it's operations.
package btree

// BTree represent a b-tree type
type BTree struct {
	root *treeNode

	t int // b-tree minimum degree
}

// NewBTree create a new b-tree with minimum-degree
func NewBTree(minimumDegree int) *BTree {
	b := &BTree{nil, minimumDegree}

	x := b.allocateNode()
	x.isLeaf = true
	x.n = 0
	diskWrite(x)

	b.root = x
	return b
}

// Insert inserts a new key into the b-tree
func (b *BTree) Insert(key int) {
	r := b.root
	if ok, _ := b.isFullNode(r); ok {
		s := b.allocateNode()
		b.root = s
		s.isLeaf = false
		s.n = 0
		s.childs = append(s.childs, r)
		b.splitChild(s, 0)
		b.insertNonFull(s, key)
	} else {
		b.insertNonFull(r, key)
	}
}
