//Package flownetwork represent flow network (directed graph + capacities) for maximum flow problem
// also defined source and target for the maximum flow problem
package flownetwork

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

// FlowUnit represent unit for capacity, flow
type FlowUnit int

type capacityStorage map[graph.EdgeID]FlowUnit

//FlowNetwork represent graph for network flow problem (maximum flow problem)
type FlowNetwork struct {
	adjGraph   graph.AdjacencyListGraph
	capacities capacityStorage

	// source, target will represent the maximum flow problem on the flow network
	// so we define them with the flow network
	source graph.NodeID
	target graph.NodeID
}

//ConstructFlowNetwork try to construct a adjacency list based graph with edge capacity,
// nodeCount and edgeCount will define V and E
// then from r to read contents for adjacency list relationship and edge attr
func ConstructFlowNetwork(nodeCount, edgeCount int, r io.Reader) (*FlowNetwork, error) {
	flowGraph := &FlowNetwork{
		adjGraph:   graph.AdjacencyListGraph{},
		capacities: capacityStorage{},
		source:     graph.InvalidNodeID,
		target:     graph.InvalidNodeID}

	for i := 0; i < nodeCount; i++ {
		flowGraph.adjGraph = append(flowGraph.adjGraph, []graph.NodeID{})
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

		flowGraph.adjGraph[fromNode] = append(flowGraph.adjGraph[fromNode], graph.NodeID(toNode))
		edge := graph.EdgeID{From: graph.NodeID(fromNode), To: graph.NodeID(toNode)}
		flowGraph.capacities[edge] = FlowUnit(edgeCapacity)

		if count >= edgeCount { // got enough edges
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	flowGraph.source = graph.NodeID(0)
	flowGraph.target = graph.NodeID(nodeCount - 1)
	return flowGraph, nil
}

// Capacity return capacity for an edge
func (f *FlowNetwork) Capacity(e graph.EdgeID) FlowUnit {
	c, ok := f.capacities[e]
	if !ok {
		return 0
	}
	return c
}

// Graph return adjcency graph (list or matrix based)
func (f *FlowNetwork) Graph() graph.Graph {
	return f.adjGraph
}

// Source represent the start point of the maximum flow problem on current flow network
func (f *FlowNetwork) Source() graph.NodeID {
	return f.source
}

// Target represent the end point of the maximum flow problem on current flow network
func (f *FlowNetwork) Target() graph.NodeID {
	return f.target
}
