// Package btree defined the b-tree strucutre and it's operations.
package btree

// BTree represent a b-tree type
type BTree struct {
	root *treeNode

	t int // b-tree minimum degree
}
