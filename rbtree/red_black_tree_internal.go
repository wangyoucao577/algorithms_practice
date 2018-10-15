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
	for node != rb.nil() {
		if key == node.key {
			return node
		}
		if key < node.key {
			node = node.leftChild
		} else {
			node = node.rightChild
		}
	}
	return rb.nil()
}

func (rb RBTree) minimumNode(node *treeNode) *treeNode {
	for node.leftChild != rb.nil() {
		node = node.leftChild
	}
	return node
}

func (rb RBTree) maximumNode(node *treeNode) *treeNode {
	for node.rightChild != rb.nil() {
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

//	|          |
//	x          y
//	 \	==>   /
//    y      x
// Test idea: will get same inorderTreeWalk result whatever rotate which node
func (rb *RBTree) leftRotate(x *treeNode) {

	y := x.rightChild          // set y
	x.rightChild = y.leftChild // turn y's left subtree into x's right subtree
	if y.leftChild != rb.nil() {
		y.leftChild.parent = x
	}
	y.parent = x.parent // link x's parent to y
	if x.parent == rb.nil() {
		rb.root = y
	} else {
		if x.parent.leftChild == x {
			x.parent.leftChild = y
		} else {
			x.parent.rightChild = y
		}
	}

	y.leftChild = x // put x on y's left
	x.parent = y
}

//	 |       |
//	 y       x
//  /	==>   \
// x           y
// Test idea: will get same inorderTreeWalk result whatever rotate which node
func (rb *RBTree) rightRotate(y *treeNode) {
	x := y.leftChild           // set x
	y.leftChild = x.rightChild // turn x's right subtree into y's left subtree
	if x.rightChild != rb.nil() {
		x.rightChild.parent = y
	}
	x.parent = y.parent // link y's parent to x
	if y.parent == rb.nil() {
		rb.root = x
	} else {
		if y == y.parent.leftChild {
			y.parent.leftChild = x
		} else {
			y.parent.rightChild = x
		}
	}
	x.rightChild = y // put y on x's right
	y.parent = x
}

func (rb *RBTree) insertFixup(z *treeNode) {
	defer func() { rb.root.color = rbBlack }() // case 0

	for z.parent.color == rbRed {
		if z.parent.parent.rightChild.color == rbRed && z.parent.parent.leftChild.color == rbRed { // z.parent.parent must exist here
			z.parent.parent.color = rbRed              // case 1
			z.parent.parent.rightChild.color = rbBlack // case 1
			z.parent.parent.leftChild.color = rbBlack  // case 1
			z = z.parent.parent                        // case 1
			continue
		}

		// case 2
		if z.parent == z.parent.parent.leftChild && z == z.parent.rightChild {
			z = z.parent // for case 3 operation
			rb.leftRotate(z)
		} else if z.parent == z.parent.parent.rightChild && z == z.parent.leftChild {
			z = z.parent // for case 3 operation
			rb.rightRotate(z)
		}

		// case 3
		if z.parent == z.parent.parent.leftChild && z == z.parent.leftChild {
			z.parent.color = rbBlack
			z.parent.parent.color = rbRed
			rb.rightRotate(z.parent.parent)
		} else if z.parent == z.parent.parent.rightChild && z == z.parent.rightChild {
			z.parent.color = rbBlack
			z.parent.parent.color = rbRed
			rb.leftRotate(z.parent.parent)
		}
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

// use `v` instead of `u`, `u` must not be `nil` here
func (rb *RBTree) transplant(u, v *treeNode) {
	if u == nil || v == nil || u == rb.nil() {
		return // could be panic actually
	}

	if u.parent == rb.nil() {
		rb.root = v
	} else {
		if u == u.parent.leftChild {
			u.parent.leftChild = v
		} else {
			u.parent.rightChild = v
		}
	}

	if v != rb.nil() {
		v.parent = u.parent
	}
}

func (rb *RBTree) deleteFixup(x *treeNode) {
	//TODO: implementation
}
