package main

import (
	"testing"
)

type TestCase struct {
	name string
	input int
	expect string
}

func TestFuzzBizz(t *testing.T) {
	testCases := []TestCase{
		{"1 should return FuzzBizz", 1, "FuzzBizz"},
		{"2 should return FuzzBizz", 2, "FuzzBizz"},
	}

	for _, tt := range testCases {
		got := FuzzBizz(tt.input)

		if got != tt.expect {
			t.Errorf("%v should return %v, but got %v", tt.input, tt.expect, got)
		}
	}
}

