package minspanningtree

import (
	"container/heap"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// Prim algorithm calculate minimum spanning tree on the input undirected graph
// Note that MinHeap in this implmenetation is based on "container/heap"
func Prim(g weightedgraph.WeightedGraph) (MinSpanningTree, error) {

	mst := MinSpanningTree{[]graph.EdgeID{}, g}

	// initialize for queryHeap
	queringHeap := &queryHeap{minHeap{}, map[graph.NodeID]*heapNode{}}
	g.IterateAllNodes(func(u graph.NodeID) {
		item := &heapNode{self: u, parent: graph.InvalidNodeID, key: maxKey}
		queringHeap.minHeap = append(queringHeap.minHeap, item)
		queringHeap.nodeToItem[u] = item
	})
	queringHeap.minHeap[0].key = 0 // random starting node
	heap.Init(&queringHeap.minHeap)

	// iterate all nodes by decreasing key
	for queringHeap.Len() > 0 {

		uItem := queringHeap.popMin() //pop min key node

		if uItem.parent != graph.InvalidNodeID {
			mst.edges = append(mst.edges, graph.EdgeID{From: uItem.parent, To: uItem.self})
		}

		g.IterateAdjacencyNodes(uItem.self, func(v graph.NodeID) {
			item, ok := queringHeap.find(v)
			uvWeight, _ := g.Weight(uItem.self, v)
			if ok && uvWeight < item.key {
				queringHeap.update(item, uItem.self, uvWeight)
			}
		})
	}

	return mst, nil
}
