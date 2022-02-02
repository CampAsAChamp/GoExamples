package main // Declare a main package (a way to group functions, and is made up of all the files in the same directory)

import (
	"fmt"

	"rsc.io/quote/v4"
) // Import the fmt package

// Main function executes by default when you run the main package
func main() {
	fmt.Println("Hello World!")
	fmt.Println(quote.Go())
}
