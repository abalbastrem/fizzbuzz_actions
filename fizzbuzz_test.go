package main

import "testing"

func TestFizzbuzzFail(t *testing.T) {
	_, err := fizzbuzz(RuleSet{
		fizz: 3,
		buzz: 5,
	})

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
