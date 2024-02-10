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

	word.Unfold(trss, cfg.N)
	fmt.Println(word.Unfold(trss, cfg.N))

}