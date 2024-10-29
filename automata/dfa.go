package automata

import "fmt"

type DFA struct {
	States      []*State
	Alphabet    []rune
	Transitions map[string]map[rune]string
	StartState  *State
	FinalStates []*State
}

func (dfa *DFA) Print() {
	fmt.Println("DFA States:")
	for _, state := range dfa.States {
		fmt.Println(state.Name)
	}
	fmt.Println("\nDFA Transitions:")

	for from, trans := range dfa.Transitions {
		for input, to := range trans {
			fmt.Printf("From %s on '%c' -> %s\n", from, input, to)
		}
	}

	fmt.Println("\nDFA Start State:")
	fmt.Println(dfa.StartState.Name)
	fmt.Println("\nDFA Final States:")
	for _, state := range dfa.FinalStates {
		fmt.Println(state.Name)
	}
}
