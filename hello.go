package main

import "fmt"

func vals() (int, int) {
	var oneNumber int
	var twoNumber int

	fmt.Println("Enter with first number: ")
	fmt.Scanln(&oneNumber) // This is the first input
	fmt.Println("Enter with second number: ")
	fmt.Scanln(&twoNumber) // This is the second input

	return oneNumber, twoNumber // OMG THIS FUNCTION RETURN A TWO VALUES ????????

}

func main() {

	a, b := vals()    // For to receive the two values of the function, It's necessary  two variables
	fmt.Println(a, b) // This 2 variables receive the two values of function "vals"
	var x int = 42

	var pointerX *int // Pointer
	pointerX = &x     // Assigning the value of x variable to the pointerX

	fmt.Println("Value of x variable is:", x) // Value of var x
	fmt.Println("Address memory:", pointerX)  // Address memory of pointer x

	const myConstantString = "DIEGO" // Constant

	// For loops

	i := 1 // Initialing the variable before the loop
	for i <= 3 {
		fmt.Println("This first loop repeat:", i, "x")
		i++ // This is for the looping don't be infinite
	}

	for j := 1; j <= 5; j++ { // This is the same thing as the first loop but in one line
		fmt.Println("This second loop repeat: ", j, "x")
	}

	// Conditionals

	// Operators
	// && (And)
	// || (Or)
	// == (Equals)

	var a1 = 20
	var a2 = 35
	if a1 > a2 {
		fmt.Println("Yes a1 is greater")
	} else {
		fmt.Println("No, a1 isn't greater than a2")
	}

	if num := 9; num > 9 { // With Go, you can create variable and assign the value in conditionally line
		fmt.Println("Yes num is grater than 9")
	}

	// You can change the values to test
	var c1 = true
	var c2 = false
	if c1 == true || c2 == true { // You can read this line as "If c1 be true or c2 be true"
		fmt.Println("True!!!")
	} else {
		fmt.Println("False")
	}

	if c1 == true && c2 == true { // You can read this line as "If c1 be true and c2 be true"
		fmt.Println("True!!!")
	} else {
		fmt.Println("False")
	}

}
