package main

import (
	"fmt"
	"os"

	"github.com/geogreck/ics9-tfl/lab1_0/config"
	"github.com/geogreck/ics9-tfl/lab1_0/trs"
	"gopkg.in/yaml.v3"
)

const inputFileName = "input.yaml"

func main() {
	file, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	cfg := config.InputConfig{}
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}

	trss, err := trs.NewTermRewritingSystemFromString(cfg.RawRules)
	if err != nil {
		panic(err)
	}
	fmt.Printf("parsed trs: %s\n", trss)

	word, err := trs.NewTermFromString(cfg.Word)
	fmt.Printf("parsed word: %s\n", word)

	// lhs, _ := trs.NewTermFromString("f(g(x,y,z),z,x)")
	// rhs, _ := trs.NewTermFromString("f(g(a,b,c),a,a)")

	// fmt.Println(lhs.IsCompatible(rhs))

	// lhs, _ := trs.NewTermFromString("h(g(t,t))")
	// rhs, _ := trs.NewTermFromString("h(t)")

	// fmt.Println(rhs.BindArguments(lhs))
	// fmt.Println(lhs.ApplyArgsBindings(lhs.BindArguments(rhs)))

	// fmt.Println()
	word.Unfold(trss, cfg.N)
	fmt.Println(word.Unfold(trss, cfg.N))
	// fmt.Println(trs.UnfoldTerm(word, trss, cfg.N))

}
