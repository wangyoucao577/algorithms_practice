package minspanningtree

import (
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// Prim2 algorithm calculate minimum spanning tree on the input undirected graph
// Note that MinHeap in this implmenetation is implemented by myself.
func Prim2(g weightedgraph.WeightedGraph) (MinSpanningTree, error) {
	mst := MinSpanningTree{[]graph.EdgeID{}, g}

	// initialize for queryHeap
	queringHeap := &myQueryHeap{myMinBinaryHeap{}, map[graph.NodeID]*heapNode{}}
	g.IterateAllNodes(func(u graph.NodeID) {
		item := &heapNode{self: u, parent: graph.InvalidNodeID, key: maxKey}
		queringHeap.insert(item)
		queringHeap.nodeToItem[u] = item
	})
	queringHeap.decreaseKey(1, 0) // random starting node

	// iterate all nodes by decreasing key
	for queringHeap.Len() > 0 {

		uItem := queringHeap.extractMin() //pop min key node

		if uItem.parent != graph.InvalidNodeID {
			mst.edges = append(mst.edges, graph.EdgeID{From: uItem.parent, To: uItem.self})
		}

		g.IterateAdjacencyNodes(uItem.self, func(v graph.NodeID) {
			item, ok := queringHeap.find(v)
			uvWeight, _ := g.Weight(uItem.self, v)
			if ok && uvWeight < item.key {
				queringHeap.decreaseKey(item.index, uvWeight)
			}
		})
	}

	return mst, nil
}
