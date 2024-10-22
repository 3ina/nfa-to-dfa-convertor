package automata

type NFA struct {
	States      []*State
	Alphabet    []rune
	Transitions []Transition
	StartState  *State
	FinalStates []*State
}

func (nfa *NFA) epsilonClosure(states []*State) []*State {
	closure := make([]*State, 0)
	visited := make(map[*State]bool)

	var stack []*State
	stack = append(stack, states...)

	for len(stack) > 0 {
		state := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[state] {
			visited[state] = true
			closure = append(closure, state)
			for _, t := range nfa.Transitions {
				if t.From == state && t.Input == 'ε' {
					for _, toState := range t.to {
						if !visited[toState] {
							stack = append(stack, toState)
						}
					}

				}
			}
		}
	}

	return closure
}

func (nfa *NFA) move(states []*State, input rune) []*State {
	result := make([]*State, 0)
	for _, state := range states {
		for _, t := range nfa.Transitions {
			if t.From == state && t.Input == input {
				result = append(result, t.to...)
			}
		}
	}
	return result
}
