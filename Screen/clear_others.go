package screen

import "fmt"

// Clear clears the screen
func Clear1() {
	fmt.Print("\033[2J")
}

// MoveTopLeft moves the cursor to the top left position of the screen
func MoveTopLeft1() {
	fmt.Print("\033[H")
}
