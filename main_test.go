package main

import (
	"bytes"
	"testing"
)


func TestCreateCode(t *testing.T) {
	if len(createCode()) < 4 {
		t.Error("Failed. The generated code must be four(4) numbers long.")
	} else {
		t.Log("Passed.")
	}
}


func TestGetInputShort(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("12\n"))

	text, err := getGuess(&stdin)
	if err != nil || text == 0{
		t.Log("Passed.")
	} else {
		t.Error("Failed.")
	}
}


func TestGetInputLong(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("12345\n"))

	text, err := getGuess(&stdin)
	if err != nil || text == 0{
		t.Log("Passed.")
	} else {
		t.Error("Failed.")
	}
}


func TestGetInputABC(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("abcd\n"))

	text, err := getGuess(&stdin)
	if err != nil || text == 0{
		t.Log("Passed.")
	} else {
		t.Error("Failed.")
	}
}


func TestGetInputPass(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("1234\n"))

	text, err := getGuess(&stdin)
	if err != nil || text == 0{
		t.Log("Passed.")
	} else {
		t.Error("Failed.")
	}
}


func TestCheckGuessCorrectPlace(t *testing.T) {
	if checkGuess([4]int{1,2,5,6}, [4]int{1,2,3,4}) != [2]int{2, 0} {
		t.Error("Failed.")
	} else {
		t.Log("Passed.")
	}
}

func TestCheckGuessCorrectDigit(t *testing.T) {
	if checkGuess([4]int{1,2,5,6}, [4]int{5,6,7,8}) != [2]int{0, 2} {
		t.Error("Failed.")
	} else {
		t.Log("Passed.")
	}
}


func TestExistsInSlice(t *testing.T) {
	if existsInSlice(1, [4]int{1,2,3,4}) {
		t.Log("Passed.")
	} else {
		t.Error("Failed.")
	}
}


func TestShowResults(t *testing.T) {
	if showResults(3,1) != "Number of correct digits in correct place:     3\n" +
		"Number of correct digits not in correct place: 1" {
		t.Error("Failed.")
	} else {
		t.Log("Passed.")
	}
}


func TestShowCode(t *testing.T) {
	if showCode([4]int{1,2,3,4}) != "The code was: 1234" {
		t.Error("Failed.")
	} else {
		t.Log("Passed.")
	}
}


func TestCorrectnessFalse(t *testing.T) {
	output := checkCorrectness(1, 2, false)
	if output {
		t.Error("Failed.")
	} else {
		t.Log("Passed.")
	}
}

func TestCorrectnessTrue(t *testing.T) {
	output := checkCorrectness(1, 4, false)
	if output {
		t.Log("Passed.")
	} else {
		t.Error("Failed.")
	}
}