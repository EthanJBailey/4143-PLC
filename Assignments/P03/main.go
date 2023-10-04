package main

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

	// Prints black space
	fmt.Print("\n")

	// Call function to process Pixel Colors
	Colors.PrintPixels()

	// Prints black space
	fmt.Print("\n")

	// Call function to grayscale image
	Grayscale.GrayScale()

	// Prints black space
	fmt.Print("\n")

	// Call function to print colored text to image
	Text.PrintText()
}
