//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {

	var myFavColour = "Green"
	fmt.Println("My fav clr is ", myFavColour)

	year, ageYears := 1996,25
	fmt.Println("my birth year is ",year," and ageYears is ",ageYears)

	var (
		intialStart = "D"
		initialEnd = "P"
	)

	fmt.Println("first & last initials", intialStart,initialEnd)

	var age int
	fmt.Println("init my age ",age)

	age = ageYears * 354

	fmt.Println("my age in days ",age)
	
}
