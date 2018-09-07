package graph

// NodeID represent each node by `unsigned int`, start from 0
type NodeID uint

const (
	// InvalidNodeID defined invalid value of NodeID
	InvalidNodeID = NodeID(^uint(0))
)
