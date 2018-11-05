package singlesourceshortestpaths

import (
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graphsample6"
)

func TestSingleSourceShortestPathSearch(t *testing.T) {

	if !BellmanFord(graphsample6.GraphSample(), graphsample6.NameToID("s")) {
		t.Errorf("expect got shortest paths from node s but failed.")
	}
}
