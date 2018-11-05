package singlesourceshortestpaths

import (
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graphsample6"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

func TestSingleSourceShortestPathSearch(t *testing.T) {

	ok, sp := BellmanFord(graphsample6.GraphSample(), graphsample6.NameToID("s"))
	if !ok {
		t.Errorf("expect got shortest paths from node s but failed.")
	}

	wantCosts := []weightedgraph.Weight{0, 2, 4, 7, -2}
	for k, v := range sp.nodesMap {
		if v.d != wantCosts[k] {
			t.Errorf("Bellman-ford calulated, expect cost %v from %v to %v, but got %v",
				v.d, sp.s, k, wantCosts[k])
		}
	}
}
