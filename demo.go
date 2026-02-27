package main

import (
	"fmt"
	// "math"
	// "strconv"
)

func main() {
	// fmt.Println("Hello, World!")
	// uint 		//Unsigned integer, only positive values, defaults to
	// uint32
	// uint64

	// int			// Signed integer for negative numbers
	// int32
	// int64

	// float64		// default float is float64
	// float32

	// byte 		//byte type equal to int8, used to stoare a number or 1 character
	// byte = 'c'

	// rune			// rune stores int32 value

	// bool			// boolean with values true or false

	// string		// anything inside " ", only double quotes

	// nil			// equal to undefined or null in Javascript or none in Python

	// var 			// dynamic variable that can change
	// const		// static variable that cannot change
	// var x string = "Hello World!"

	// var x uint32 = 12

	// const x uint8 = 2
	// fmt.Println(x)

	// Implicit assignment, decides the type of variable at compile time, without needing to set it when assigned
	// 90% of the time we use implicit assignment, we use explicit when we need to assign a variable without giving it yet a value
	// y := 2.3
	// fmt.Println(y)

	// Typecast can be used with implicit assignment
	// y := uint32(0)

	// Typecasting doesn't work as expected in certain situations
	// x := -9
	// y := uint(x)
	// fmt.Println(x, y)
	// -9 18446744073709551607

	// fmt.Printf("%T", y) // print the type of the variable
	// x := false

	// 1. Format function fmt

	// Pringln
	// fmt.Println("hello", 2, x) // prints line and adds a new line automatically, adds spaces between each argument

	// Printf
	// fmt.Printf("%T", x)        // Print formatted is use to print specific strings
	// fmt.Printf("%T %T", x, x)				// %T for type

	// fmt.Printf("The value of x is: %v", x) 	// %v for value

	// x := 7
	// fmt.Printf("%b", x) 						// %b prints the binary representation of the value

	// x := 7.436332
	// fmt.Printf("%e", x) 						// %e print scientific notation of the value
	// 7.436332e+00

	// fmt.Printf("%f", x) 						// %f float represantation of the value
	// 7.436332

	// x := "hello"
	// fmt.Printf("%s", x) 						// %s prints the string

	// x := 2.476332
	// fmt.Printf("%.2f", x)					// %.2f print only the first 2 decimals of the float
	// 2.48

	// x := 2.476332
	// fmt.Printf("%10.2f", x)					// %10.2f  how many spaces to add
	//          2.48

	// x := 2.476332
	// fmt.Printf("%.2f%%", x) 					// %.2f%%	to add the % sign we use 2 %
	// 2.48%

	// Sprintf
	// x := 2.476332
	// y := fmt.Sprintf("%10.2f%%", x) // Sprintf creates a formatted string to use later without printing it
	// fmt.Println(y, y)
	// 2.48%       2.48%

	// 2. Arithmetic operators

	/*
		+	addition
		-	subtraction
		*	multiplication
		/	division
		--	decrement
		++	increment
		%	modulo
		We do not have the exponenent operator ** but it can be done another way
	*/
	// x := 7
	// y := 9
	// z := x + y
	// fmt.Println(z)	// we receive no error as they are the same type
	// 16

	// x := uint8(7)
	// y := 9
	// z := x + y // invalid operation: x + y (mismatched types uint8 and int)compilerMismatchedTypes

	// fmt.Println(z)

	// x := uint8(7)
	// y := 1000
	// z := x + uint8(y)

	// fmt.Println(z)
	// 239				// the larger number does not fit 8 bit so is converted

	// x := uint8(7)
	// y := 1000
	// z := int(x) + y // to solve this we need to convert the smaller value to the larger type

	// fmt.Println(z)
	// 1007

	// x := 7
	// y := 1000
	// z := y / x

	// fmt.Println(z)
	// 142				// we do not get a float but the same type, integer

	// x := 7
	// y := 1000
	// z := float64(y) / float64(x)

	// fmt.Println(z)
	// 142.85714285714286		// we need to convert both variables to float if we want the full result

	// x := "hello"
	// y := "world"
	// z := x + y

	// fmt.Println(z)
	// helloworld					// string concatination

	// x := "hello"
	// y := 2
	// z := x + string(y) // conversion from int to string yields a string of one rune, not a string of digitsstringintconvdefault
	// fmt.Println(z)
	// hello߷

	// x := "hello"
	// y := 2
	// z := x + fmt.Sprint(y)	//correct way to convert int to string
	// fmt.Println(z)
	// hello2

	// 3. Math package

	// fmt.Println(math.Min(4, 5)) // Minimum value
	// 4

	// fmt.Println(math.Max(4, 5)) // Maximum value
	// 5

	// fmt.Println(math.Sqrt(256)) // Square root of the argument
	// 16

	// fmt.Println(math.Pow(4, 2)) // raise x by the power of y - the way to do exponents
	// 16

	// fmt.Println(math.Round(4.245)) // Round a value. We also have math.Ceil() and math.Floor() to round or down.
	// 4

	// 4. String conversion with strconv package

	// x := "1234"
	// y, err := strconv.Atoi(x)	// Atoi takes a string value and converts it into a integer base 10 plus and error
	// fmt.Println(y, err)
	// 1234 <nil>

	// x := "1234"
	// y, err := strconv.ParseInt(x, 10, 0) // ParseInt we pass the string value, base of integer and bitSize (8, 16, 32, 64) or 0 for all
	// fmt.Println(y, err)
	// 1234 <nil>

	// x := "1111011"
	// y, err := strconv.ParseInt(x, 2, 0) // we can also get a binary number if we pass the base 2
	// fmt.Println(y, err)
	// 123 <nil>

	// 5. Conditions and conditionals

	// 5.1 Comparison operators
	/*
		<	less than
		>	greater than
		<=	less then or equal to
		>=	greater than or equal to
		==	equal to
		!= 	not equal to
	*/

	// x := uint(8)
	// y := 10
	// z := x < y //invalid operation: x < y (mismatched types uint and int)compilerMismatchedTypes
	// z := int(x) < y
	// fmt.Println(z)
	// true

	// x := uint(8)
	// z := x < 9 // no error because 9 can be treated as uint,  but if it were -9 it would give an error
	// fmt.Println(z)
	// true

	// 5.2 Logical operators
	/*
		||	OR
		&&	AND
		!	NOT

	*/

	// 5.3 If statements
	// x := 5
	// if x < 3 {
	// 	fmt.Println("run")

	// } else if x > 4 {
	// 	fmt.Println("do not run")

	// } else {
	// 	fmt.Println("unknown")
	// }

	// 5.4 Switch statements

	// a := 1
	// switch a {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("default")
	// }
	// one
	// In Go it automatically breaks after the matching case, unlike Javascript where you have to specify break after each case

	// Naked switch statement
	// a := 1
	// switch {
	// case a < 1:
	// 	fmt.Println("one")
	// case a > 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("default")
	// }
	// default

	// a := -2
	// switch {
	// case a < -1:
	// 	fmt.Println("a is less than -1")
	// 	fallthrough // there are reasons where you want to fallthrough the next case, for ex if one case is true you want to follow through the rest of the cases
	// case a < 0:
	// 	fmt.Println("a is less than 0")
	// 	fallthrough
	// case a < 1:
	// 	fmt.Println("a is less than 1")
	// default:
	// 	fmt.Println("default")
	// }
	// a is less than -1
	// a is less than 0
	// a is less than 1

	// Switch statements with multiple values

	// a := "h"

	// switch {
	// case a == "a", a == "b", a == "c":
	// 	fmt.Println("a is a, b or c")
	// default:
	// 	fmt.Println("default")
	// }

	// 5.5 For Loops

	// for idx := 0; idx < 10; idx++ {
	// 	fmt.Println(idx)
	// }

	// Go does not have a While loop, we need to use a condition in a for loop
	// a := 1
	// for a < 10 {
	// 	fmt.Println("loop")
	// 	a++
	// }

	// Looping through strings
	// str := "hello world"

	// fmt.Println(str[0]) // we get the integer representation of the character, not the expected indexx 0
	// 104

	// In Go strings are actually slices of bytes, so we get the value of the first byte
	// fmt.Printf("%T", str[0])
	// uint8

	// fmt.Println(string(str[0])) // We need to convert it back to a string
	// h

	// str := "hello world"

	// for idx := 0; idx < len(str); idx++ {
	// 	fmt.Printf("%c", str[idx])
	// }
	// hello world

	// Strings have different encodings: ASCII and UTF-8
	// ASCII - 1 byte to represent the different characters, which gives us 256 values
	// UTF-8 - 4 bytes which permits characters like emojis, Chinese alphabet, symbols etc. This means that we cannot use the index of the string since it might be more than 1 byte.

	// str := "😂"

	// for idx := 0; idx < len(str); idx++ {
	// 	fmt.Printf("%c", str[idx])
	// }
	// ð		//it cannot retrieve the emoji by index

	str := "hello world"

	// Instead we use the range syntax to handle the special characters
	for _, char := range str { // if we do not need to print or retrieve the index, we use _ instead of i
		fmt.Printf("%c", char)

	}
	// hello world

	// 6. Arrays

}
