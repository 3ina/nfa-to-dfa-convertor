package main

import (
	"fmt"
	"github.com/rivo/tview"
	"strings"
)

func SummeryFlexInit(pages *tview.Pages,
	app *tview.Application,
	states,
	alphabet,
	transitions,
	startState,
	finalStates *string) *tview.Flex {
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
			*states, *alphabet, *transitions, *startState, *finalStates))
		pages.SwitchToPage("Summary")
	})
	return summaryLayout
}

func FinalStatesFormInit(pages *tview.Pages, states *[]string, finalStates *string, app *tview.Application) *tview.Form {
	formFinal := tview.NewForm().
		AddInputField("Final States (comma-separated)", "", 30, nil, func(text string) {
			*finalStates = text
		}).
		AddButton("Next", func() {
			if *finalStates == "" {
				showError(pages, "Please enter at least one final state.")
				return
			}
			if checkDuplicates(*finalStates) {
				showError(pages, "Duplicate final states found. Please ensure all final states are unique.")
				return
			}
			for _, s := range strings.Split(*finalStates, ",") {
				if !isValidState(s, *states) {
					showError(pages, fmt.Sprintf("Invalid state: %s", s))
					return
				}
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

	return formFinal
}

func StartStateFormInit(pages *tview.Pages, states *[]string, startState *string, app *tview.Application) *tview.Form {
	formStartState := tview.NewForm().
		AddInputField("Start State", "", 20, nil, func(text string) {
			*startState = text
		}).
		AddButton("Next", func() {
			if *startState == "" {
				showError(pages, "Please enter a start state.")
				return
			}
			if !isValidState(strings.TrimSpace(*startState), *states) {
				showError(pages, fmt.Sprintf("Invalid state: %s", *startState))
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

	formStartState.SetBorder(true).
		SetTitle("Step 4: Define Start State").
		SetTitleAlign(tview.AlignLeft)

	return formStartState
}

func TransitionsFormInit(pages *tview.Pages, alphabet *[]string, states *[]string, transitions *string, app *tview.Application) *tview.Form {

	formTransitions := tview.NewForm().
		AddInputField("Transitions (e.g., S0,a->S1; S0,ε->S2)", "", 50, nil, func(text string) {
			if !validateTransitionFormat(text) {
				showError(pages, "Invalid format. Transitions must follow the pattern S0,a->S1; S0,ε->S2.")
				return
			}
			*transitions = text
		}).
		AddButton("Next", func() {
			if *transitions == "" {
				showError(pages, "Please enter at least one transition.")
				return
			}
			if hasDuplicateTransitions(*transitions) {
				showError(pages, "Duplicate transitions found. Please ensure all transitions are unique.")
				return
			}

			for _, transition := range strings.Split(*transitions, ";") {
				parts := strings.Split(transition, "->")
				if len(parts) != 2 {
					showError(pages, "Invalid transition format.")
					return
				}
				fromStateInput := strings.Split(parts[0], ",")
				fromState, input := strings.TrimSpace(fromStateInput[0]), strings.TrimSpace(fromStateInput[1])

				if !isValidState(fromState, *states) {
					showError(pages, fmt.Sprintf("Invalid state in transition: %s", fromState))
					showError(pages, fmt.Sprintf("Invalid state arr in transition: %s", *states))

					return
				}
				if !isValidAlphabet(input, *alphabet) {
					showError(pages, fmt.Sprintf("Invalid alphabet symbol in transition: %s", input))
					return
				}

				toState := strings.TrimSpace(parts[1])
				if !isValidState(toState, *states) {
					showError(pages, fmt.Sprintf("Invalid target state in transition: %s", toState))
					return
				}
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

	return formTransitions
}

func AlphabetFormInit(pages *tview.Pages, alphabet *string, alphabetArr *[]string, app *tview.Application) *tview.Form {

	formAlphabet := tview.NewForm().
		AddInputField("Alphabet (comma-separated, e.g., a,b)", "", 30, nil, func(text string) {
			if !validateCommaSeparatedAlphabet(text) {
				showError(pages, "Invalid format. Alphabet must be single letters separated by commas (e.g., a,b).")
				return
			}
			*alphabet = text
			*alphabetArr = strings.Split(text, ",")
		}).
		AddButton("Next", func() {
			if *alphabet == "" {
				showError(pages, "Please enter at least one symbol in the alphabet.")
				return
			}
			if checkDuplicates(*alphabet) {
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

	return formAlphabet
}

func StatesFormInit(pages *tview.Pages, states *string, stateArr *[]string, app *tview.Application) *tview.Form {

	formStates := tview.NewForm().
		AddInputField("States (comma-seprated)",
			"", 30, nil, func(text string) {
				if !validateCommaSeparatedWords(text) {
					showError(pages, "Invalid format. States must be comma-separated alphanumeric values.")
					return
				}
				*states = text
				*stateArr = strings.Split(text, ",")
			}).
		AddButton("Next", func() {
			if *states == "" {
				showError(pages, "Please enter at least one state.")
				return
			}
			if checkDuplicates(*states) {
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

	return formStates
}
