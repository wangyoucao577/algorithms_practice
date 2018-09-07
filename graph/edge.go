package graph

// EdgeID represent an edge between two nodes.
// If it's an undirected edge, From and To can be swapped.
type EdgeID struct {
	From NodeID
	To   NodeID
}

//UndirectedEqual to check whether two undirected edges are equal
func (e EdgeID) UndirectedEqual(f EdgeID) bool {
	if e == f {
		return true
	}
	if e.From == f.To && e.To == f.From {
		return true
	}
	return false
}

//Reverse return reverse edge of current one
func (e EdgeID) Reverse() EdgeID {
	return EdgeID{e.To, e.From}
}

//IsValid return whether the edgeID is valid
//if From or To is InvalidNodeID, the edgeID is invalid
func (e EdgeID) IsValid() bool {
	return e.From != InvalidNodeID && e.To != InvalidNodeID
}
