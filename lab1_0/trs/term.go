package trs

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type TermType int

const (
	TermTypeVariable TermType = 1 << iota
	TermTypeFunction
	TermTypeConstant
)

type Term struct {
	Type   TermType
	Symbol string

	Arguments []Term
}

func NewTermFromString(str string, variables []string) (Term, error) {
	parser := TermParser{
		Variables: variables,
	}
	return parser.parseTerm(str)
}

type TermParser struct {
	Variables []string
}

func (tp TermParser) parseTerm(s string) (Term, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return Term{}, errors.New("empty term")
	}

	if tp.isContinousSetOfLetters(s) {
		if slices.Contains(tp.Variables, s) {
			return Term{Type: TermTypeVariable, Symbol: s}, nil
		}
		return Term{Type: TermTypeConstant, Symbol: s}, nil
	}

	symbol, args, err := tp.parseFunction(s)
	if err != nil {
		return Term{}, err
	}
	return Term{Type: TermTypeFunction, Symbol: symbol, Arguments: args}, nil
}

func (tp TermParser) isContinousSetOfLetters(s string) bool {
	for _, r := range s {
		if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", r) {
			return false
		}
	}
	return true
}

func (tp TermParser) parseFunction(s string) (symbol string, args []Term, err error) {
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
				term, err := tp.parseTerm(argsStr[argStart:i])
				if err != nil {
					return "", nil, err
				}
				args = append(args, term)
				argStart = i + 1
			}
		}
	}

	if argStart < len(argsStr) {
		term, err := tp.parseTerm(argsStr[argStart:])
		if err != nil {
			return "", nil, err
		}
		args = append(args, term)
	}

	return symbol, args, nil
}

func (t Term) String() string {
	switch t.Type {
	case TermTypeVariable, TermTypeConstant:
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
	if lhs.Type == TermTypeConstant {
		if rhs.Type == TermTypeConstant {
			if lhs.Symbol == rhs.Symbol {
				(*argsMap)[lhs.Symbol] = rhs
				return nil
			}
		}
		return fmt.Errorf("error in args binding: no pair for constant term found")
	}
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
		return fmt.Errorf("error in args binding: unmatched terms signature, '%s' and '%s'", lhs, rhs)
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

func (t Term) UnfoldDeep(trs TermRewritingSystem, n int) []Term {
	words := []Term{t}
	for i := 0; i < n; i++ {
		newWords := make([]Term, 0)
		for _, word := range words {
			newWords = append(newWords, word.Unfold(trs, 1)...)
		}
		words = newWords
	}
	return removeDuplicateTerms(words)
}

func removeDuplicateTerms(termSlice []Term) []Term {
	allKeys := make(map[string]bool)
	list := []Term{}
	for _, item := range termSlice {
		if _, value := allKeys[item.String()]; !value {
			allKeys[item.String()] = true
			list = append(list, item)
		}
	}
	return list
}

func (t Term) Unfold(trs TermRewritingSystem, n int) []Term {
	// fmt.Printf("n = %d, word = '%s'\n", n, t)
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
		// fmt.Println("!!!!", bindings, "word", t, "rule", rule)

		newTerm := rule.RightTerm.ApplyArgsBindings(bindings)
		// fmt.Printf("n = %d, queueing rebinded word = '%s'\n", n, newTerm)
		someRes := newTerm.Unfold(trs, n-1)
		// fmt.Printf("n = %d, rebinded word resukt = '%s', someRes = %v\n", n, newTerm, someRes)
		res = append(res, someRes...)

	}

	// fmt.Printf("n = %d, word = '%s', starting arguments unwrap\n", n, t)
	for i := 0; i < len(t.Arguments); i++ {
		// fmt.Println("here")
		newTerm := t.DeepCopy()
		// fmt.Println("newterm:", newTerm)

		// fmt.Printf("n = %d, queueing subterm word = '%s'\n", n, newTerm.Arguments[i])
		newRes := newTerm.Arguments[i].Unfold(trs, n)
		// fmt.Printf("n = %d, subterm word = '%s', newRes = %v\n", n, newTerm.Arguments[i], newRes)

		// fmt.Println(newRes)

		for _, newnewTerm := range newRes {
			newnewnewTerm := newTerm.DeepCopy()
			newnewnewTerm.Arguments[i] = newnewTerm
			// fmt.Println(newnewTerm)
			// fmt.Println(newTerm)
			// fmt.Println(t)
			res = append(res, newnewnewTerm)
			// fmt.Println(*res)
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
