package trs

import (
	avkstream "github.com/geogreck/ics9-tfl/lab1_0/avk_stream"
)

type TermRewritingSystem struct {
	Rules []TermRule
}

func NewTermRewritingSystemFromString(str string) (TermRewritingSystem, error) {
	var res TermRewritingSystem
	streamer := avkstream.NewStream(str, []rune{})

	for streamer.Peek() != 0 {
		buf := ""
		for streamer.Peek() != '\n' && streamer.Peek() != 0 {
			buf += string(streamer.Next())
		}
		rule, err := NewTermRuleFromString(buf)
		if err != nil {
			return TermRewritingSystem{}, err
		}
		res.Rules = append(res.Rules, rule)
		streamer.Next()
	}

	return res, nil
}
