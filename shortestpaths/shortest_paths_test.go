package shortestpaths

import (
	"reflect"
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/graphsample6"
	"github.com/wangyoucao577/algorithms_practice/graphsample8"
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

	wantShortestPaths := []struct {
		destination graph.NodeID
		path        graph.Path
	}{
		{3, graph.Path{0, 3}},
		{1, graph.Path{0, 3, 2, 1}},
		{4, graph.Path{0, 3, 2, 1, 4}},
	}

	for _, v := range wantShortestPaths {
		path, _ := sp.RetrievePath(v.destination)
		if !reflect.DeepEqual(path, v.path) {
			t.Errorf("want shortest path %v, but got %v", v.path, path)
		}
	}
}

func TestDijkstra(t *testing.T) {
	sp := Dijkstra(graphsample8.GraphSample(), graphsample8.NameToID("s"))

	wantCosts := []weightedgraph.Weight{0, 8, 9, 5, 7}
	for k, v := range sp.nodesMap {
		if v.d != wantCosts[k] {
			t.Errorf("Dijkstra calulated, expect cost %v from %v to %v, but got %v",
				v.d, sp.s, k, wantCosts[k])
		}
	}

	wantShortestPaths := []struct {
		destination graph.NodeID
		path        graph.Path
	}{
		{3, graph.Path{0, 3}},
		{1, graph.Path{0, 3, 1}},
		{4, graph.Path{0, 3, 4}},
	}

	for _, v := range wantShortestPaths {
		path, _ := sp.RetrievePath(v.destination)
		if !reflect.DeepEqual(path, v.path) {
			t.Errorf("want shortest path %v, but got %v", v.path, path)
		}
	}

}
