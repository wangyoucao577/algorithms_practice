package graph

// Path represented by nodes
type Path []NodeID

//Equal to compare whether current Path equal to another one
func (p Path) Equal(q Path) bool {
	//NOTE: could use reflect.DeepEqual() instead.
	// but we implement it manually to avoid `import reflect`.
	if len(p) != len(q) {
		return false
	}

	if (p == nil) != (q == nil) {
		return false
	}

	for i := range p {
		if p[i] != q[i] {
			return false
		}
	}

	return true
}
