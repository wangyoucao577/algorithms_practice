package dfs

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

type color int

const (
	white color = iota // Not find
	gray               // Have found but not scan adjacency list
	black              // Have scaned adjacency list
)

type nodeAttr struct {
	timestampD int          // when this node first found (i.e. when set it to gray)
	timestampF int          // when all subnodes of this node have been searched (i.e. when set it to black)
	nodeColor  color        // record node status during search
	parent     graph.NodeID // remember parent node, InvalidNodeID means no parent
}
type nodeAttrArray map[graph.NodeID]*nodeAttr // nodeID is not a int start from 0, so we use `map` instread of `array`

type edgeType int

// below 4 types of edges defined in "Introduction to Algorithms" ch22.3
const (
	unknown  edgeType = iota // unknown before DFS
	branch                   // the edge is one of branches of a DFS tree
	forward                  // it's not a branch but also connected to a descendant
	backward                 // connected to a ancestor (include spin edge)
	cross                    // others, i.e. not above three types
)

type edgeAttr struct {
	t edgeType
}

func (t edgeType) String() string {
	switch t {
	case branch:
		return fmt.Sprintf("branch")
	case forward:
		return fmt.Sprintf("forward")
	case backward:
		return fmt.Sprintf("backward")
	case cross:
		return fmt.Sprintf("cross")
	default:
		return fmt.Sprintf("unknown")
	}
}

type edgeAttrArray map[graph.EdgeID]*edgeAttr
