package automata

import "sort"

type State struct {
	Name string
}

type Transition struct {
	From  *State
	Input rune
	to    []*State
}

func contains(states []*State, target *State) bool {
	for _, state := range states {
		if state == target {
			return true
		}
	}

	return false
}

func unionSets(a, b []*State) []*State {
	unique := make(map[*State]bool)
	for _, state := range a {
		unique[state] = true
	}

	for _, state := range b {
		unique[state] = true
	}

	var result []*State

	for state := range unique {
		result = append(result, state)
	}

	return result
}

func stateSetName(states []*State) string {
	names := []string{}

	for _, state := range states {
		names = append(names, state.Name)
	}
	sort.Strings(names)
	return "{" + join(names, ",") + "}"
}

func join(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		if i != 0 {
			result += sep
		}
		result += s
	}

	return result
}

func containsAny(a []*State, b []*State) bool {
	for _, stateA := range a {
		for _, stateB := range b {
			if stateA == stateB {
				return true
			}
		}
	}
	return false
}
