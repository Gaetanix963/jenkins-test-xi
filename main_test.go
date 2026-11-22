package main

// Commento su una singola riga per testare il push automatico su Jenkins

import "testing"

func TestSomma(t *testing.T) {
	if 2+2 != 4 {
		t.Fatal("2+2 non fa 4")
	}
}