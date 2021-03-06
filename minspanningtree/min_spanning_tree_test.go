package minspanningtree

import (
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graphsamples/graphsample5"
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

	//Prim2 algorithm - with my implemented min heap
	pm2, err := Prim2(*graphsample5.GraphSample())
	if err != nil {
		t.Error(err)
	}
	pmw2, err := pm2.Weight()
	if err != nil {
		t.Error(err)
	}
	if pmw != want {
		t.Errorf("minimum spanning tree by prim algorithm (based on my implemented min heap), weight %v but want %v", pmw2, want)
	}

}
