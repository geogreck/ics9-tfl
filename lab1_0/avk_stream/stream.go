package avkstream

import (
	"slices"
)

// https://github.com/geogreck/ics9-dm/blob/master/module1/polish.go#L11
type Stream struct {
	expr          []rune
	i             int
	symbolsToSkip []rune
}

func NewStream(str string, symbolsToSkip []rune) Stream {
	return Stream{
		expr:          []rune(str),
		i:             0,
		symbolsToSkip: symbolsToSkip,
	}
}

func (s Stream) Peek() rune {
	for s.i < len(s.expr) && (slices.Contains(s.symbolsToSkip, s.expr[s.i])) {
		s.i++
	}
	var symb rune
	if s.i >= len(s.expr) {
		symb = 0
	} else {
		symb = s.expr[s.i]
	}
	return symb
}

func (s *Stream) Next() rune {
	for s.i < len(s.expr) && (slices.Contains(s.symbolsToSkip, s.expr[s.i])) {
		s.i++
	}
	var symb rune
	if s.i >= len(s.expr) {
		symb = 0
	} else {
		symb = s.expr[s.i]
	}
	s.i++
	return symb
}
