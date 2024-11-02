package main

import (
	"github.com/3ina/nfa-to-dfa-convertor/automata"
)

func main() {
	s0 := &automata.State{Name: "S0"}
	s1 := &automata.State{Name: "S1"}
	s2 := &automata.State{Name: "S2"}
	s3 := &automata.State{Name: "S3"}

	nfa := automata.NFA{
		States:   []*automata.State{s0, s1, s2, s3},
		Alphabet: []rune{'a', 'b'},
		Transitions: []automata.Transition{
			{From: s0, Input: 'ε', To: []*automata.State{s1, s2}},
			{From: s1, Input: 'a', To: []*automata.State{s1}},
			{From: s1, Input: 'b', To: []*automata.State{s1, s3}},
			{From: s2, Input: 'a', To: []*automata.State{s2}},
			{From: s2, Input: 'b', To: []*automata.State{s2}},
			{From: s3, Input: 'a', To: []*automata.State{s3}},
			{From: s3, Input: 'b', To: []*automata.State{s3}},
		},
		StartState:  s0,
		FinalStates: []*automata.State{s3},
	}

	dfa := nfa.ConvertToDfa()

	dfa.Print()
}
