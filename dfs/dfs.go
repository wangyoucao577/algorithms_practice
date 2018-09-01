// Package dfs - Depth First Search
package dfs

import "github.com/wangyoucao577/algorithms_practice/graph"

type color int

const (
	white color = iota // Not find
	gray               // Have found but not scan adjacency list
	black              // Have scaned adjacency list
)

type nodeAttr struct {
	TimestampD int          // when this node first found (i.e. when set it to gray)
	TimestampF int          // when all subnodes of this node have been searched (i.e. when set it to black)
	Color      color        // record node status during search
	Parent     graph.NodeID // remember parent node, InvalidNodeID means no parent
}
type nodeAttrArray map[graph.NodeID]*nodeAttr // nodeID is not a int start from 0, so we use `map` instread of `array`

type dfsTree struct {
	Root graph.NodeID // DFS start node, i.e. root of a tree
}

// Dfs defined a structure to store result after DFS search
type Dfs struct {
	Time      int           //as timestamp during DFS, should be a global var during DFS
	Forest    []dfsTree     //generated forest by DFS
	NodesAttr nodeAttrArray // store depth/parent/timestamp during DFS
}

// NewDfs execute the DFS search on a graph
func NewDfs(g graph.Graph) (*Dfs, error) {

	// Initialize
	dfsContext := &Dfs{0, []dfsTree{}, nodeAttrArray{}}
	g.IterateAllNodes(func(k graph.NodeID) {
		dfsContext.NodesAttr[k] = &nodeAttr{0, 0, white, graph.InvalidNodeID} // create node attr for each node
	})

	// DFS
	g.IterateAllNodes(func(k graph.NodeID) {
		if dfsContext.NodesAttr[k].Color == white {
			dfsContext.Forest = append(dfsContext.Forest, dfsTree{k}) //record a tree's root

			// execute one tree search
			dfsContext.dfsVisit(g, k)
		}
	})
	return dfsContext, nil
}

func (d *Dfs) dfsVisit(g graph.Graph, currNode graph.NodeID) {
	d.Time++
	d.NodesAttr[currNode].Color = gray
	d.NodesAttr[currNode].TimestampD = d.Time

	g.IterateAdjacencyNodes(currNode, func(v graph.NodeID) {
		if d.NodesAttr[v].Color == white {
			d.NodesAttr[v].Parent = currNode
			d.dfsVisit(g, v)
		}
	})

	d.Time++
	d.NodesAttr[currNode].Color = black
	d.NodesAttr[currNode].TimestampF = d.Time
}
