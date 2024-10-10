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
