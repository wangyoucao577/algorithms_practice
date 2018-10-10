package rbtree

type rbColor int

const (
	rbBlack rbColor = iota
	rbRed
)

type treeNode struct {
	parent     *treeNode
	leftChild  *treeNode
	rightChild *treeNode

	key     int
	payload interface{}

	color rbColor // the only new item compare to binarysearchtree node
}

// use the nil() to represent a nil node
func (rb RBTree) nil() *treeNode {
	return rb.nilNode
}

func (rb RBTree) searchNode(key int) *treeNode {

	node := rb.root
	for node != rb.nilNode {
		if key == node.key {
			return node
		}
		if key < node.key {
			node = node.leftChild
		} else {
			node = node.rightChild
		}
	}
	return rb.nilNode
}

func (rb RBTree) minimumNode(node *treeNode) *treeNode {
	for node.leftChild != rb.nilNode {
		node = node.leftChild
	}
	return node
}

func (rb RBTree) maximumNode(node *treeNode) *treeNode {
	for node.rightChild != rb.nilNode {
		node = node.rightChild
	}
	return node
}

type iterateActionOnNode func(*treeNode)

func (rb RBTree) inorderTreeWalkNodes(node *treeNode, action iterateActionOnNode) {
	if node != rb.nil() {
		rb.inorderTreeWalkNodes(node.leftChild, action)
		action(node)
		rb.inorderTreeWalkNodes(node.rightChild, action)
	}
}

func (rb RBTree) inorderTreeWalk(node *treeNode, action IterateAction) {
	if node != rb.nil() {
		rb.inorderTreeWalk(node.leftChild, action)
		action(node.key, node.payload)
		rb.inorderTreeWalk(node.rightChild, action)
	}
}

func (rb RBTree) preorderTreeWalk(node *treeNode, action IterateAction) {
	if node != rb.nil() {
		action(node.key, node.payload)
		rb.preorderTreeWalk(node.leftChild, action)
		rb.preorderTreeWalk(node.rightChild, action)
	}
}

func (rb RBTree) postorderTreeWalk(node *treeNode, action IterateAction) {
	if node != rb.nil() {
		rb.postorderTreeWalk(node.leftChild, action)
		rb.postorderTreeWalk(node.rightChild, action)
		action(node.key, node.payload)
	}
}

func (rb RBTree) validateBinarySearchTreeProperties() bool {
	if rb.Empty() {
		return true
	}

	// same code with binarysearchtree.Validate()

	var walked []int
	rb.InorderTreeWalk(func(key int, payload interface{}) {
		walked = append(walked, key)
	})

	if len(walked) != rb.count {
		return false
	}

	for i := 1; i < len(walked); i++ {
		if walked[i-1] > walked[i] {
			return false
		}
	}
	return true
}
