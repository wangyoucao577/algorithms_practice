package rbtree

import (
	"math/rand"
	"reflect"
	"testing"
)

// similar functional test comes from binarysearchtree.TestBinarySearchTreeSample1()
func TestRedBlackTreeSample1(t *testing.T) {

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
		preorderWalked:  []int{13, 6, 3, 2, 4, 7, 9, 17, 15, 18, 20},
		postorderWalked: []int{2, 4, 3, 9, 7, 6, 15, 20, 18, 17, 13},
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

	payloadFactor := 2 // use a simple factor for payload, easy to verify

	// construct tree
	tree := NewRBTree()
	tree.Insert(15, 15*payloadFactor)
	tree.Insert(6, 6*payloadFactor)
	tree.Insert(3, 3*payloadFactor)
	tree.Insert(2, 2*payloadFactor)
	tree.Insert(4, 4*payloadFactor)
	tree.Insert(7, 7*payloadFactor)
	tree.Insert(13, 13*payloadFactor)
	tree.Insert(9, 9*payloadFactor)
	tree.Insert(18, 18*payloadFactor)
	tree.Insert(17, 17*payloadFactor)
	tree.Insert(20, 20*payloadFactor)

	/* the new red-black tree will be:
			 13
			/   \
		   6     17
		  / \    /  \
		 3   7  15  18
		/ \   \       \
	   2   4   9      20
	*/

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

	for _, key := range want.postorderWalked {
		payload, err := tree.Search(key)
		if err != nil {
			t.Errorf("search key %v failed, err %v", key, err)
		}
		payloadValue := payload.(int)
		if payloadValue != key*payloadFactor {
			t.Errorf("search key %v, expect payload %v but got %v", key, key*payloadFactor, payloadValue)
		}
	}
	_, err := tree.Search(100)
	if err == nil {
		t.Error("search key 100, expect failed but return successed")
	}

	if tree.Delete(13) != nil {
		t.Error("Delete 13, expect succeed but failed")
	}
	if tree.Delete(13) == nil {
		t.Error("Delete 13 again, expect failed but succeed")
	}

	tree.Clear()
	if !tree.Empty() || tree.Count() > 0 {
		t.Errorf("Expect empty tree after clear, but got count %v", tree.Count())
	}
}

func TestRedBlackTreeRandomizedInsertDelete(t *testing.T) {

	maxTreeCount := 100
	maxRandomCount := 1000
	maxTestKeyCount := 2000

	for i := 0; i < maxTreeCount; i++ {

		tree := NewRBTree()
		var insert bool
		keys := make([]bool, maxTestKeyCount, maxTestKeyCount)

		for j := 0; j < maxRandomCount; j++ {

			if rand.Intn(2) != 0 {
				insert = true
			}

			key := rand.Intn(maxTestKeyCount)
			if insert {
				countBeforeInsert := tree.Count()
				tree.Insert(key, key)
				countAfterInsert := tree.Count()

				if countAfterInsert != countBeforeInsert+1 {
					t.Errorf("insert key %v, but count before insert %v +1 != count after insert %v", key, countBeforeInsert, countAfterInsert)
				}
				keys[key] = true

				if !tree.Validate() {
					t.Errorf("tree going to invalid after insert key %v", key)
				}

			} else { //delete
				countBeforeDelete := tree.Count()
				deleteErr := tree.Delete(key)
				countAfterDelete := tree.Count()

				if keys[key] {
					// expect delete succeed

					if deleteErr != nil {
						t.Errorf("delete key %v except succeed, but got failed, err %v", key, deleteErr)
					}

					if countAfterDelete != countBeforeDelete-1 {
						t.Errorf("delete key %v except succeed, but count before delete %v -1 != count after delete %v", key, countBeforeDelete, countAfterDelete)
					}

				} else {
					// expect delete failed

					if deleteErr == nil {
						t.Errorf("delete key %v except failed, but got succeed", key)
					}

					if countAfterDelete != countBeforeDelete {
						t.Errorf("delete key %v except failed, but count before delete %v != count after delete %v", key, countBeforeDelete, countAfterDelete)
					}

				}
				keys[key] = false

				if !tree.Validate() {
					t.Errorf("tree going to invalid after delete key %v", key)
				}

			}
		}
	}
}
