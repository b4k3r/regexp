package main

type Analyzer struct {
	expression    string
	states        map[string]*State
	currentState  *State
	tmpMatchedExp []rune
	matchedExp    string
}

func NewAnalyzer(expression string, states map[string]*State) *Analyzer {
	a := new(Analyzer)

	a.expression = expression
	a.states = states
	a.currentState = states["A1"]

	return a
}

func (a *Analyzer) Run() {
	for _, c := range a.expression {
		if state, exists := a.currentState.transitions[c]; exists != false {
			a.tmpMatchedExp = append(a.tmpMatchedExp, c)

			if state.isAcceptable {
				a.currentState = a.states["A1"]
				a.matchedExp = string(a.tmpMatchedExp)
			} else {
				a.currentState = state
			}
		} else {
			a.tmpMatchedExp = []rune{}
			a.currentState = a.states["A1"]
		}
	}
}
