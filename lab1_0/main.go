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

	trss, err := trs.NewTermRewritingSystemFromString(cfg.RawRules, cfg.Variables)
	if err != nil {
		panic(err)
	}
	fmt.Printf("parsed trs: %s\n", trss)

	word, err := trs.NewTermFromString(cfg.Word, cfg.Variables)
	if err != nil {
		panic(err)
	}
	fmt.Printf("parsed word: %+v\n", word)

	fmt.Println(word.UnfoldDeep(trss, cfg.N))

	// word.Unfold(trss, cfg.N)
	// fmt.Println(word.Unfold(trss, cfg.N))

	// lhs, _ := trs.NewTermFromString("g(g(t,t),h(t))", []string{"t", "x"})
	// rhs, _ := trs.NewTermFromString("g(x,x)", []string{"t", "x"})

	// fmt.Println(rhs.BindArguments(lhs))
}
