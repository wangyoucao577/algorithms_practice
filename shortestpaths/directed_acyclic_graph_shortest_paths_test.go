package shortestpaths

import (
	"reflect"
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/graphsample7"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

func TestDirectedAcyclicGraphShortestPaths(t *testing.T) {
	sp, err := DirectedAcyclicGraphShortestPaths(graphsample7.GraphSample(), graphsample7.NameToID("s"))
	if err != nil {
		t.Error(err)
	}

	wantShortestPaths := []struct {
		destination graph.NodeID
		path        graph.Path
		weight      weightedgraph.Weight
	}{
		{3, graph.Path{1, 3}, 6},
		//{1, graph.Path{1}, 0},
		//{4, graph.Path{1, 3, 4}, 5},
		//{5, graph.Path{1, 3, 4, 5}, 3},
		//{0, graph.Path{}, infinitelyWeight},
	}

	for _, v := range wantShortestPaths {
		path, weight := sp.RetrievePath(v.destination)
		if weight != v.weight {
			t.Errorf("to %v, want shortest path weight %v, but got %v", v.destination, v.weight, weight)
		}
		if !reflect.DeepEqual(path, v.path) {
			t.Errorf("want shortest path %v, but got %v", v.path, path)
		}
	}

}
