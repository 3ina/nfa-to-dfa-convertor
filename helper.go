package main

import (
	"fmt"
	"github.com/3ina/nfa-to-dfa-convertor/automata"
	"github.com/rivo/tview"
	"regexp"
	"strings"
)

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

func isValidState(state string, states []string) bool {
	for _, s := range states {
		if s == state {
			return true
		}
	}
	return false
}

func isValidAlphabet(symbol string, alphabet []string) bool {
	for _, a := range alphabet {
		if a == symbol || symbol == "ε" {
			return true
		}
	}

	return false
}

func validateCommaSeparatedWords(input string) bool {
	matched, _ := regexp.MatchString(`^(\w+)(,\w+)*$`, input)
	return matched
}

func validateCommaSeparatedAlphabet(input string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z](,[a-zA-Z])*$`, input)
	return matched
}

func validateTransitionFormat(input string) bool {
	matched, _ := regexp.MatchString(`^(\w+,[a-zA-Zε]->\w+)(;\s*\w+,[a-zA-Zε]->\w+)*$`, input)
	return matched
}

func parseNFA(statesStr, alphabetStr, transitionsStr, startStateStr, finalStatesStr string) *automata.NFA {
	// Create states
	statesNames := strings.Split(statesStr, ",")
	states := make([]*automata.State, len(statesNames))
	for i, name := range statesNames {
		states[i] = &automata.State{Name: strings.TrimSpace(name)}
	}

	// Create alphabet
	alphabet := strings.Split(alphabetStr, ",")

	// Find start state
	var startState *automata.State
	for _, state := range states {
		if state.Name == strings.TrimSpace(startStateStr) {
			startState = state
			break
		}
	}

	// Find final states
	finalStatesArr := strings.Split(finalStatesStr, ",")
	var finalStates []*automata.State
	for _, finalStateName := range finalStatesArr {
		for _, state := range states {
			if state.Name == strings.TrimSpace(finalStateName) {
				finalStates = append(finalStates, state)
			}
		}
	}

	// Parse transitions
	transitions := []automata.Transition{}
	transitionsArr := strings.Split(transitionsStr, ";")
	for _, transitionStr := range transitionsArr {
		parts := strings.Split(transitionStr, "->")
		if len(parts) != 2 {
			fmt.Printf("Invalid transition format: %s\n", transitionStr)
			continue
		}
		inputParts := strings.Split(parts[0], ",")
		if len(inputParts) != 2 {
			fmt.Printf("Invalid transition format: %s\n", transitionStr)
			continue
		}

		fromStateName := strings.TrimSpace(inputParts[0])
		input := strings.TrimSpace(inputParts[1])
		toStateName := strings.TrimSpace(parts[1])

		fromState := findState(fromStateName, states)
		toState := findState(toStateName, states)
		if fromState == nil || toState == nil {
			fmt.Printf("State not found in transition: %s\n", transitionStr)
			continue
		}

		transitions = append(transitions, automata.Transition{
			From:  fromState,
			Input: input,
			To:    []*automata.State{toState},
		})
	}

	// Return the parsed NFA
	return &automata.NFA{
		States:      states,
		Alphabet:    alphabet,
		Transitions: transitions,
		StartState:  startState,
		FinalStates: finalStates,
	}
}

func findState(name string, states []*automata.State) *automata.State {
	for _, state := range states {
		if state.Name == name {
			return state
		}
	}
	return nil
}

//func parseNFA(statesStr, alphabetStr, transitionsStr, startStateStr, finalStateStr string) *automata.NFA {
//
//	// create states arr
//	statesName := strings.Split(statesStr, ",")
//	states := make([]*automata.State, len(statesName))
//	for i, name := range statesName {
//		states[i] = &automata.State{Name: strings.TrimSpace(name)}
//	}
//
//	// alphabet
//	alphabetArr := strings.Split(alphabetStr, ",")
//	alphabet := make([]string, len(alphabetArr))
//	for i, ch := range alphabetArr {
//		alphabet[i] = strconv.Itoa(int(strings.TrimSpace(ch)[0]))
//	}
//
//	// add start state
//	var startState *automata.State
//	for _, state := range states {
//		if strings.TrimSpace(state.Name) == strings.TrimSpace(startStateStr) {
//			startState = state
//		}
//	}
//
//	// add final states
//	finalStatesArr := strings.Split(finalStateStr, ",")
//	var finalStates []*automata.State
//	for _, fState := range finalStatesArr {
//		for _, state := range states {
//			if strings.TrimSpace(fState) == strings.TrimSpace(state.Name) {
//				finalStates = append(finalStates, state)
//			}
//		}
//	}
//
//	// add transitions
//	var transitions []automata.Transition
//	transitionsArr := strings.Split(transitionsStr, ";")
//	for _, transitionStr := range transitionsArr {
//		splitInput := strings.Split(transitionStr, "->")
//		startStateName := strings.Split(splitInput[0], ",")[0]
//		startState := automata.FindState(startStateName, states)
//		char := strings.Split(splitInput[0], ",")[1]
//		endStateName := splitInput[1]
//		endState := automata.FindState(endStateName, states)
//		transitions = append(transitions, automata.Transition{
//			From:  startState,
//			Input: strconv.Itoa(int(char[0])),
//			To:    []*automata.State{endState},
//		})
//	}
//
//	return &automata.NFA{
//		States:      states,
//		Alphabet:    alphabet,
//		Transitions: transitions,
//		StartState:  startState,
//		FinalStates: finalStates,
//	}
//}

func formatStates(states []*automata.State) string {
	var names []string
	for _, state := range states {
		if state == nil {
			continue
		}
		names = append(names, state.Name)
	}
	return strings.Join(names, ", ")
}

func formatAlphabet(alphabet []string) string {
	var symbols []string
	for _, r := range alphabet {
		symbols = append(symbols, r)
	}
	return strings.Join(symbols, ", ")
}

func formatTransitions(transitions map[string]map[string]string) string {
	var result []string
	for from, inputs := range transitions {
		for input, to := range inputs {
			result = append(result, fmt.Sprintf("%s, %s -> %s", from, input, to))
		}
	}
	return strings.Join(result, "; ")
}

func formatTransitionsNfa(transitions []automata.Transition) string {
	var result []string
	for _, transition := range transitions {
		toStates := []string{}
		for _, toState := range transition.To {
			toStates = append(toStates, toState.Name)
		}
		result = append(result, fmt.Sprintf("%s, %s -> [%s]",
			transition.From.Name, string(transition.Input), strings.Join(toStates, ", ")))
	}
	return strings.Join(result, "; ")
}
