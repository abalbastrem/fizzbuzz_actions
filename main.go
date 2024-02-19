package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	errNegativeNumber = errors.New("ERROR: no zero nor negative numbers allowed")
)

const (
	fizzbuzzN = 5000
)

func main() {
	fmt.Println("fizzbuzz actions")
}

type RuleSet struct {
	fizz int
	buzz int
}

func fizzbuzz(ruleset RuleSet) ([]string, error) {
	var fizzbuzzes []string

	if ruleset.fizz <= 0 || ruleset.buzz <= 0 {
		return []string{}, errNegativeNumber
	}
	for i := 1; i <= fizzbuzzN; i++ {
		var result string
		if i%ruleset.fizz != 0 && i%ruleset.buzz != 0 {
			result += strconv.Itoa(i)
		}
		if i%ruleset.fizz == 0 {
			result += "fizz"
		}
		if i%ruleset.buzz == 0 {
			result += "buzz"
		}
		fizzbuzzes = append(fizzbuzzes, result)
	}

	return fizzbuzzes, nil
}
