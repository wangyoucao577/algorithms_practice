// Package binarysearchtree implemented a binary search tree structure
package binarysearchtree

import (
	"fmt"
)

// BinarySearchTree represent a binary search tree structure
type BinarySearchTree struct {
	root *treeNode
}

// Empty to check whether tree is empty,
// return true if empty, false if not empty
func (b BinarySearchTree) Empty() bool {
	return b.root == nil
}

// Search to find the key in the binarySearchTree,
// return payload if found, nil if not found
// NOTE that if there're multiple nodes with same key, will only return the first one
func (b BinarySearchTree) Search(key int) (interface{}, error) {
	node := b.searchNode(key)
	if node == nil {
		return nil, fmt.Errorf("key %v does not exist", key)
	}
	return node.key, nil
}

// Minimum return the minimum key node in the binary search tree
// return key and payload
func (b BinarySearchTree) Minimum() (int, interface{}, error) {
	if b.Empty() {
		return 0, nil, fmt.Errorf("empty tree")
	}
	node := b.root
	for node.leftChild != nil {
		node = node.leftChild
	}
	return node.key, node.payload, nil
}

// Maximum return the maximum key node in the binary search tree
// return key and payload
func (b BinarySearchTree) Maximum() (int, interface{}, error) {
	if b.Empty() {
		return 0, nil, fmt.Errorf("empty tree")
	}
	node := b.root
	for node.rightChild != nil {
		node = node.rightChild
	}
	return node.key, node.payload, nil
}

// Successor return the smallest newKey, which newKey > key
func (b BinarySearchTree) Successor(key int) (int, error) {
	node := b.searchNode(key)
	if node == nil {
		return 0, fmt.Errorf("key %v does not exist", key)
	}

	if node.rightChild != nil {
		subTree := BinarySearchTree{node.rightChild}
		newKey, _, err := subTree.Minimum()
		return newKey, err
	}

	p := node.parent
	for p != nil {
		if node == p.leftChild {
			return p.key, nil
		}

		node = p
		p = node.parent
	}
	return 0, fmt.Errorf("key %v already the max key", key)
}

// Predecessor return the biggest newKey, which newKey < key
func (b BinarySearchTree) Predecessor(key int) (int, error) {
	node := b.searchNode(key)
	if node == nil {
		return 0, fmt.Errorf("key %v does not exist", key)
	}

	if node.leftChild != nil {
		subTree := BinarySearchTree{node.leftChild}
		newKey, _, err := subTree.Maximum()
		return newKey, err
	}

	p := node.parent
	for p != nil {
		if node == p.rightChild {
			return p.key, nil
		}

		node = p
		p = node.parent
	}
	return 0, fmt.Errorf("key %v already the min key", key)
}

// IterateAction define interface to operate on key/payload during walk/iterate
type IterateAction func(int, interface{})

// InorderTreeWalk walk the tree by inorder
func (b BinarySearchTree) InorderTreeWalk(action IterateAction) {
	inorderTreeWalk(b.root, action)
}

// PreorderTreeWalk walk the tree by preorder
func (b BinarySearchTree) PreorderTreeWalk(action IterateAction) {
	preorderTreeWalk(b.root, action)
}

// PostorderTreeWalk walk the tree by postorder
func (b BinarySearchTree) PostorderTreeWalk(action IterateAction) {
	postorderTreeWalk(b.root, action)
}

// Insert to insert a key-payload pair into tree
func (b *BinarySearchTree) Insert(key int, payload interface{}) {

	newNode := &treeNode{}
	newNode.key = key
	newNode.payload = payload

	if b.root == nil {
		b.root = newNode
		return
	}

	node := b.root
	for node != nil {
		if newNode.key < node.key {
			if node.leftChild == nil {
				node.leftChild = newNode
				newNode.parent = node
				return
			}
			node = node.leftChild
		} else { // >= can allow multiple nodes have same key
			if node.rightChild == nil {
				node.rightChild = newNode
				newNode.parent = node
				return
			}
			node = node.rightChild
		}
	}
}

// Delete to delete the node with the key from the tree
func (b *BinarySearchTree) Delete(key int) {
	//TODO:
}
