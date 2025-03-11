package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Board struct {
	state [9]string
}
type Game struct {
	board         *Board
	currentPlayer int
}

func (game *Game) play() {
	for i := 0; i < 9; i++ {
		fmt.Printf("player %d, enter your preferred position: ", game.currentPlayer)

		input := game.getPlayerInput()

		game.board.state[input-1] = getPlayerString(game.currentPlayer)

		// display the current game state and the position just inserted

		game.displayGameState()

		fmt.Printf("Player %d played position %d\n", game.currentPlayer, input)

		// check winner
		isWin := game.checkWinner()
		if isWin {
			fmt.Printf("Player %d won \n", game.currentPlayer)
			return
		} else {
			time.Sleep(time.Second * 2)
			clearStdout()
			if game.currentPlayer == 1 {
				game.currentPlayer = 2
			} else {
				game.currentPlayer = 1
			}
		}
	}

	fmt.Println("game ended in a draw")
}

func (game *Game) getPlayerInput() int {
	var input int
	stdin := bufio.NewReader(os.Stdin)

	for {
		_, _ = fmt.Fscan(stdin, &input)
		if input > 0 && input < 10 {
			if game.board.state[input-1] != "" {
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

func (game *Game) displayGameState() {
	for i := 0; i < 9; i += 3 {
		fmt.Printf(" %s  |  %s  |  %s\n",
			getDisplayValue(game.board.state[i]),
			getDisplayValue(game.board.state[i+1]),
			getDisplayValue(game.board.state[i+2]),
		)
		if i < 6 {
			fmt.Println("--- | --- | ---")
		}
	}

}

func (game *Game) checkWinner() bool {
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
	lastPlayerString := getPlayerString(game.currentPlayer)
	for _, winArr := range winPositions {
		a, b, c := winArr[0], winArr[1], winArr[2]
		if game.board.state[a] == lastPlayerString && game.board.state[b] == lastPlayerString && game.board.state[c] == lastPlayerString {
			return true
		}
	}
	return false
}

func main() {
	// build a tic tac toe game
	fmt.Printf("Welcome to tic tac toe in the terminal\n")
	fmt.Printf("Below is the game layout\n\n")
	displayGamePosition()
	fmt.Printf("\nwe will expect each player to enter the position they want to play")
	time.Sleep(time.Second * 4)
	clearStdout()

	game := Game{
		board:         new(Board),
		currentPlayer: 1,
	}

	game.play()
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

func clearStdout() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getPlayerString(player int) string {
	if player == 1 {
		return "X"
	}
	return "O"
}
