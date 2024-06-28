package main

import (
	"testing"
)

func TestFuzzBizz(t *testing.T) {
	got := FuzzBizz(1)

	want := "FuzzBizz"

	if got != want {
		t.Errorf("1 should return FuzzBizz")
	} 
}
