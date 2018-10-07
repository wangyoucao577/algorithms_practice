package binarysearchtree

import (
	"reflect"
	"testing"
)

func TestBinarySearchTreeSample1(t *testing.T) {

	/* This sample tree comes from
		"Introduction to Algorithms - Third Edition" 12.2

		define tree as below:
	                   15
	      			 /    \
		           6       18
				  / \      / \
				 3   7    17  20
				/ \   \
			   2   4   13
					  /
				     9
	*/

	want := struct {
		count           int
		min             int
		max             int
		inorderWalked   []int
		preorderWalked  []int
		postorderWalked []int
	}{
		count:           11,
		min:             2,
		max:             20,
		inorderWalked:   []int{2, 3, 4, 6, 7, 9, 13, 15, 17, 18, 20},
		preorderWalked:  []int{15, 6, 3, 2, 4, 7, 13, 9, 18, 17, 20},
		postorderWalked: []int{2, 4, 3, 9, 13, 7, 6, 17, 20, 18, 15},
	}

	wantPrecessorSuccessor := []struct {
		precessor int
		successor int
	}{
		{13, 15},
		{6, 7},
		{2, 3},
		{4, 6},
		{20, 2}, //error
	}

	// construct tree
	var tree BinarySearchTree
	tree.Insert(15, nil)
	tree.Insert(6, nil)
	tree.Insert(3, nil)
	tree.Insert(2, nil)
	tree.Insert(4, nil)
	tree.Insert(7, nil)
	tree.Insert(13, nil)
	tree.Insert(9, nil)
	tree.Insert(18, nil)
	tree.Insert(17, nil)
	tree.Insert(20, nil)

	if tree.Count() != want.count {
		t.Errorf("expect tree count %v but got %v", want.count, tree.Count())
	}

	gotMin, _, _ := tree.Minimum()
	if gotMin != want.min {
		t.Errorf("expect tree minimum %v but got %v", want.min, gotMin)
	}

	gotMax, _, _ := tree.Maximum()
	if gotMax != want.max {
		t.Errorf("expect tree maximum %v but got %v", want.max, gotMax)
	}

	gotInorderWalk := []int{}
	tree.InorderTreeWalk(func(key int, payload interface{}) {
		gotInorderWalk = append(gotInorderWalk, key)
	})
	if !reflect.DeepEqual(gotInorderWalk, want.inorderWalked) {
		t.Errorf("expect tree inorder walk %v but got %v", want.inorderWalked, gotInorderWalk)
	}

	gotPreorderWalk := []int{}
	tree.PreorderTreeWalk(func(key int, payload interface{}) {
		gotPreorderWalk = append(gotPreorderWalk, key)
	})
	if !reflect.DeepEqual(gotPreorderWalk, want.preorderWalked) {
		t.Errorf("expect tree preorder walk %v but got %v", want.preorderWalked, gotPreorderWalk)
	}

	gotPostorderWalk := []int{}
	tree.PostorderTreeWalk(func(key int, payload interface{}) {
		gotPostorderWalk = append(gotPostorderWalk, key)
	})
	if !reflect.DeepEqual(gotPostorderWalk, want.postorderWalked) {
		t.Errorf("expect tree postorder walk %v but got %v", want.postorderWalked, gotPostorderWalk)
	}

	for _, v := range wantPrecessorSuccessor {
		gotSuccessor, err := tree.Successor(v.precessor)
		if v.precessor == 20 { //expect fail
			if err == nil {
				t.Errorf("expect error for successor of key %v but got %v", v.precessor, gotSuccessor)
			}
			continue
		}
		if gotSuccessor != v.successor {
			t.Errorf("expect %v for successor of key %v but got %v", v.successor, v.precessor, gotSuccessor)
		}
	}

	for _, v := range wantPrecessorSuccessor {
		gotPrecessor, err := tree.Predecessor(v.successor)
		if v.successor == 2 { //expect fail
			if err == nil {
				t.Errorf("expect error for precessor of key %v but got %v", v.successor, gotPrecessor)
			}
			continue
		}
		if gotPrecessor != v.precessor {
			t.Errorf("expect %v for precessor of key %v but got %v", v.precessor, v.successor, gotPrecessor)
		}
	}

}
