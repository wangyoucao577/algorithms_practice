package levelgraph

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

// Level defines level unit for a level graph
type Level uint

// LevelGraph defines a level graph which can be constructed by BFS
type LevelGraph struct {
	graph.Graph

	nodesLevel []Level
}

// NodeLevel to return level value of a node
func (lg LevelGraph) NodeLevel(node graph.NodeID) Level {
	return lg.nodesLevel[node]
}

// SetNodeLevel set a level for a node
func (lg LevelGraph) SetNodeLevel(node graph.NodeID, level Level) error {
	if !lg.IsNodeValid(node) {
		return fmt.Errorf("node %v is invalid", node)
	}

	lg.nodesLevel[node] = level
	return nil
}

// NewLevelGraph create a new level graph with nodes, but with only empty levels
func NewLevelGraph(g graph.Graph) *LevelGraph {

	return &LevelGraph{
		graph.NewAdjacencyListGraph(g.NodeCount(), g.Directed()),
		make([]Level, g.NodeCount(), g.NodeCount()),
	}
}
