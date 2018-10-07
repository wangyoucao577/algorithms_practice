package binarysearchtree

type treeNode struct {
	parent     *treeNode
	leftChild  *treeNode
	rightChild *treeNode

	key     int
	payload interface{}
}

func (b *BinarySearchTree) searchNode(key int) *treeNode {

	node := b.root
	for node != nil {
		if key == node.key {
			return node
		}
		if key < node.key {
			node = node.leftChild
		} else {
			node = node.rightChild
		}
	}
	return nil
}

func inorderTreeWalk(node *treeNode, action IterateAction) {
	if node != nil {
		inorderTreeWalk(node.leftChild, action)
		action(node.key, node.payload)
		inorderTreeWalk(node.rightChild, action)
	}
}

func preorderTreeWalk(node *treeNode, action IterateAction) {
	if node != nil {
		action(node.key, node.payload)
		preorderTreeWalk(node.leftChild, action)
		preorderTreeWalk(node.rightChild, action)
	}
}

func postorderTreeWalk(node *treeNode, action IterateAction) {
	if node != nil {
		postorderTreeWalk(node.leftChild, action)
		postorderTreeWalk(node.rightChild, action)
		action(node.key, node.payload)
	}
}
