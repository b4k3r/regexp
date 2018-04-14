package main

type State struct {
	char         rune
	transitions  map[rune]*State
	isAcceptable bool
}

func NewState() *State {
	s := new(State)
	s.transitions = make(map[rune]*State)
	s.isAcceptable = false

	return s
}

func GenerateStates() map[rune]*State {
	states := make(map[rune]*State)

	// Create empty states
	states['A'] = NewState()
	states['B'] = NewState()
	states['D'] = NewState()
	states['E'] = NewState()
	states['F'] = NewState()
	states['G'] = NewState()
	states['H'] = NewState()

	// Create state A
	for i := 97; i <= 122; i++ {
		states['A'].transitions[rune(i)] = states['B']
	}

	// Create state B
	for i := 97; i <= 122; i++ {
		states['B'].transitions[rune(i)] = states['B']
	}
	states['B'].transitions['.'] = states['D']

	// Create state D
	states['D'].transitions['p'] = states['E']
	states['D'].transitions['c'] = states['F']

	// Create state E
	states['E'].transitions['l'] = states['G']

	// Create state E
	states['F'].transitions['o'] = states['H']

	// Create state H
	states['H'].transitions['m'] = states['G']

	states['G'].isAcceptable = true

	return states
}
