// Name of the package...
// All Golang files belong to a package
// here main.go belongs to main package
// It is not necessary to name this file as main.go, this could be anyting(may be helloworld.go)
package main

// Importing standarf Go Package fmt
import "fmt"

// main function this is entry point of main function
func main() {
	// Printing welcome message using fmt.Println
	fmt.Println("Welcome to Golang! This is going to be a wonderful journey!!!")

	// fmt package has three methods to print on standard output (console)
	//fmt.Println(), fmt.Print(), fmt.Printf

	fmt.Print("\nHey!! I am print this methis using fmt.Print\n")

	fmt.Printf("\nThis print is using fmt.Printf, this is the C stype print function, we can use format specifiers!!!\n")
}
