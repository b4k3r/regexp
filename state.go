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

func GenerateStates() map[string]*State {
	states := make(map[string]*State)

	// Create empty states
	states["A1"] = NewState()
	states["A2"] = NewState()
	states["A3"] = NewState()
	states["A4"] = NewState()
	states["A5"] = NewState()
	states["A6"] = NewState()
	states["A7"] = NewState()
	states["A8"] = NewState()
	states["B"] = NewState()
	states["E"] = NewState()
	states["F"] = NewState()
	states["I"] = NewState()
	states["I1"] = NewState()
	states["M"] = NewState()
	states["M1"] = NewState()
	states["J"] = NewState()
	states["N"] = NewState()
	states["K"] = NewState()

	states["A1"].transitions['#'] = states["A2"]
	states["A2"].transitions['i'] = states["A3"]
	states["A3"].transitions['n'] = states["A4"]
	states["A4"].transitions['c'] = states["A5"]
	states["A5"].transitions['l'] = states["A6"]
	states["A6"].transitions['u'] = states["A7"]
	states["A7"].transitions['d'] = states["A8"]
	states["A8"].transitions['e'] = states["B"]

	states["B"].transitions[' '] = states["B"]
	states["B"].transitions['\t'] = states["B"]
	states["B"].transitions['\r'] = states["B"]
	states["B"].transitions['"'] = states["E"]
	states["B"].transitions['<'] = states["F"]

	for i := 97; i <= 122; i++ {
		states["E"].transitions[rune(i)] = states["I"]
	}
	for i := 97; i <= 122; i++ {
		states["I"].transitions[rune(i)] = states["I"]
	}

	states["I"].transitions['.'] = states["I1"]
	states["I"].transitions['"'] = states["K"]
	states["I1"].transitions['h'] = states["J"]
	states["J"].transitions['"'] = states["K"]

	for i := 97; i <= 122; i++ {
		states["F"].transitions[rune(i)] = states["M"]
	}
	for i := 97; i <= 122; i++ {
		states["M"].transitions[rune(i)] = states["M"]
	}

	states["M"].transitions['.'] = states["M1"]
	states["M"].transitions['>'] = states["K"]
	states["M1"].transitions['h'] = states["N"]
	states["N"].transitions['>'] = states["K"]

	states["K"].isAcceptable = true

	return states
}
