package main

import "slices"

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
