package main

type Node struct {
	ID          int
	Neighbors   []*Node
	Visited     bool
	Manipulator Manipulator
	IsStart     bool
}

func createNode(id int, manipulator Manipulator) Node {
	return Node{
		ID:          id,
		Visited:     false,
		Manipulator: manipulator,
	}
}

func (n *Node) getUnvisitedNeighbors() []*Node {
	var nbrs []*Node
	for _, nbr := range n.Neighbors {
		if nbr.Visited == false {
			nbrs = append(nbrs, nbr)
		}
	}
	return nbrs
}

func addNeighbor(n1 *Node, n2 *Node) {
	n1.Neighbors = append(n1.Neighbors, n2)
	n2.Neighbors = append(n2.Neighbors, n1)
}
