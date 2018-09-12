package dfs

import (
	"reflect"
	"sort"
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graphsample4"
)

func TestSCC(t *testing.T) {

	want := []StronglyConnectedComponent{
		{0, 1, 4},
		{2, 3},
		{5, 6},
		{7},
	}

	sccs, err := SplitToStronglyConnectedComponents(graphsample4.AdjacencyListGraphSample())
	if err != nil {
		t.Error(err)
		return
	}

	// sort each component before compare, easy to compare
	for _, v := range sccs {
		sort.Sort(v)
	}

	if !reflect.DeepEqual(sccs, want) {
		t.Errorf("want strongly connected components %v\n, got %v\n", want, sccs)
	}
}
