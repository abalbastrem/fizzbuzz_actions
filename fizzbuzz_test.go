package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzbuzzFailNegativeNumber(t *testing.T) {
	_, err := fizzbuzz(RuleSet{
		fizz: -3,
		buzz: 5,
	})

	if !errors.Is(err, errNegativeNumber) {
		t.Errorf("expectedNegativeNumber error, got %v", err)
	}
}

func TestFizzbuzzSuccessStandardNumbers(t *testing.T) {
	res, err := fizzbuzz(RuleSet{
		fizz: 3,
		buzz: 5,
	})

	assert.Equal(t, "1", res[0])
	assert.Equal(t, "2", res[1])
	assert.Equal(t, "fizz", res[2])
	assert.Equal(t, "4", res[3])
	assert.Equal(t, "buzz", res[4])
	assert.Equal(t, "fizz", res[5])
	assert.Equal(t, "fizzbuzz", res[14])
	assert.Equal(t, "fizzbuzz", res[29])
	assert.Equal(t, "buzz", res[99])
	assert.Equal(t, "fizzbuzz", res[299])

	assert.NoError(t, err)
}

func TestFizzbuzzSuccessHighNumbers(t *testing.T) {
	res, err := fizzbuzz(RuleSet{
		fizz: 100,
		buzz: 250,
	})

	assert.Equal(t, "1", res[0])
	assert.Equal(t, "2", res[1])
	assert.Equal(t, "3", res[2])
	assert.Equal(t, "fizz", res[99])
	assert.Equal(t, "buzz", res[249])
	assert.Equal(t, "fizzbuzz", res[2499])
	assert.Equal(t, "2501", res[2500])

	assert.NoError(t, err)
}

func TestFizzbuzzSuccessLowNumbers(t *testing.T) {
	res, err := fizzbuzz(RuleSet{
		fizz: 2,
		buzz: 3,
	})

	assert.Equal(t, res[0], "1")
	assert.Equal(t, res[1], "fizz")
	assert.Equal(t, res[2], "buzz")
	assert.Equal(t, res[3], "fizz")
	assert.Equal(t, res[4], "5")
	assert.Equal(t, res[5], "fizzbuzz")
	assert.Equal(t, res[6], "7")
	assert.Equal(t, res[7], "fizz")
	assert.Equal(t, res[199], "fizz")
	assert.Equal(t, res[299], "fizzbuzz")
	assert.Equal(t, res[332], "buzz")
	assert.Equal(t, res[599], "fizzbuzz")

	assert.NoError(t, err)
}

func TestFizzbuzzSuccessSupercharged(t *testing.T) {
	res, err := fizzbuzz(RuleSet{
		fizz: 1,
		buzz: 1,
	})

	assert.Equal(t, res[0], "fizzbuzz")
	assert.Equal(t, res[1], "fizzbuzz")
	assert.Equal(t, res[2], "fizzbuzz")
	assert.Equal(t, res[3], "fizzbuzz")
	assert.Equal(t, res[4], "fizzbuzz")
	assert.Equal(t, res[5], "fizzbuzz")
	assert.Equal(t, res[6], "fizzbuzz")
	assert.Equal(t, res[7], "fizzbuzz")
	assert.Equal(t, res[199], "fizzbuzz")
	assert.Equal(t, res[299], "fizzbuzz")
	assert.Equal(t, res[599], "fizzbuzz")
	assert.Equal(t, res[999], "fizzbuzz")

	assert.NoError(t, err)
}
