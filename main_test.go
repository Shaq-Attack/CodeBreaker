package main

import (
	"testing"
)


func TestCreateCode(t *testing.T) {

	if len(createCode()) < 4 {
		t.Error("Failed. The generated code must be four(4) numbers long.")
	} else {
		t.Log("Success. The generated code is correct!")
	}
}


func TestCheckGuessCorrectPlace(t *testing.T) {
	if checkGuess([4]int{1,2,5,6}, [4]int{1,2,3,4}) != [2]int{2, 0} {
		t.Error("Failed.")
	} else {
		t.Log("Success.")
	}
}

func TestCheckGuessCorrectDigit(t *testing.T) {
	if checkGuess([4]int{1,2,5,6}, [4]int{5,6,7,8}) != [2]int{0, 2} {
		t.Error("Failed.")
	} else {
		t.Log("Success.")
	}
}


func TestExistsInSlice(t *testing.T) {
	if existsInSlice(1, [4]int{1,2,3,4}) {
		t.Log("Success.")
	} else {
		t.Error("Failed.")
	}
}