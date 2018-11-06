package shortestpaths

import (
	"container/heap"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// Dijkstra implements the dijkstra shortest path algorithm
func Dijkstra(g *weightedgraph.WeightedGraph, s graph.NodeID) *ShortestPaths {

	sp := &ShortestPaths{}
	sp.initializeSingleSource(g, s)

	h := queryHeap{priorityQueue{}, map[graph.NodeID]*pqItem{}}
	g.IterateAllNodes(func(u graph.NodeID) {
		item := &pqItem{u, sp.nodesMap[u].d, h.Len()}
		h.priorityQueue = append(h.priorityQueue, item)
		h.m[u] = item
	})
	heap.Init(&h.priorityQueue)

	for h.Len() > 0 {
		u := sp.extractMin(&h)

		g.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			w, _ := g.Weight(u, v)
			newWeight := sp.relax(u, v, w)
			if newWeight > 0 {
				// update key
				if vItem, ok := h.m[v]; ok {
					vItem.priority = newWeight
					heap.Fix(&h.priorityQueue, vItem.index)
				}
			}
		})
	}

	return sp
}

type pqItem struct {
	u        graph.NodeID         // The value of the item; arbitrary.
	priority weightedgraph.Weight // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type priorityQueue []*pqItem

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type queryHeap struct {
	priorityQueue
	m map[graph.NodeID]*pqItem
}

func (sp *ShortestPaths) extractMin(h *queryHeap) graph.NodeID {

	min := heap.Pop(h).(*pqItem)
	delete(h.m, min.u) // also remove from the map
	return min.u
}
