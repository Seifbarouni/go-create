package colorizetext

import "fmt"

var (
	// Colors
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

// Prints text with a specified color
func PrintWithColor(text string, color string) {
	fmt.Print(color + text + Reset)
}
