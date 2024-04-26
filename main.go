package main

import (
	"fmt"
)

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

	p := Processor{currNode: &node4, prevNode: &Node{}, word: startWord}
	p.ProcessWord()
	fmt.Printf("Your final word: %v\n", p.Word())
	fmt.Printf("Correct word: %v", endWord)
}
