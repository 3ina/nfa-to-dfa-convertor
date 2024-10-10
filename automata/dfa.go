package automata

type DFA struct {
	States      []*State
	Alphabet    []rune
	Transitions map[string]map[rune]string
	StartState  *State
	FinalStates []*State
}
