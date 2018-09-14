package minspanningtree

import (
	"sort"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// Kruskal calculate minimum spanning tree on the input undirected graph
func Kruskal(g weightedgraph.WeightedGraph) (MinSpanningTree, error) {

	mst := MinSpanningTree{[]graph.EdgeID{}, g}

	// collect all edges with weight
	edges := edgeWithWeightArray{}
	g.IterateAllNodes(func(u graph.NodeID) {
		g.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			weight, _ := g.Weight(u, v) //TODO: check err if necessary
			edge := graph.EdgeID{From: u, To: v}
			if !edges.Exist(edge) { // filter equal undirected edge
				edges = append(edges, edgeWithWeight{edge, weight})
			}
		})
	})

	sort.Sort(edges) // sort by weight increasing

	//MAKE-SETS
	//TODO: should be a better way to implement the sets
	nodesSets := sets{map[int]map[graph.NodeID]struct{}{}, 0}

	for _, e := range edges {

		if !nodesSets.nodesInSameSet(e.From, e.To) {
			nodesSets.union(e.From, e.To)
			mst.edges = append(mst.edges, e.EdgeID)
		}
	}

	return mst, nil
}

type edgeWithWeight struct {
	graph.EdgeID
	w weightedgraph.Weight
}
type edgeWithWeightArray []edgeWithWeight

// TODO: it should be a func of graph inteface
func (e edgeWithWeightArray) Exist(edge graph.EdgeID) bool {
	for _, v := range e {
		if v.UndirectedEqual(edge) {
			return true
		}
	}
	return false
}

func (e edgeWithWeightArray) Len() int { return len(e) }
func (e edgeWithWeightArray) Less(i, j int) bool {
	return e[i].w < e[j].w
}
func (e edgeWithWeightArray) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// use map as a set, use map[int]map as many sets
type sets struct {
	s     map[int]map[graph.NodeID]struct{}
	count int
}

const (
	invalidSetID int = -1
)

func (s sets) nodesInSameSet(n1, n2 graph.NodeID) bool {
	id1 := s.find(n1)
	id2 := s.find(n2)

	if id1 == invalidSetID || id2 == invalidSetID {
		return false
	}
	return id1 == id2
}

// find set of n1 and n2, then union the two sets into one
func (s *sets) union(n1, n2 graph.NodeID) {
	id1 := s.find(n1)
	id2 := s.find(n2)

	if id1 == invalidSetID && id2 == invalidSetID {
		newSet := map[graph.NodeID]struct{}{}
		newSet[n1] = struct{}{}
		newSet[n2] = struct{}{}
		s.s[s.count] = newSet
		s.count++ // will small than g.NodeCount()
		return
	}

	if id1 == invalidSetID { // id2 != invalideSetID
		s.s[id2][n1] = struct{}{}
		return
	}

	if id2 == invalidSetID { // id1 != invalideSetID
		s.s[id1][n2] = struct{}{}
		return
	}

	//both id1 and id2 not invalid,
	// move all nodes into set1, then delete set2
	for k, v := range s.s[id2] {
		s.s[id1][k] = v
	}
	delete(s.s, id2)
}

func (s sets) find(node graph.NodeID) int {
	setID := invalidSetID
	for k, v := range s.s {
		if _, ok := v[node]; ok {
			setID = k
			break
		}
	}
	return setID
}
