package minspanningtree

import (
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graphsample5"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

func TestMinSpanningTree(t *testing.T) {

	want := weightedgraph.Weight(37)

	//Kruskal algorithm
	km, err := Kruskal(*graphsample5.GraphSample())
	if err != nil {
		t.Error(err)
	}
	kmw, err := km.Weight()
	if err != nil {
		t.Error(err)
	}
	if kmw != want {
		t.Errorf("minimum spanning tree by kruskal algorithm, weight %v but want %v", kmw, want)
	}

	//Prim algorithm
	pm, err := Prim(*graphsample5.GraphSample())
	if err != nil {
		t.Error(err)
	}
	pmw, err := pm.Weight()
	if err != nil {
		t.Error(err)
	}
	if pmw != want {
		t.Errorf("minimum spanning tree by prim algorithm, weight %v but want %v", pmw, want)
	}

}
