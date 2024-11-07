package main

import (
	"github.com/rivo/tview"
)

func main() {
	//s0 := &automata.State{Name: "S0"}
	//s1 := &automata.State{Name: "S1"}
	//s2 := &automata.State{Name: "S2"}
	//s3 := &automata.State{Name: "S3"}
	//s4 := &automata.State{Name: "S4"}
	//s5 := &automata.State{Name: "S5"}
	//s6 := &automata.State{Name: "S6"}
	//s7 := &automata.State{Name: "S7"}
	//s8 := &automata.State{Name: "S8"}
	//
	//nfa := automata.NFA{
	//	States:   []*automata.State{s0, s1, s2, s3, s4, s5, s6, s7, s8},
	//	Alphabet: []rune{'a', 'b'},
	//	Transitions: []automata.Transition{
	//		{From: s0, Input: 'ε', To: []*automata.State{s1, s7}},
	//		{From: s1, Input: 'ε', To: []*automata.State{s2, s4}},
	//		{From: s2, Input: 'a', To: []*automata.State{s3}},
	//		{From: s4, Input: 'b', To: []*automata.State{s5}},
	//		{From: s5, Input: 'ε', To: []*automata.State{s6}},
	//		{From: s3, Input: 'ε', To: []*automata.State{s6}},
	//		{From: s6, Input: 'ε', To: []*automata.State{s1, s7}},
	//		{From: s7, Input: 'a', To: []*automata.State{s1, s8}},
	//	},
	//	StartState:  s0,
	//	FinalStates: []*automata.State{s8},
	//}
	//
	//dfa := nfa.ConvertToDfa()
	//
	//dfa.Print()

	app := tview.NewApplication()
	var states, alphabet, transitions, startState, finalStates string

	pages := tview.NewPages()

	//page 1 for state input
	formStates := StatesFormInit(pages, &states, app)
	//page 2 for alphabet input
	formAlphabet := AlphabetFormInit(pages, &alphabet, app)
	//Page 3: Transition Input
	formTransitions := TransitionsFormInit(pages, &transitions, app)
	// page 4 for start state
	formStartState := StartStateFormInit(pages, &startState, app)
	// page 5 for Final States

	formFinalStates := FinalStatesFormInit(pages, &finalStates, app)

	// Page 6 summary about nfa and

	summaryLayout := SummeryFlexInit(pages, app, states, alphabet,
		transitions, states, finalStates)
	pages.AddPage("States", formStates, true, true)
	pages.AddPage("Alphabet", formAlphabet, true, false)
	pages.AddPage("Transitions", formTransitions, true, false)
	pages.AddPage("StartState", formStartState, true, false)
	pages.AddPage("FinalStates", formFinalStates, true, false)
	pages.AddPage("Summary", summaryLayout, true, false)

	// Run the application
	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}

}
