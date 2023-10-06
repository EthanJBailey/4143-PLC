/*****************************************************************************
*
*  Email:	     ejbailey1030@my.msutexas.edu
*  Label:            P03
*  Title:            Program 3 - Image Ascii Art
*  Course:           CMPS 4143
*  Semester:         Fall 2023
*
*  Description:
*        This program hosts four different packages that allow a user to 
*        "go get" them on their local machine from anywhere. 
*        They all perform varous form of Image Manipulation or ASCII art.
*
*  Usage:
*        - Run the 'go get' command to pull the libraries created
*
*  Files:
*        main.go   			 :  driverprogram
*        go.mod	   			 :  defines the module's module path        
*
*****************************************************************************/

// Package Declaration
package main

// Import required packages
import (
	"fmt"

	"github.com/EthanJBailey/img_mod/Colors"
	"github.com/EthanJBailey/img_mod/Getpic"
	"github.com/EthanJBailey/img_mod/Grayscale"
	"github.com/EthanJBailey/img_mod/Text"
)

func main() {
	// Prints Heading of the Program
	fmt.Println("Ethan Bailey ")
	fmt.Println("Program 3 - Image Ascii Art")

	// Prints black space
	fmt.Print("\n")

	// Call function to get picture from URL
	Getpic.DownloadPicture()

	// Prints blank space
	fmt.Print("\n")

	// Call function to process Pixel Colors
	Colors.PrintPixels()

	// Prints blank space
	fmt.Print("\n")

	// Call function to grayscale image
	Grayscale.GrayScale()

	// Prints blank space
	fmt.Print("\n")

	// Call function to print colored text to image
	Text.PrintText()
}
