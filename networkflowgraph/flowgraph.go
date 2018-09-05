//Package networkflowgraph represent graph for maximum flow problem
// i.e. a adjacency list based graph with edge capacity
package networkflowgraph

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

type edgeAttr struct {
	capacity int
}
type edgeAttrArray map[graph.EdgeID]*edgeAttr

//NetworkFlowGraph represent graph for network flow problem (maximum flow problem)
type NetworkFlowGraph struct {
	AdjGraph  graph.AdjacencyListGraph
	edgesAttr edgeAttrArray
	Source    graph.NodeID
	Target    graph.NodeID
}

//ConstructNetworkFlowGraph try to construct a adjacency list based graph with edge capacity,
// nodeCount and edgeCount will define V and E
// then from r to read contents for adjacency list relationship and edge attr
func ConstructNetworkFlowGraph(nodeCount, edgeCount int, r io.Reader) (*NetworkFlowGraph, error) {
	flowGraph := &NetworkFlowGraph{
		AdjGraph:  graph.AdjacencyListGraph{},
		edgesAttr: edgeAttrArray{},
		Source:    graph.InvalidNodeID,
		Target:    graph.InvalidNodeID}

	for i := 0; i < nodeCount; i++ {
		flowGraph.AdjGraph = append(flowGraph.AdjGraph, []graph.NodeID{})
	}

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var count int
	for scanner.Scan() {
		count++

		var fromNode, toNode, edgeCapacity int
		fmt.Sscanf(scanner.Text(), "%d%d%d", &fromNode, &toNode, &edgeCapacity)

		//NOTE: input nodeID will start from 1, but in code we prefer start from 0
		fromNode--
		toNode--

		if fromNode < 0 || fromNode >= nodeCount {
			return nil, fmt.Errorf("invalid fromNode %d, nodeCount %d", fromNode, nodeCount)
		}
		if toNode < 0 || toNode >= nodeCount {
			return nil, fmt.Errorf("invalid toNode %d, nodeCount %d", toNode, nodeCount)
		}
		if edgeCapacity <= 0 {
			return nil, fmt.Errorf("invalid edgeCapacity %d", edgeCapacity)
		}

		flowGraph.AdjGraph[fromNode] = append(flowGraph.AdjGraph[fromNode], graph.NodeID(toNode))
		edge := graph.EdgeID{From: graph.NodeID(fromNode), To: graph.NodeID(toNode)}
		flowGraph.edgesAttr[edge] = &edgeAttr{edgeCapacity}

		if count >= edgeCount { // got enough edges
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	flowGraph.Source = graph.NodeID(0)
	flowGraph.Target = graph.NodeID(nodeCount - 1)
	return flowGraph, nil
}

// Capacity return capacity for an edge
func (g *NetworkFlowGraph) Capacity(e graph.EdgeID) int {
	c, ok := g.edgesAttr[e]
	if !ok {
		return 0
	}
	return c.capacity
}
