/*****************************************************************************
*
*  Email:	     ejbailey1030@my.msutexas.edu
*  Label:            P02
*  Title:            Program 2 - Baby Steps
*  Course:           CMPS 4143
*  Semester:         Fall 2023
*
*  Description:
*        Draws a rectangle on an existing image "mustangs.jpg" and saves as a new
*		 file
*
*  Usage:
*        - Run the main.go file
* 		 - OR execute the 'go run main.go' command
*
*  Files:
*        main.go   			 :  driverprogram
*        imageManipulator.go :  contains the mascot package and function.
*        go.mod	   			 :  defines the module's module path
*        go.sum	   			 :  contains path information used to get quote from web
*
*****************************************************************************/

// Package Declaration:
package main

// Import required packages
import (
	"fmt"
	"myimageapp/imageManipulator"
)

// Main function of the program
func main() {
	// Create an ImageManipulator instance with an existing image
	im, err := imageManipulator.NewImageManipulatorWithImage("mustangs.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Draw a rectangle
	im.DrawRectangle(150, 50, 560, 411)

	// Save the modified image to a file
	im.SaveToFile("mustangs_edited.png")
}
