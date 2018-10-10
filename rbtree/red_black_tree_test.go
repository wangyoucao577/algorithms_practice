package rbtree

import (
	"math/rand"
	"testing"
)

func TestBinarySearchTreeRandomizedInsertDelete(t *testing.T) {

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
