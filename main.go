package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strconv"
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


func getGuess(stdin io.Reader) (int, error) {
	fmt.Printf("Input 4 digit code: ")
	reader := bufio.NewReader(stdin)
	text, err := reader.ReadString('\n')
	if err != nil || len(text) != 4{
		return 0, err
	} else {
		text, err := strconv.Atoi(text[0:4])
		if err != nil {
			return 0, err
		}
		return text, err
	}
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


func showResults(correctPlace, correctDigit int) string {
	return "Number of correct digits in correct place:     " + strconv.Itoa(correctPlace) +
		"\nNumber of correct digits not in correct place: " + strconv.Itoa(correctDigit)
}


func showCode(code [4]int) string{
	var answer string
	for x := 0; x < len(code); x++ {
		answer += strconv.Itoa(code[x])
	}

	return "The code was: " + answer
}


func checkCorrectness(turns, correctPlace int, correct bool) bool {
	if correctPlace == 4 {
		correct = true
		fmt.Println("Congratulations! You are a codebreaker!")
	} else {
		turns -= 12
		if turns == 0 {
			fmt.Println("You've run out of guesses!")
		} else {
			fmt.Printf("Turns left: %v", turns)
		}
	}
	return correct
}


func main() {
	//correct := false
	//turns := 0
	//code := createCode()
	//fmt.Println("4-digit Code has been set. Digits in range 1 to 8. You have 12 turns to break it.")
	//
	//for turns < 12 {
	//	var text int
	//	for true {
	//		text, err := getGuess(os.Stdin)
	//		if err != nil {
	//			fmt.Println("Try again")
	//		} else {
	//			break
	//		}
	//	}
	//
	//	temp := make([]int, text)
	//	var ans [4]int
	//	for x :=0; x<4;x++ {
	//		ans[x] = temp[x]
	//	}
	//	correctPnD := checkGuess(code, ans)
	//}
}
