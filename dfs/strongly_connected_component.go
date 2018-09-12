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

// SplitToStronglyConnectedComponents to split the input directed graph
// to several strongly connected components
func SplitToStronglyConnectedComponents(g graph.Graph) ([]StronglyConnectedComponent, error) {
	if !g.Directed() {
		return nil, fmt.Errorf("It's not a directed graph")
	}

	firstDfs, err := NewDfsForest(g, StackBased)
	if err != nil {
		return nil, err
	}

	// sort nodes by increasing timestampF
	sortedNodes, err := sortNodesByTimestampF(firstDfs.nodesAttr)
	if err != nil {
		return nil, err
	}

	transposedG, err := graph.Transpose(g, graph.NewAdjacencyListGraph) // calculate transposed graph of original graph
	if err != nil {
		return nil, err
	}

	// initialize for second DFS on transposed graph
	secondDfs := &Dfs{0, []dfsTree{}, nodeAttrArray{}, edgeAttrArray{}}
	secondDfs.initialize(transposedG)

	secondRootIndicators := make([]bool, transposedG.NodeCount(), transposedG.NodeCount())
	for i := len(sortedNodes) - 1; i >= 0; i-- { //second iteration by reverse order of topoSort
		currNode := sortedNodes[i]
		if secondDfs.nodesAttr[currNode].nodeColor == white {
			secondDfs.forest = append(secondDfs.forest, dfsTree{currNode})
			secondRootIndicators[currNode] = true

			secondDfs.dfsStackBasedVisit(transposedG, currNode, nil)
		}
	}

	// retrieve components split by root of forest
	secondSortedNodes, err := sortNodesByTimestampF(secondDfs.nodesAttr)
	if err != nil {
		return nil, err
	}

	components := []StronglyConnectedComponent{}

	scc := StronglyConnectedComponent{}
	for _, v := range secondSortedNodes {
		scc = append(scc, v)

		// root will always have biggeest timestampF
		// so the last node MUST be root, and will be append into components
		if secondRootIndicators[v] {
			components = append(components, scc)
			scc = StronglyConnectedComponent{} //clear
		}
	}

	return components, nil
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
