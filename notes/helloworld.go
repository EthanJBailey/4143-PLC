/*****************************************************************************
 *
 *  Author:           Ethan Bailey
 *  Title:            HelloWorld Go-lang
 *  Course:           CMPS 4143
 *  Semester:         Fall 2023
 *
 *  Description:
 *        This program outputs 'Hello World!' to the console when executed
 *
 *  Usage:
 *        - Run the driver program using 'go run nameofile.go' command
 *
 *  Files:
 *        helloworld.go  :  driverprogram
*****************************************************************************/

// Instructions & Notes for the program.

// First download the go language for windows on 'go.dev'
// Every program must start with the package declaration.
// In Go language, packages are used to organize and reuse the code.
// In Go, there are two types of program available one is executable and another one is the library
// To run the program you have to use shell command 'go run nameofile.go'
// Therefore for this file 'go run helloworld.go' after I located and cd to where the file is located

// Resources

// https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-20-04
// https://zetcode.com/golang/build/

package main // added the package main in your program:
import "fmt" // import packages in your program and fmt package is used to implement formatted Input/Output with functions.

// Main function
func main() {
	fmt.Println("Hello World!")
}
