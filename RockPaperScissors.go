/*
Created by Miles Colbert 04/06/2022
Rock Paper Scissors in Golang
Assignment given for Nephron Interview
Create rock paper scissors game that will compare the user's string of "rock", "paper", or "scissors" to the computers string
*/
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	newGame := true
	fmt.Println("\nWelcome to the rock paper scissors game")
	for newGame == true {

		// Gets the user's input
		fmt.Println("Please enter \"rock\", \"paper\", or \"scissors\"")
		var user string
		fmt.Scanln(&user)
		user = strings.ToLower(user)

		// User has invalid move
		if (user != "rock") && (user != "paper") && (user != "scissors") {
			fmt.Println("Invalid move, exiting the game")
			break
		}

		// Gets the cpu's move
		var x int
		cpu := cpuMove(x)

		// Compares user to the cpu
		output := moveComparison(user, cpu)
		// Prints the results
		fmt.Printf(output)

		// Prompts user to continue
		fmt.Println("Would you like to continue? Enter \"Yes\" or \"No\"")
		var retry string
		fmt.Scanln(&retry)

		var choice string
		choice, newGame = Retry(retry)
		fmt.Println(choice)

	}

}

// Randomly makes the user's move and returns the string value
func cpuMove(x int) string {
	enemy := ""
	rand.Seed(time.Now().UnixNano()) // Necessary for random numbers
	max := 3
	min := 0
	x = rand.Intn(max-min) + min
	// x = rand.Intn(3) will always give the same result, seed is needed

	if x == 0 { // Rock
		enemy = "rock"
	}
	if x == 1 { // Paper
		enemy = "paper"
	}
	if x == 2 { // Scissors
		enemy = "scissors"
	}
	return enemy
}

// Compares the user's move to the cpu's move and returns a string value with the result
func moveComparison(userInput, cpuInput string) string {
	var result string

	// Tie
	if userInput == cpuInput {
		result = "You tied, both of you have chose " + userInput + "\n"
	}
	// You win
	if ((userInput == "rock") && (cpuInput == "scissors")) || ((userInput == "paper") && (cpuInput == "rock")) || ((userInput == "scissors") && (cpuInput == "paper")) {
		result = "You won, you chose " + userInput + " and the computer chose " + cpuInput + "\n"
	}
	// You lose
	if ((userInput == "rock") && (cpuInput == "paper")) || ((userInput == "paper") && (cpuInput == "scissors")) || ((userInput == "scissors") && (cpuInput == "rock")) {
		result = "You lost, you chose " + userInput + " and the computer chose " + cpuInput + "\n"
	}
	return result
}

func Retry(retry string) (string, bool) {
	var answer bool // ends or continues loop
	var text string // returns result of what you chose

	// User enters invalid input
	if strings.ToLower(retry) != "yes" && strings.ToLower(retry) != "no" {
		text = ("You have entered an invalid input, exiting the game")
		answer = false
	}

	// User quits
	if strings.ToLower(retry) == "no" {
		text = ("You have chosen to quit, thank you for playing")
		answer = false
	}

	// User continues
	if strings.ToLower(retry) == "yes" {
		text = ("You have chosen to play again, good luck!\n\n")
		answer = true
	}
	return text, answer
}
