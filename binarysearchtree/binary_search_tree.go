// Package binarysearchtree implemented a binary search tree structure
package binarysearchtree

import (
	"fmt"
)

// BinarySearchTree represent a binary search tree structure
type BinarySearchTree struct {
	root  *treeNode
	count int
}

// Empty to check whether tree is empty,
// return true if empty, false if not empty
func (b BinarySearchTree) Empty() bool {
	return b.root == nil
}

// Clear to clear all nodes of the tree
func (b *BinarySearchTree) Clear() {
	b.root = nil
	b.count = 0
}

// Search to find the key in the binarySearchTree,
// return payload if found, nil if not found
// NOTE that if there're multiple nodes with same key, will only return the first one
func (b BinarySearchTree) Search(key int) (interface{}, error) {
	node := b.searchNode(key)
	if node == nil {
		return nil, fmt.Errorf("key %v does not exist", key)
	}
	return node.payload, nil
}

// Minimum return the minimum key node in the binary search tree
// return key and payload
func (b BinarySearchTree) Minimum() (int, interface{}, error) {
	if b.Empty() {
		return 0, nil, fmt.Errorf("empty tree")
	}
	node := minimumNode(b.root)
	return node.key, node.payload, nil
}

// Maximum return the maximum key node in the binary search tree
// return key and payload
func (b BinarySearchTree) Maximum() (int, interface{}, error) {
	if b.Empty() {
		return 0, nil, fmt.Errorf("empty tree")
	}
	node := maximumNode(b.root)
	return node.key, node.payload, nil
}

// Successor return the smallest newKey, which newKey > key
func (b BinarySearchTree) Successor(key int) (int, error) {
	node := b.searchNode(key)
	if node == nil {
		return 0, fmt.Errorf("key %v does not exist", key)
	}

	if node.rightChild != nil {
		newNode := minimumNode(node.rightChild)
		return newNode.key, nil
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
		newNode := maximumNode(node.leftChild)
		return newNode.key, nil
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

// Count how many nodes in the tree
func (b BinarySearchTree) Count() int {
	return b.count
}

// Insert to insert a key-payload pair into tree
func (b *BinarySearchTree) Insert(key int, payload interface{}) {

	newNode := &treeNode{}
	newNode.key = key
	newNode.payload = payload

	if b.root == nil {
		b.root = newNode
		b.count++
		return
	}

	node := b.root
	for node != nil {
		if newNode.key < node.key {
			if node.leftChild == nil {
				node.leftChild = newNode
				newNode.parent = node
				b.count++
				return
			}
			node = node.leftChild
		} else { // >= can allow multiple nodes have same key
			if node.rightChild == nil {
				node.rightChild = newNode
				newNode.parent = node
				b.count++
				return
			}
			node = node.rightChild
		}
	}
}

// Delete to delete the node with the key from the tree
func (b *BinarySearchTree) Delete(key int) error {
	node := b.searchNode(key)
	if node == nil {
		return fmt.Errorf("key %v does not exist", key)
	}

	if node.leftChild == nil { // only have a rightChild or no child
		b.transplant(node, node.rightChild) //whatever node.rightChild is nil or not nil
	} else if node.rightChild == nil { // only have a leftChild
		b.transplant(node, node.leftChild)
	} else { // both have leftChild and rightChild

		// Here's the most complex process of binary search tree.
		// Refer to "Introduction to Algorithms - Third Edition" ch12.3 for more details if necessary.

		// find successor of current node
		// The successor of node MUST be in right-sub-tree of current node,
		// and it will not have a leftChild.
		y := minimumNode(node.rightChild)

		if y != node.rightChild { // y in node's subTree but not the rightChild of node
			b.transplant(y, y.rightChild)
			y.rightChild = node.rightChild
			y.rightChild.parent = y
		}

		b.transplant(node, y)
		y.leftChild = node.leftChild
		y.leftChild.parent = y
	}

	b.count--
	return nil
}

// Validate to validate whether b is a valid binary search tree.
// return true if it's valid, false if it's not match the property of binary search tree.
func (b BinarySearchTree) Validate() bool {
	if b.Empty() {
		return true
	}

	var walked []int
	b.InorderTreeWalk(func(key int, payload interface{}) {
		walked = append(walked, key)
	})

	if len(walked) != b.count {
		return false
	}

	for i := 1; i < len(walked); i++ {
		if walked[i-1] > walked[i] {
			return false
		}
	}

	return true
}
