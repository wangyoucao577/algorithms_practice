package dfs

import (
	"fmt"
	"sort"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

// StronglyConnectedComponent represent a strongly connected component,
// include all vertexs within
type StronglyConnectedComponent []graph.NodeID

// For Sort Interfaces
func (s StronglyConnectedComponent) Len() int { return len(s) }
func (s StronglyConnectedComponent) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s StronglyConnectedComponent) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type dfsSCC struct {
	Dfs
	scc StronglyConnectedComponent
}

// SplitToStronglyConnectedComponents to split the input graph
// to several strongly connected components
func SplitToStronglyConnectedComponents(g graph.Graph) ([]StronglyConnectedComponent, error) {
	components := []StronglyConnectedComponent{}

	firstDfs, err := NewDfsForest(g, StackBased)
	if err != nil {
		return components, err
	}

	// sort nodes by increasing timestampF
	sortedNodes, err := sortNodesByTimestampF(firstDfs.nodesAttr)
	if err != nil {
		return components, err
	}

	transposedG := transposeGraph(g) // calculate transposed graph of original graph

	// initialize for second DFS on transposed graph
	secondDfsContext := dfsSCC{
		Dfs{0, []dfsTree{}, nodeAttrArray{}, edgeAttrArray{}},
		StronglyConnectedComponent{}}
	secondDfsContext.initialize(transposedG)

	for i := len(sortedNodes) - 1; i >= 0; i-- { //second iteration by reverse order of topoSort
		currNode := sortedNodes[i]
		if secondDfsContext.nodesAttr[currNode].nodeColor == white {
			secondDfsContext.forest = append(secondDfsContext.forest, dfsTree{currNode})
			secondDfsContext.scc = StronglyConnectedComponent{} // clear before each dfs

			secondDfsContext.stackBasedVisitForSCC(transposedG, currNode)
			components = append(components, secondDfsContext.scc)
		}
	}

	return components, nil
}

// generate a new graph based on current one, but reverse all edges
func transposeGraph(g graph.Graph) graph.Graph {

	newGraph := graph.NewAdjacencyListGraph(g.NodeCount())

	g.IterateAllNodes(func(u graph.NodeID) {
		g.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			// means u->v exist in current graph
			newGraph.AddEdge(v, u)
		})
	})

	return newGraph
}

// default increasing
func sortNodesByTimestampF(nodesAttr nodeAttrArray) ([]graph.NodeID, error) {

	nodesCount := len(nodesAttr)
	if nodesCount <= 0 {
		return nil, fmt.Errorf("Empty nodes attr array")
	}

	timestampFArray := make(nodeWithTimestampFArray, 0, nodesCount) //reserve array
	for k, v := range nodesAttr {
		if v == nil {
			return nil, fmt.Errorf("null pointer of node %v attr. dfs not exeucted?", k)
		}
		timestampFArray = append(timestampFArray, nodeWithTimestampF{k, v.timestampF})
	}

	sort.Sort(timestampFArray) //sort
	//fmt.Println(timestampFArray)

	sortedNodes := make([]graph.NodeID, 0, nodesCount) //reserve array
	for _, v := range timestampFArray {
		sortedNodes = append(sortedNodes, v.nodeID)
	}

	return sortedNodes, nil
}

func (d *dfsSCC) stackBasedVisitForSCC(g graph.Graph, root graph.NodeID) {
	d.time++
	d.nodesAttr[root].nodeColor = gray
	d.nodesAttr[root].timestampD = d.time

	var stack = []graph.NodeID{}
	stack = append(stack, root)

	for len(stack) > 0 {
		currNode := stack[len(stack)-1]

		newWhiteNodeFound := false
		g.ControllableIterateAdjacencyNodes(currNode, func(v graph.NodeID) graph.IterateControl {
			if d.nodesAttr[v].nodeColor == white {
				newWhiteNodeFound = true

				d.nodesAttr[v].parent = currNode

				d.time++
				d.nodesAttr[v].nodeColor = gray
				d.nodesAttr[v].timestampD = d.time

				stack = append(stack, v) // push stack: push to the end

				return graph.BreakIterate
			}

			return graph.ContinueIterate
		})

		if !newWhiteNodeFound {
			d.time++
			d.nodesAttr[currNode].nodeColor = black
			d.nodesAttr[currNode].timestampF = d.time

			d.scc = append(d.scc, currNode) //remember for current strongly connect component

			stack = stack[:len(stack)-1] // pop from stack
		}
	}
}

type nodeWithTimestampF struct {
	nodeID     graph.NodeID
	timestampF int
}

type nodeWithTimestampFArray []nodeWithTimestampF

func (n nodeWithTimestampFArray) Len() int { return len(n) }
func (n nodeWithTimestampFArray) Less(i, j int) bool {
	return n[i].timestampF < n[j].timestampF
}
func (n nodeWithTimestampFArray) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
