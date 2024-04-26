package main

import (
	"fmt"
	"slices"
)

type Graph struct {
	currNode *Node
	prevNode *Node
	word     []string
}

func (g *Graph) ProcessWord() {
	g.currNode.Visited = true
	g.word = g.currNode.Manipulator.Apply(g.word)
	unvisitedNbrs := g.currNode.getUnvisitedNeighbors()

	if g.prevNode != nil {
		prevNodeUnvisited := g.prevNode.getUnvisitedNeighbors()
		if len(prevNodeUnvisited) > 0 {
			g.currNode = g.prevNode
			g.prevNode = nil
			g.ProcessWord()
		}
	}

	if len(unvisitedNbrs) == 0 {
		return
	}

	g.prevNode = g.currNode
	g.currNode = unvisitedNbrs[0]
	g.ProcessWord()
}

func (g *Graph) Word() []string {
	return g.word
}

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

type Manipulator interface {
	Apply([]string) []string
}

type Reverser struct{}

func (r *Reverser) Apply(word []string) []string {
	slices.Reverse(word)
	return word
}

type LetterSwapper struct {
	letter string
	pos    int
}

func (ls *LetterSwapper) Apply(word []string) []string {
	word[ls.pos] = ls.letter
	return word
}

type PositionSwapper struct {
	pos1 int
	pos2 int
}

func (ps *PositionSwapper) Apply(word []string) []string {
	word[ps.pos1], word[ps.pos2] = word[ps.pos2], word[ps.pos1]
	return word
}

func main() {
	node1 := createNode(1, &Reverser{})
	node2 := createNode(2, &Reverser{})
	node3 := createNode(3, &LetterSwapper{letter: "y", pos: 0})
	node4 := createNode(4, &LetterSwapper{letter: "z", pos: 3})
	node5 := createNode(5, &LetterSwapper{letter: "i", pos: 2})
	node6 := createNode(6, &PositionSwapper{pos1: 1, pos2: 2})

	addNeighbor(&node4, &node1)
	addNeighbor(&node4, &node5)
	addNeighbor(&node5, &node2)
	addNeighbor(&node2, &node3)
	addNeighbor(&node2, &node6)

	startWord := []string{"c", "r", "e", "s", "t"}
	endWord := []string{"t", "i", "z", "z", "y"}

	g := Graph{currNode: &node4, prevNode: &Node{}, word: startWord}
	g.ProcessWord()
	fmt.Printf("Your final word: %v\n", g.Word())
	fmt.Printf("Correct word: %v", endWord)

}

func addNeighbor(n1 *Node, n2 *Node) {
	n1.Neighbors = append(n1.Neighbors, n2)
	n2.Neighbors = append(n2.Neighbors, n1)
}
