package main

import "testing"

func TestSomma(t *testing.T) {
	if 2+2 != 4 {
		t.Fatal("2+2 non fa 4")
	}
}