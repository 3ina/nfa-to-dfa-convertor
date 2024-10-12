package automata

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
