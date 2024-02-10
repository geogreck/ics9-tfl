package trs

import (
	"fmt"

	avkstream "github.com/geogreck/ics9-tfl/lab1_0/avk_stream"
)

type TermRule struct {
	LeftTerm  Term
	RightTerm Term
}

func NewTermRuleFromString(str string) (TermRule, error) {

	streamer := avkstream.NewStream(str, []rune{' '})

	buf := ""
	for streamer.Peek() != '-' {
		buf += string(streamer.Next())
	}
	leftTerm, err := NewTermFromString(buf)
	if err != nil {
		return TermRule{}, fmt.Errorf("term rule parse error: %w", err)
	}

	for streamer.Peek() != '>' {
		streamer.Next()
	}
	streamer.Next()

	buf = ""
	for streamer.Peek() != 0 {
		buf += string(streamer.Next())
	}
	rightTerm, err := NewTermFromString(buf)
	if err != nil {
		return TermRule{}, fmt.Errorf("term rule parse error: %w", err)
	}

	res := TermRule{
		LeftTerm:  leftTerm,
		RightTerm: rightTerm,
	}

	return res, nil
}

func (tr TermRule) String() string {
	return fmt.Sprintf("%s -> %s", tr.LeftTerm, tr.RightTerm)
}
