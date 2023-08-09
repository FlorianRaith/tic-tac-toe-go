package main

import (
	"fmt"
)

type Board [3][3]byte

func main() {
	board := Board{
		{'_', '_', '_'},
		{'_', '_', '_'},
		{'_', '_', '_'},
	}

	player := 0

	for {
		printBoard(board)

		x, y := getInputCoordinates(player, board)
		fmt.Println()

		board[y][x] = playerChar(player)

		if checkWin(player, board) {
			printBoard(board)
			fmt.Println()
			fmt.Printf("Player%c has won", playerChar(player))
			fmt.Println()

			break
		}

		player = 1 - player
	}

}

// checkWin checks if the player has won the game
func checkWin(player int, board Board) bool {
	switch {
	case checkSequence(player, board[0][0], board[0][1], board[0][2]):
		return true
	case checkSequence(player, board[1][0], board[1][1], board[1][2]):
		return true
	case checkSequence(player, board[2][0], board[2][1], board[2][2]):
		return true
	case checkSequence(player, board[0][0], board[1][0], board[2][0]):
		return true
	case checkSequence(player, board[0][1], board[1][1], board[2][1]):
		return true
	case checkSequence(player, board[0][2], board[1][2], board[2][2]):
		return true
	case checkSequence(player, board[0][0], board[1][1], board[2][2]):
		return true
	case checkSequence(player, board[0][2], board[1][1], board[2][0]):
		return true
	}

	return false
}

// checkSequence checks if the given sequence only contains the player's character
//
// Example:
//
//	checkSequence(0, 'X', 'X', 'X') // true
//	checkSequence(0, 'X', 'O', 'X') // false
func checkSequence(player int, a byte, b byte, c byte) bool {
	char := playerChar(player)
	return char == a && char == b && char == c
}

// playerChar returns the character for the given player
// panics if the player is not 0 or 1
//
// playerChar(0) // 'X'
// playerChar(1) // 'O'
func playerChar(player int) byte {
	switch player {
	case 0:
		return 'X'
	case 1:
		return 'O'
	}

	panic(fmt.Sprintf("player must be 0 or 1, instead %d given.", player))
}

// getInputCoordinates asks the user for the coordinates and returns them
// if the input is invalid, the user is asked again
func getInputCoordinates(player int, board Board) (int, int) {
	for {
		var x, y int

		fmt.Println()
		fmt.Printf("[Player%c] Type x and y coordinates: ", playerChar(player))
		_, err := fmt.Scanf("%d %d", &x, &y)

		if err != nil {
			fmt.Println("[Error] The input could not be parsed")
			continue
		}

		if !validateCoordinates(board, x, y) {
			continue
		}

		return x, y
	}
}

// validateCoordinates checks if the given coordinates are valid
// if not, it prints an error message to the console
func validateCoordinates(board Board, x int, y int) bool {
	if x < 0 || x > 2 {
		fmt.Println("0 <= x <= 2. Try again")
		return false
	}

	if y < 0 || y > 2 {
		fmt.Println("0 <= y <= 2. Try again")
		return false
	}

	if board[y][x] != '_' {
		fmt.Printf("[Error] The field at (%d, %d) is already taken\n", x, y)
		return false
	}

	return true
}

// printBoard prints the board including coordinates to the console
//
// Example output:
//
//	  0 1 2
//	0 _ X _
//	1 _ _ O
//	2 _ _ _
func printBoard(board Board) {
	fmt.Print("  ")
	for x, _ := range board[0] {
		fmt.Printf("%d ", x)
	}

	fmt.Println()

	for y, row := range board {
		fmt.Printf("%d ", y)
		for _, char := range row {
			fmt.Printf("%c ", char)
		}

		fmt.Println()
	}
}
