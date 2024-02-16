package main

import (
	"errors"
	"fmt"
)

var (
	errNegativeNumber = errors.New("ERROR: no negative numbers allowed")
)

func main() {
	fmt.Println("hello world")
}

type RuleSet struct {
	fizz int
	buzz int
}

func fizzbuzz(ruleset RuleSet) ([]string, error) {
	var fizzbuzz []string
	// TODO
	return fizzbuzz, errNegativeNumber
}
