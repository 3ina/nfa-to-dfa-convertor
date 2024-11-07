package main

import (
	"fmt"
	"github.com/rivo/tview"
	"strings"
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
	formStates := tview.NewForm().
		AddInputField("States (comma-seprated)",
			"", 30, nil, func(text string) {
				states = text
			}).
		AddButton("Next", func() {
			if states == "" {
				showError(pages, "Please enter at least one state.")
				return
			}
			if checkDuplicates(states) {
				showError(pages, "Duplicate states found. Please ensure all states are unique.")
				return
			}
			pages.SwitchToPage("Alphabet")
		}).
		AddButton("Cancel", func() {
			app.Stop()
		})

	formStates.SetBorder(true).
		SetTitle("Step 1: Define States").
		SetTitleAlign(tview.AlignLeft)

	//page 2 for alphabet input
	formAlphabet := tview.NewForm().
		AddInputField("Alphabet (comma-separated, e.g., a,b)", "", 30, nil, func(text string) {
			alphabet = text
		}).
		AddButton("Next", func() {
			if alphabet == "" {
				showError(pages, "Please enter at least one symbol in the alphabet.")
				return
			}
			if checkDuplicates(alphabet) {
				showError(pages, "Duplicate symbols in the alphabet. Please ensure all symbols are unique.")
				return
			}
			pages.SwitchToPage("Transitions")
		}).
		AddButton("Back", func() {
			pages.SwitchToPage("States")
		}).
		AddButton("Cancel", func() {
			app.Stop()
		})

	formAlphabet.
		SetBorder(true).
		SetTitle("Step 2: Define Alphabet").
		SetTitleAlign(tview.AlignLeft)

	//Page 3: Transition Input
	formTransitions := tview.NewForm().
		AddInputField("Transitions (e.g., S0,a->S1; S0,ε->S2)", "", 50, nil, func(text string) {
			transitions = text
		}).
		AddButton("Next", func() {
			if transitions == "" {
				showError(pages, "Please enter at least one transition.")
				return
			}
			if hasDuplicateTransitions(transitions) {
				showError(pages, "Duplicate transitions found. Please ensure all transitions are unique.")
				return
			}
			pages.SwitchToPage("StartState")
		}).
		AddButton("Back", func() {
			pages.SwitchToPage("Alphabet")
		}).
		AddButton("Cancel", func() {
			app.Stop()
		})
	formTransitions.
		SetBorder(true).
		SetTitle("Step 3: Define Transitions").
		SetTitleAlign(tview.AlignLeft)

	// page 4 for start state

	formStart := tview.NewForm().
		AddInputField("Start State", "", 20, nil, func(text string) {
			startState = text
		}).
		AddButton("Next", func() {
			if startState == "" {
				showError(pages, "Please enter a start state.")
				return
			}
			pages.SwitchToPage("FinalStates")
		}).
		AddButton("Back", func() {
			pages.SwitchToPage("Transitions")
		}).
		AddButton("Cancel", func() {
			app.Stop()
		})

	formStart.SetBorder(true).
		SetTitle("Step 4: Define Start State").
		SetTitleAlign(tview.AlignLeft)

	// page 5 for Final States

	formFinal := tview.NewForm().
		AddInputField("Final States (comma-separated)", "", 30, nil, func(text string) {
			finalStates = text
		}).
		AddButton("Next", func() {
			if finalStates == "" {
				showError(pages, "Please enter at least one final state.")
				return
			}
			if checkDuplicates(finalStates) {
				showError(pages, "Duplicate final states found. Please ensure all final states are unique.")
				return
			}
			pages.SwitchToPage("Summary")
		}).
		AddButton("Back", func() {
			pages.SwitchToPage("StartState")
		}).
		AddButton("Cancel", func() {
			app.Stop()
		})

	formFinal.
		SetBorder(true).
		SetTitle("Step 5: Define Final States").
		SetTitleAlign(tview.AlignLeft)

	// Page 6 summary about nfa and

	summary := tview.
		NewTextView().
		SetDynamicColors(true).
		SetWrap(true)

	summary.
		SetBorder(true).
		SetTitle("NFA Summary").
		SetTitleAlign(tview.AlignLeft)

	confirmForm := tview.NewForm().
		AddButton("Confirm", func() {
			fmt.Println("NFA Details:")
			fmt.Printf("States: %s\nAlphabet: %s\nTransitions: %s\nStart State: %s\nFinal States: %s\n",
				states, alphabet, transitions, startState, finalStates)
			app.Stop() // End the app after submission
		}).
		AddButton("Back", func() {
			pages.SwitchToPage("FinalStates")
		}).
		AddButton("Cancel", func() {
			app.Stop()
		})

	// Layout for Summary Page
	summaryLayout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(summary, 0, 1, false).
		AddItem(confirmForm, 3, 1, true)

	confirmForm.AddButton("Show Summary", func() {
		summary.SetText(fmt.Sprintf("[yellow]States:[-] %s\n[cyan]Alphabet:[-] %s\n[green]Transitions:[-] %s\n[red]Start State:[-] %s\n[blue]Final States:[-] %s",
			states, alphabet, transitions, startState, finalStates))
		pages.SwitchToPage("Summary")
	})

	pages.AddPage("States", formStates, true, true)
	pages.AddPage("Alphabet", formAlphabet, true, false)
	pages.AddPage("Transitions", formTransitions, true, false)
	pages.AddPage("StartState", formStart, true, false)
	pages.AddPage("FinalStates", formFinal, true, false)
	pages.AddPage("Summary", summaryLayout, true, false)

	// Run the application
	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}

}

func showError(pages *tview.Pages, message string) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.RemovePage("modal")
		})
	pages.AddPage("modal", modal, true, true)
}

func checkDuplicates(input string) bool {
	seen := make(map[string]bool)
	for _, item := range strings.Split(input, ",") {
		item = strings.TrimSpace(item)
		if seen[item] {
			return true
		}
		seen[item] = true
	}
	return false
}

func hasDuplicateTransitions(input string) bool {
	seen := make(map[string]bool)
	for _, transition := range strings.Split(input, ";") {
		transition = strings.TrimSpace(transition)
		if seen[transition] {
			return true
		}
		seen[transition] = true
	}
	return false
}
