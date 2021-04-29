package main

import (
	"testing"
)

func TestCalulcate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Excepted 2 + 2 to equal 4")
	}
}