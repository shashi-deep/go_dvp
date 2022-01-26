//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"

	"math/rand"
)

func greeting(name string) {
	fmt.Println("Hello , Have a great day ", name)
}

func addition(one, two, three int) int {
	return one + two + three
}

func randomNumber() int {

	return rand.Intn(500)
}

func randomTwoNumber() (int, int) {
	return rand.Intn(600), rand.Intn(600)
}

func main() {
	greeting("Shashi")

	var one int
	var two int
	var three int

	fmt.Scanln(&one)
	fmt.Scanln(&two)
	fmt.Scanln(&three)

	addito := addition(one, two, three)
	fmt.Println("addito ", addito)

	ones := randomNumber()
	twos, threes := randomTwoNumber()

	fmt.Println("one ", ones, "two ", twos, "three", threes)

	addThree := ones + twos + threes
	fmt.Println("randthree ", addThree)

}
