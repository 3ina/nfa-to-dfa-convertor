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

func (nfa *NFA) ConvertToDfa() *DFA {
	dfa := &DFA{
		Alphabet:    nfa.Alphabet,
		Transitions: make(map[string]map[rune]string),
	}

	stateMap := make(map[string]string)
	startClosure := nfa.epsilonClosure([]*State{nfa.StartState})
	startName := stateSetName(startClosure)
	dfa.StartState = &State{Name: startName}
	dfa.States = append(dfa.States, dfa.StartState)
	stateMap[startName] = startName

	if containsAny(startClosure, nfa.FinalStates) {
		dfa.FinalStates = append(dfa.FinalStates, dfa.StartState)
	}

	queue := [][]*State{startClosure}

	for len(queue) > 0 {
		currentSet := queue[0]
		queue = queue[1:]
		currentName := stateSetName(currentSet)
		for _, input := range nfa.Alphabet {
			moved := nfa.move(currentSet, input)
			closure := nfa.epsilonClosure(moved)
			if len(closure) == 0 {
				continue
			}
			newName := stateSetName(closure)
			if _, exists := stateMap[newName]; !exists {
				stateMap[newName] = newName
				newState := &State{Name: newName}
				dfa.States = append(dfa.States, newState)
				if containsAny(closure, nfa.FinalStates) {
					dfa.FinalStates = append(dfa.FinalStates, newState)
				}
				queue = append(queue, closure)
			}

			if dfa.Transitions[currentName] == nil {
				dfa.Transitions[currentName] = make(map[rune]string)
			}
			dfa.Transitions[currentName][input] = newName

		}
	}

	return dfa
}
