package automata

import "sort"

type State struct {
	Name string
}

type Transition struct {
	From  *State
	Input string
	To    []*State
}

func FindState(name string, states []*State) *State {
	for _, state := range states {
		if state.Name == name {
			return state
		}
	}

	return nil
}

func contains(states []*State, target *State) bool {
	for _, state := range states {
		if state.Name == target.Name {
			return true
		}
	}
	return false
}

func unionSets(a, b []*State) []*State {
	unique := make(map[string]*State)
	for _, state := range a {
		unique[state.Name] = state
	}
	for _, state := range b {
		unique[state.Name] = state
	}

	var result []*State
	for _, state := range unique {
		result = append(result, state)
	}

	return result
}

func stateSetName(states []*State) string {
	var names []string
	for _, state := range states {
		if state != nil {
			names = append(names, state.Name)
		}
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
	stateSet := make(map[string]bool)
	for _, state := range a {
		stateSet[state.Name] = true
	}
	for _, state := range b {
		if stateSet[state.Name] {
			return true
		}
	}
	return false
}
