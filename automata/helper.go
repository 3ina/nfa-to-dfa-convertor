package automata

type State struct {
	Name string
}

type Transition struct {
	From  *State
	Input rune
	to    []*State
}
