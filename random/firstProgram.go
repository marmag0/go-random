package main

// standard library import
import "fmt"

// main function, code execution starts from there
func main() {
	// STDOUT standard formatting print
	fmt.Println("hellow world")

	// variables

	// int	 	int8 	int16 	int32 	int64
	// uint		uint8	uint16 	uint32 	uint64 uintptr (always positive)

	// byte (alias for uint8)

	// rune (alias for uint32 - unicode code point)

	// float32 float64

	// complex64 complex128

	// assigning value to variables
	var cost float64 = 0.02
	var parts int = 11

	// short variable declaration (autoassigned by go)
	totalCost := cost * float64(parts) // float64

	fmt.Println(totalCost)
	fmt.Printf("The type of cost is: %T\n", cost)

	// constants (they only can be specified during compile time)
	const rockType string = "hard"
	fmt.Printf("The rock is " + rockType + "\n")

	// string formatting
	// prints to STDOUT
	fmt.Printf("The rock is %v\n", rockType)
	// returns formatted string
	var some_text string = fmt.Sprintf("The rock is %v", rockType)

	// inserting variable
	fmt.Printf("Last print was: %v\n", some_text)
	// interpolating string
	fmt.Printf("Some %s here\n", "text")
	// interpolating an integer in decimal form
	fmt.Printf("My fave number is %d\n", 8)
	// interpolating an float
	fmt.Printf("My fave float number is %.2f\n", 8.9)
}
