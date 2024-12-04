package main

import (
	"fmt"
	"github.com/3ina/nfa-to-dfa-convertor/automata"
	"github.com/rivo/tview"
)

func main() {

	//ε

	app := tview.NewApplication().EnableMouse(true).EnablePaste(true)
	var states, alphabet, transitions, startState, finalStates string
	nfaInput := &automata.NFA{
		States:      nil,
		Alphabet:    nil,
		Transitions: nil,
		StartState:  nil,
		FinalStates: nil,
	}
	dfaResult := &automata.DFA{
		States:      nil,
		Alphabet:    nil,
		Transitions: nil,
		StartState:  nil,
		FinalStates: nil,
	}
	var statesArr []string
	var alphabetArr []string
	pages := tview.NewPages()

	//page 1 for state input
	formStates := StatesFormInit(pages, &states, &statesArr, app)
	//page 2 for alphabet input
	formAlphabet := AlphabetFormInit(pages, &alphabet, &alphabetArr, app)
	//Page 3: Transition Input
	formTransitions := TransitionsFormInit(pages, &alphabetArr, &statesArr, &transitions, app)
	// page 4 for start state
	formStartState := StartStateFormInit(pages, &statesArr, &startState, app)
	// page 5 for Final States
	formFinalStates := FinalStatesFormInit(pages, &statesArr, &finalStates, app)
	// Page 6 summary about nfa and
	summaryLayout := SummeryFlexInit(pages, app, &states, &alphabet,
		&transitions, &states, &finalStates, nfaInput, dfaResult)

	pages.AddPage("States", formStates, true, true)
	pages.AddPage("Alphabet", formAlphabet, true, false)

	pages.AddPage("Transitions", formTransitions, true, false)
	pages.AddPage("StartState", formStartState, true, false)
	pages.AddPage("FinalStates", formFinalStates, true, false)
	pages.AddPage("Summary", summaryLayout, true, false)

	// Run the application
	if err := app.SetRoot(pages, true).Run(); err != nil {
		fmt.Printf(err.Error())
	}

}
