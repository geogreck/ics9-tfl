package trs

import (
	"errors"
	"fmt"
	"strings"
)

type TermType int

const (
	TermTypeVariable TermType = 1 << iota
	TermTypeFunction
)

type Term struct {
	Type   TermType
	Symbol string

	Arguments []Term
}

func NewTermFromString(str string) (Term, error) {
	return parseTerm(str)
}

func parseTerm(s string) (Term, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return Term{}, errors.New("empty term")
	}

	if isVariable(s) {
		return Term{Type: TermTypeVariable, Symbol: s}, nil
	}

	symbol, args, err := parseFunction(s)
	if err != nil {
		return Term{}, err
	}
	return Term{Type: TermTypeFunction, Symbol: symbol, Arguments: args}, nil
}

func isVariable(s string) bool {
	for _, r := range s {
		if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", r) {
			return false
		}
	}
	return true
}

func parseFunction(s string) (symbol string, args []Term, err error) {
	openingParenIndex := strings.Index(s, "(")
	if openingParenIndex == -1 {
		return "", nil, errors.New("no opening parenthesis found for function")
	}

	symbol = s[:openingParenIndex]
	argsStr := s[openingParenIndex+1 : len(s)-1]

	argStart := 0
	parenCount := 0
	for i, char := range argsStr {
		switch char {
		case '(':
			parenCount++
		case ')':
			parenCount--
		case ',':
			if parenCount == 0 {
				term, err := parseTerm(argsStr[argStart:i])
				if err != nil {
					return "", nil, err
				}
				args = append(args, term)
				argStart = i + 1
			}
		}
	}

	if argStart < len(argsStr) {
		term, err := parseTerm(argsStr[argStart:])
		if err != nil {
			return "", nil, err
		}
		args = append(args, term)
	}

	return symbol, args, nil
}

func (t Term) String() string {
	switch t.Type {
	case TermTypeVariable:
		return t.Symbol
	case TermTypeFunction:
		args := make([]string, 0, len(t.Arguments))
		for _, arg := range t.Arguments {
			args = append(args, arg.String())
		}
		return t.Symbol + "(" + strings.Join(args, ",") + ")"
	default:
		return ""
	}
}

func (lhs *Term) BindArguments(rhs Term) (map[string]Term, error) {
	res := make(map[string]Term)
	err := lhs.bindArguments(rhs, &res)

	return res, err
}

func (lhs *Term) bindArguments(rhs Term, argsMap *map[string]Term) error {
	if lhs.Type == TermTypeVariable {
		if otherVar, exists := (*argsMap)[lhs.Symbol]; exists {
			if !otherVar.IsEquival(rhs) {
				return fmt.Errorf("error in args binding: got diffent binding for same arg")
			}
			return nil
		}
		(*argsMap)[lhs.Symbol] = rhs
		return nil
	}

	if lhs.Symbol != rhs.Symbol || len(lhs.Arguments) != len(rhs.Arguments) {
		return fmt.Errorf("error in args binding: unmatched terms signature")
	}

	for i := 0; i != len(lhs.Arguments); i++ {
		err := lhs.Arguments[i].bindArguments(rhs.Arguments[i], argsMap)
		if err != nil {
			return fmt.Errorf("error in args binding: recurse binding failed: %w", err)
		}
	}

	return nil
}

func (t Term) ApplyArgsBindings(bindings map[string]Term) Term {
	if t.Type == TermTypeVariable {
		if interpretation, exists := bindings[t.Symbol]; exists {
			return interpretation
		}
		return t
	}

	newArgs := make([]Term, len(t.Arguments))
	for i, arg := range t.Arguments {
		newArgs[i] = arg.ApplyArgsBindings(bindings)
	}
	return Term{Type: TermTypeFunction, Symbol: t.Symbol, Arguments: newArgs}
}

func (lhs Term) IsEquival(rhs Term) bool {
	res, _ := lhs.isEquival(rhs, make(map[string]string))
	return res
}

func (lhs Term) isEquival(rhs Term, argsMap map[string]string) (bool, map[string]string) {
	if lhs.Type != rhs.Type {
		return false, argsMap
	}

	if lhs.Type == TermTypeVariable {
		if otherVar, exists := argsMap[lhs.Symbol]; exists {
			return otherVar == rhs.Symbol, argsMap
		} else {
			argsMap[lhs.Symbol] = rhs.Symbol
			return true, argsMap
		}
	}

	if lhs.Symbol != rhs.Symbol || len(lhs.Arguments) != len(rhs.Arguments) {
		return false, argsMap
	}

	for i := range lhs.Arguments {
		var areArgsEquivalent bool
		areArgsEquivalent, argsMap = lhs.Arguments[i].isEquival(rhs.Arguments[i], argsMap)
		if !areArgsEquivalent {
			return false, argsMap
		}
	}

	return true, argsMap
}

func (t Term) Unfold(trs TermRewritingSystem, n int) []Term {

	res := make([]Term, 0)
	if n == 0 {
		res = append(res, t)
		return res
	}

	if t.Type == TermTypeVariable {
		return res
	}

	for _, rule := range trs.Rules {
		bindings, err := rule.LeftTerm.BindArguments(t)

		if err != nil {
			continue
		}

		newTerm := rule.RightTerm.ApplyArgsBindings(bindings)

		res = append(res, newTerm.Unfold(trs, n-1)...)
	}

	for i := 0; i < len(t.Arguments); i++ {

		newTerm := t.DeepCopy()

		newRes := newTerm.Arguments[i].Unfold(trs, n)

		for _, newnewTerm := range newRes {
			newnewnewTerm := newTerm.DeepCopy()
			newnewnewTerm.Arguments[i] = newnewTerm

			res = append(res, newnewnewTerm)
		}
	}
	return res
}

func (term Term) DeepCopy() Term {
	copy := Term{
		Type:      term.Type,
		Symbol:    term.Symbol,
		Arguments: make([]Term, len(term.Arguments)),
	}

	for i, arg := range term.Arguments {
		copy.Arguments[i] = arg.DeepCopy()
	}

	return copy
}
