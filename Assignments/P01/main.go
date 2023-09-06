// Package Declaration
package main

// Import required packages
import (
	"fmt"

	"example.com/p01/mascot"
	"rsc.io/quote"
)

// Main Function of program
func main() {
	// Prints the output of BestMascot function to the console
	fmt.Println(mascot.BestMascot())
	// Prints quote from package pulled from the web
	fmt.Println(quote.Go())
}
