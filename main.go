package main

import (
	"math/rand"
	"time"
)


func existsInSlice(a int, list [4]int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}


func createCode() [4]int{
	code := [4]int{0, 0, 0, 0}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		code[i] = rand.Intn(8)
	}

	return code
}

func checkGuess(code, guess [4]int) [2]int{
	correctPlace := 0
	correctDigit := 0
	for i := 0; i < 4; i++ {
		if code[i] == guess[i] {
			correctPlace += 1
		} else if existsInSlice(guess[i], code) {
			correctDigit += 1
		}
	}

	return [2]int{correctPlace, correctDigit}
}

func main() {
	//correct := false
	//turns := 0
	//code := createCode()
	//fmt.Println("4-digit Code has been set. Digits in range 1 to 8. You have 12 turns to break it.")
	//
	//for turns < 12 {
	//
	//}
}
