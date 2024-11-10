package main

import (
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
		if a == symbol {
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
