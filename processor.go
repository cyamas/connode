package main

type Processor struct {
	currNode *Node
	prevNode *Node
	word     []string
}

func (p *Processor) ProcessWord() {
	p.currNode.Visited = true
	p.word = p.currNode.Manipulator.Apply(p.word)
	unvisitedNbrs := p.currNode.getUnvisitedNeighbors()

	if p.prevNode != nil {
		prevNodeUnvisited := p.prevNode.getUnvisitedNeighbors()
		if len(prevNodeUnvisited) > 0 {
			p.currNode = p.prevNode
			p.prevNode = nil
			p.ProcessWord()
		}
	}

	if len(unvisitedNbrs) == 0 {
		return
	}

	p.prevNode = p.currNode
	p.currNode = unvisitedNbrs[0]
	p.ProcessWord()
}

func (p *Processor) Word() []string {
	return p.word
}
