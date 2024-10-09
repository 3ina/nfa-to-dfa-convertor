package automata

type NFA struct {
	States      []*State
	Alphabet    []rune
	Transitions []Transition
	StartState  *State
	FinalStates []*State
}
