/*****************************************************************************
*
*  Email:	     ejbailey1030@my.msutexas.edu
*  Label:            P01
*  Title:            Program 1 - Run a Go Program
*  Course:           CMPS 4143
*  Semester:         Fall 2023
*
*  Description:
*        Ouputs the name of the Best Mascot and a quote pulled from the internet.
*
*  Usage:
*        - Run the main.go file
* 		 - OR execute the 'go run main.go' command
*
*  Files:
*        main.go   :  driverprogram
*        mascot.go :  contains the mascot package and function.
*        go.mod	   :  defines the module's module path
*        go.sum	   :  contains path information used to get quote from web
*
*****************************************************************************/

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
