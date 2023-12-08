package main

import "fmt"

func vals() (int, int) {
	var oneNumber int
	var twoNumber int

	fmt.Println("Enter with first number: ")
	fmt.Scanln(&oneNumber) // This is the first input
	fmt.Println("Enter with second number: ")
	fmt.Scanln(&twoNumber) // This is the first input

	return oneNumber, twoNumber // OMG THIS FUNCTIONS RETURN A TWO VALUES ????????

}

func main() {

	a, b := vals()    // For to receive the two values of the function, It's necessary  two variables
	fmt.Println(a, b) // This 2 variables receive the two values of function "vals"
	var x int = 42

	var pointerX *int // Pointer
	pointerX = &x     // Assigning the value of x variable to the pointerX

	fmt.Println(x)        // Value of var x
	fmt.Println(pointerX) // Address memory of pointer x

	const myConstantString = "DIEGO" // Constant

	// For loops

	i := 1 // Initialing the variable before for loop
	for i <= 3 {
		fmt.Println("This loop repeat ", i, "x")
		i++ // This is for the looping don't be infinite
	}
}
