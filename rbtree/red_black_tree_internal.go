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

func (rb *RBTree) clearDoubleBlackNil() {
	rb.doubleBlackNil.parent = nil
	rb.doubleBlackNil.leftChild = nil
	rb.doubleBlackNil.rightChild = nil
	rb.doubleBlackNil.key = 0
	rb.doubleBlackNil.payload = nil
	rb.doubleBlackNil.color = rbBlack
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

	// 以下的 case xxx 都对应于 "红黑树-Delete-我的推导" 笔记中的case xxx,
	//  而非算法导论书上的 case
	// 本函数中主要会处理 case 2 的所有case, 以及case 4的着色

	for x != rb.root && x.color == rbBlack {
		if x == x.parent.leftChild {
			if x.parent.rightChild.color == rbRed { // case 2-4
				x.parent.color = rbRed              // case 2-4
				x.parent.rightChild.color = rbBlack // case 2-4
				rb.leftRotate(x.parent)             // case 2-4
				// then will going to case 2-1 or 2-2 or 2-3-1 in next loop
			} else {
				if x.parent.rightChild.rightChild.color == rbBlack && x.parent.rightChild.leftChild.color == rbBlack { // case 2-3
					if x.parent.color == rbRed { // case 2-3-1
						x.parent.color = rbBlack          // case 2-3-1
						x.parent.rightChild.color = rbRed // case 2-3-1

						// can be done here
						break // case 2-3-1
					} else { // case 2-3-2
						x.parent.rightChild.color = rbRed // case 2-3-2
						x = x.parent                      // case 2-3-2
						if x.leftChild == rb.doubleBlackNil {
							rb.transplant(x.leftChild, rb.nil()) // replace by normal sentinal nil node
						}
					}
				} else if x.parent.rightChild.rightChild.color == rbRed { // case 2-1
					x.parent.rightChild.rightChild.color = rbBlack // case 2-1
					x.parent.rightChild.color = x.parent.color     // case 2-1
					x.parent.color = rbBlack                       // case 2-1
					rb.leftRotate(x.parent)                        // case 2-1

					// can be done here
					break // case 2-1
				} else { // case 2-2, x.parent.rightChild.leftChild.color == rbRed
					x.parent.rightChild.color = rbRed             // case 2-2
					x.parent.rightChild.leftChild.color = rbBlack // case 2-2
					rb.rightRotate(x.parent.rightChild)           // case 2-2

					// next loop going into case 2-1
				}
			}
		} else { // 与上述处理对称的, x == x.parent.rightChild
			if x.parent.leftChild.color == rbRed { // case 2-4
				x.parent.color = rbRed             // case 2-4
				x.parent.leftChild.color = rbBlack // case 2-4
				rb.rightRotate(x.parent)           // case 2-4
				// then will going to case 2-1 or 2-2 or 2-3-1 in next loop
			} else {
				if x.parent.leftChild.rightChild.color == rbBlack && x.parent.leftChild.leftChild.color == rbBlack { // case 2-3
					if x.parent.color == rbRed { // case 2-3-1
						x.parent.color = rbBlack         // case 2-3-1
						x.parent.leftChild.color = rbRed // case 2-3-1

						// can be done here
						break // case 2-3-1
					} else { // case 2-3-2
						x.parent.leftChild.color = rbRed // case 2-3-2
						x = x.parent                     // case 2-3-2
						if x.rightChild == rb.doubleBlackNil {
							rb.transplant(x.rightChild, rb.nil()) // replace by normal sentinal nil node
						}
					}
				} else if x.parent.leftChild.leftChild.color == rbRed { // case 2-1
					x.parent.leftChild.leftChild.color = rbBlack // case 2-1
					x.parent.leftChild.color = x.parent.color    // case 2-1
					x.parent.color = rbBlack                     // case 2-1
					rb.rightRotate(x.parent)                     // case 2-1

					// can be done here
					break // case 2-1
				} else { // case 2-2, x.parent.leftChild.rightChild.color == rbRed
					x.parent.leftChild.color = rbRed              // case 2-2
					x.parent.leftChild.rightChild.color = rbBlack // case 2-2
					rb.leftRotate(x.parent.leftChild)             // case 2-2

					// next loop going into case 2-1
				}
			}
		}
	}

	if x == rb.doubleBlackNil {
		rb.transplant(x, rb.nil())
	}
	x.color = rbBlack //  case 4
}
