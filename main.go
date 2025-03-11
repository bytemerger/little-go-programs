package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// build a tic tac toe game
	gameState := make([]string, 9)
	playerTurn := 1
	fmt.Printf("Welcome to tic tac toe in the terminal\n")
	fmt.Printf("Below is the game layout\n\n")
	displayGamePosition()
	fmt.Printf("\nwe will expect each player to enter the position they want to play")
	time.Sleep(time.Second * 4)
	clearStdout()

	for i := 0; i < 9; i++ {
		fmt.Printf("player %d, enter your preferred position: ", playerTurn)

		input := getPlayerInput(gameState)

		if playerTurn == 1 {
			gameState[input-1] = "X"
		} else {
			gameState[input-1] = "O"
		}

		// display the current game state and the position just inserted

		displayGameState(gameState)

		fmt.Printf("Player %d played position %d\n", playerTurn, input)

		// check winner
		isWin := checkWinner(gameState, playerTurn)
		if isWin {
			fmt.Printf("Player %d won \n", playerTurn)
			return
		} else {
			time.Sleep(time.Second * 2)
			clearStdout()
			if playerTurn == 1 {
				playerTurn = 2
			} else {
				playerTurn = 1
			}
		}

	}
	fmt.Println("game ended in a draw")

}

func displayGamePosition() {
	for i := 0; i < 9; i += 3 {
		fmt.Printf(" %d  |  %d  |  %d\n", i+1, i+2, i+3)
		if i < 6 {
			fmt.Println("--- | --- | ---")
		}
	}

}

func getDisplayValue(value string) string {
	if value == "" {
		return "-"
	}
	return value
}

func displayGameState(state []string) {
	for i := 0; i < 9; i += 3 {
		fmt.Printf(" %s  |  %s  |  %s\n",
			getDisplayValue(state[i]),
			getDisplayValue(state[i+1]),
			getDisplayValue(state[i+2]),
		)
		if i < 6 {
			fmt.Println("--- | --- | ---")
		}
	}

}

func clearStdout() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getPlayerInput(gameState []string) int {
	var input int
	stdin := bufio.NewReader(os.Stdin)

	for {
		_, _ = fmt.Fscan(stdin, &input)
		if input > 0 && input < 10 {
			if gameState[input-1] != "" {
				clearStdout()
				fmt.Print("Sorry invalid position, please enter a valid input: ")
				continue
			}
			break
		}
		// clear the remaing things in the buffer to prevent infinite looop
		stdin.ReadString('\n')

		clearStdout()
		fmt.Print("Sorry invalid position, please enter a valid input: ")
	}
	return input
}

func checkWinner(gameState []string, lastPlayer int) bool {
	winPositions := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
	var lastPlayerString string
	if lastPlayer == 1 {
		lastPlayerString = "X"
	} else {
		lastPlayerString = "O"
	}
	for _, winArr := range winPositions {
		if gameState[winArr[0]] == lastPlayerString && gameState[winArr[1]] == lastPlayerString && gameState[winArr[2]] == lastPlayerString {
			return true
		}
	}
	return false
}
