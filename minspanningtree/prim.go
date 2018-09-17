package minspanningtree

import (
	"sort"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// Prim algorithm calculate minimum spanning tree on the input undirected graph
func Prim(g weightedgraph.WeightedGraph) (MinSpanningTree, error) {

	mst := MinSpanningTree{[]graph.EdgeID{}, g}

	priorityQueue := nodeItemsArray{}
	g.IterateAllNodes(func(u graph.NodeID) {
		priorityQueue = append(priorityQueue, &nodeItem{self: u, parent: graph.InvalidNodeID, key: maxKey})
	})
	priorityQueue[0].key = 0 // starting node

	//assign key for each node
	for len(priorityQueue) > 0 {

		sort.Sort(priorityQueue)

		u := priorityQueue[0]
		priorityQueue = priorityQueue[1:] //pop min node

		if u.parent != graph.InvalidNodeID {
			mst.edges = append(mst.edges, graph.EdgeID{From: u.parent, To: u.self})
		}

		g.IterateAdjacencyNodes(u.self, func(v graph.NodeID) {
			item, ok := priorityQueue.find(v)
			uvWeight, _ := g.Weight(u.self, v)
			if ok && uvWeight < item.key {
				item.parent = u.self
				item.key = uvWeight
			}
		})
	}

	return mst, nil
}

const (
	maxKey weightedgraph.Weight = weightedgraph.Weight((^uint(0)) >> 1)
)

type nodeItem struct {
	self   graph.NodeID
	parent graph.NodeID

	key weightedgraph.Weight
}
type nodeItemsArray []*nodeItem

func (n nodeItemsArray) Len() int { return len(n) }
func (n nodeItemsArray) Less(i, j int) bool {
	return n[i].key < n[j].key
}
func (n nodeItemsArray) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n nodeItemsArray) find(v graph.NodeID) (*nodeItem, bool) {
	for _, vItem := range n {
		if vItem.self == v {
			return vItem, true
		}
	}
	return nil, false
}
