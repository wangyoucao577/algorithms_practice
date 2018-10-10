package rbtree

import "fmt"

// RBTree represent a red-black tree structure
type RBTree struct {
	root    *treeNode
	nilNode *treeNode // sentinel node to represent nil
	count   int       // won't count the nil node
}

// NewRBTree create a new emtpy red-black tree
func NewRBTree() *RBTree {
	var tree RBTree
	tree.nilNode = &treeNode{nil, nil, nil, -1, nil, rbBlack}
	tree.root = tree.nilNode
	return &tree
}

// Empty to check whether tree is empty,
// return true if empty, false if not empty
func (rb RBTree) Empty() bool {
	return rb.root == rb.nil()
}

// Clear to clear all nodes of the tree
func (rb *RBTree) Clear() {
	rb.root = rb.nil()
	rb.count = 0
}

// Search to find the key in the red-black tree,
// return payload if found, nil if not found
// NOTE that if there're multiple nodes with same key, will only return the first one
func (rb RBTree) Search(key int) (interface{}, error) {
	node := rb.searchNode(key)
	if node == rb.nil() {
		return nil, fmt.Errorf("key %v does not exist", key)
	}
	return node.payload, nil
}

// Minimum return the minimum key node in the red-black tree
// return key and payload
func (rb RBTree) Minimum() (int, interface{}, error) {
	if rb.Empty() {
		return 0, nil, fmt.Errorf("empty tree")
	}
	node := rb.minimumNode(rb.root)
	return node.key, node.payload, nil
}

// Maximum return the maximum key node in the red-black tree
// return key and payload
func (rb RBTree) Maximum() (int, interface{}, error) {
	if rb.Empty() {
		return 0, nil, fmt.Errorf("empty tree")
	}
	node := rb.maximumNode(rb.root)
	return node.key, node.payload, nil
}

// Successor return the smallest newKey, which newKey > key
func (rb RBTree) Successor(key int) (int, error) {
	node := rb.searchNode(key)
	if node == rb.nil() {
		return 0, fmt.Errorf("key %v does not exist", key)
	}

	if node.rightChild != rb.nil() {
		newNode := rb.minimumNode(node.rightChild)
		return newNode.key, nil
	}

	p := node.parent
	for p != rb.nil() {
		if node == p.leftChild {
			return p.key, nil
		}

		node = p
		p = node.parent
	}
	return 0, fmt.Errorf("key %v already the max key", key)
}

// Predecessor return the biggest newKey, which newKey < key
func (rb RBTree) Predecessor(key int) (int, error) {
	node := rb.searchNode(key)
	if node == rb.nil() {
		return 0, fmt.Errorf("key %v does not exist", key)
	}

	if node.leftChild != rb.nil() {
		newNode := rb.maximumNode(node.leftChild)
		return newNode.key, nil
	}

	p := node.parent
	for p != rb.nil() {
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
func (rb RBTree) InorderTreeWalk(action IterateAction) {
	rb.inorderTreeWalk(rb.root, action)
}

// PreorderTreeWalk walk the tree by preorder
func (rb RBTree) PreorderTreeWalk(action IterateAction) {
	rb.preorderTreeWalk(rb.root, action)
}

// PostorderTreeWalk walk the tree by postorder
func (rb RBTree) PostorderTreeWalk(action IterateAction) {
	rb.postorderTreeWalk(rb.root, action)
}

// Count how many nodes in the tree
func (rb RBTree) Count() int {
	return rb.count
}

// Insert to insert a key-payload pair into tree
func (rb *RBTree) Insert(key int, payload interface{}) {
	//TODO: implement
}

// Delete to delete the node with the key from the tree
func (rb *RBTree) Delete(key int) error {

	//TODO: implement
	return nil
}

// Validate to validate whether b is a valid red-black tree.
// return true if it's valid, false if it's not match the property of red-black tree.
func (rb RBTree) Validate() bool {

	//0. First of all, it should be a binary search tree
	if !rb.validateBinarySearchTreeProperties() {
		return false
	}

	//1. Every node is either RED or BLACK.
	// Don't need to verify.

	//2. Every nil node is BLACK.
	if rb.nil().color != rbBlack || rb.nilNode.color != rbBlack {
		return false
	}

	//5. The root node is always BLACK.
	if rb.root.color != rbBlack {
		return false
	}

	//3. Every RED node has two BLACK child nodes.
	valid := true
	rb.inorderTreeWalkNodes(rb.root, func(node *treeNode) {
		if node.color == rbRed {
			if node.leftChild.color != rbBlack || node.rightChild.color != rbBlack {
				valid = false
			}
		}
	})
	if !valid {
		return false
	}

	//4. Every path from node x (calculte without x) down to leaf node(must be rb.nil()) has the same number of BLACK nodes.
	//TODO: implement

	return true
}
