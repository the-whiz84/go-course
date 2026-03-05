package main

import "fmt"

// "math"
// "strconv"

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

	// str := "hello world"

	// Instead we use the range syntax to handle the special characters
	// for _, char := range str { // if we do not need to print or retrieve the index, we use _ instead of i
	// 	fmt.Printf("%c", char)

	// }
	// hello world

	// 6. Arrays

	//Fixed sized data structure that can store values of the same type.

	// var arr [2]int //arrays that stores integers of size 2
	// var arr [3]int //different array type. It cannot change once its initialized
	// fmt.Println(arr)
	// [0 0]

	// var arr [2]bool
	// fmt.Println(arr)
	// [false false]

	// we can also initialize arrays using the implicit assignment
	// arr := [2]int{1, 2} // this is known as an array literal where we provide the values in the array already
	// fmt.Println(arr)
	// [1 2]

	// arr := [2][2]int{{1, 2}, {2, 3}} //nested arrays
	// fmt.Println(arr)
	// [[1 2] [2 3]]

	// there is one more way to initialize an array without providing the count. The compiler will analyze the literal and intialize the correct value
	// arr := [...][2]int{{1, 2}, {2, 3}, {3, 2}}
	// fmt.Printf("%T", arr)
	// [3][2]int

	// Array properties

	// arr := [...][2]int{{1, 2}, {2, 3}, {3, 2}}
	// arr[0] = [2]int{10, 11} //access index and setting value

	// fmt.Println(arr)
	// [[10 11] [2 3] [3 2]]

	// fmt.Println(len(arr)) //access length of array

	// Loop through arrays
	// for i, value := range arr {
	// 	fmt.Println(i, value)
	// }
	// 0 [10 11]
	// 1 [2 3]
	// 2 [3 2]

	// Loop through nested array
	// for _, nested := range arr {
	// 	for _, value := range nested {
	// 		fmt.Println(value)
	// 	}
	// }
	// 10
	// 11
	// 2
	// 3
	// 3
	// 2

	// Arrays in functions
	// arr := [...][2]int{{1, 2}, {2, 3}, {3, 2}}
	// test(arr) //this doesn't actually mutate the index 0 by passing the array to the test function. In Go, by defauult, it passes a copy of the array and mutates the copy
	// fmt.Println(arr)
	// [[1 2] [2 3] [3 2]]

	// 7. Slices

	// Slices are much more flexible in Go than arrays and are used over them in most situations.
	// A slice is a view of an arrays and allows to increase or decrease the capacity of an array.

	// arr := [5]int{1, 2, 3, 4, 5}
	// sl := arr[:3] //view the array up to but not including index 3
	// fmt.Println(sl)
	// [1 2 3]

	// sl := arr[1:3]   // slice from and including index 1 up to index 3
	// fmt.Println(sl)
	// [2 3]

	// sl := arr[1:] //slice to and including last index
	// fmt.Println(sl)
	// [2 3 4 5]

	// We can access index of slices
	// sl := arr[1:]
	// fmt.Println(sl[0]) // gives us the first element in the slice
	// 2

	// We can mutate the slices
	// sl := arr[1:]
	// sl[0] = 100
	// fmt.Println(sl, arr)
	// [100 3 4 5] [1 100 3 4 5]	// when we modify the slice, we also modify that element in the array

	// Slice properties:
	// pointer -> the first element the slice points to in the underlying array
	// length  - length of slice
	// capacity  - how many elements exist from the pointer to the end of the array

	// arr := [5]int{1, 2, 3, 4, 5}
	// sl := arr[:3] 	// pointer is arr[0]
	// 					// len is 3
	// 					// capacity is 5

	// fmt.Println(sl, len(sl), cap(sl))
	// [1 2 3] 3 5

	// arr := [5]int{1, 2, 3, 4, 5}
	// sl := arr[1:3] 	// pointer is arr[1]
	// len is 2
	// capacity is 4

	// fmt.Println(sl, len(sl), cap(sl))
	// [2 3] 2 4

	// We can re-slice or expand the capacity of the slice
	// arr := [5]int{1, 2, 3, 4, 5}
	// sl := arr[1:3]
	// sl = sl[:4]

	// fmt.Println(sl, len(sl), cap(sl))
	// [2 3 4 5] 4 4

	// How to create slices without having and underlying array
	// sl := []string{"hello", "world"} // this is different than an array, an implicit array will be created with the pointer for this slice will point to it, with a len=2, cap=2
	// fmt.Println(sl)
	// // [hello world]
	// fmt.Printf("%T", sl)
	// []string   // this is the slice type, different than what we had before [3][2]int which was the array type

	// Appending to the slice
	// sl := []string{"hello", "world"}

	// for x := 0; x < 10; x++ {
	// 	sl = append(sl, "radu")
	// 	fmt.Println(sl, len(sl), cap(sl))
	// }
	// [hello world radu] 3 4
	// [hello world radu radu] 4 4
	// [hello world radu radu radu] 5 8
	// [hello world radu radu radu radu] 6 8
	// [hello world radu radu radu radu radu] 7 8
	// [hello world radu radu radu radu radu radu] 8 8
	// [hello world radu radu radu radu radu radu radu] 9 16
	// [hello world radu radu radu radu radu radu radu radu] 10 16
	// [hello world radu radu radu radu radu radu radu radu radu] 11 16
	// [hello world radu radu radu radu radu radu radu radu radu radu] 12 16

	// We can observe that the len and capacity are not the same. What happens is that the underlying array douubles its capacity to allow for the new value.
	// When it reaches the max and cannot add again, it doubles again.

	// Other ways to create slices:
	//  By specifying type, size and capacity
	// sl := make([]int, 10, 20) // type int, size 10, capacity 20

	// Iterate over slices
	// sl := []string{"hello", "world", "from", "go"}

	// for i, value := range sl {
	// 	fmt.Println(i, value)
	// }
	// 0 hello
	// 1 world
	// 2 from
	// 3 go

	// Passing a slice into a function
	// sl := []string{"hello", "world", "from", "go"}
	// test2(sl) // different behavior than with arrays. We mutate the slice when we pass it to a function.
	// fmt.Println(sl)
	// [changed this world from go]

	// 8. Maps

	// How to create and initialize a map

	// var mp map[string]int = map[string]int{"a": 1}   //or
	// mp := map[string]int{"a": 1}
	// fmt.Println(mp)
	// map[a:1]

	// We can also use the make function
	// mp := make(map[string]int) //we create an empty map
	// fmt.Println(mp)
	// map[]

	// We can create more complicated maps
	// mp := map[string][]int{"a": {1, 2, 3}} //key is a string and value is a slice
	// fmt.Println(mp)
	// fmt.Printf("%T", mp)
	// map[a:[1 2 3]]
	// map[string][]int

	// How to add values to the map
	// mp := map[string][]int{"a": {1, 2, 3}}
	// mp["b"] = []int{4, 5, 6}
	// fmt.Println(mp)
	// map[a:[1 2 3] b:[4 5 6]]

	// We can also modify values
	// mp := map[string][]int{"a": {1, 2, 3}}
	// mp["b"] = []int{4, 5, 6}
	// mp["b"] = []int{}
	// fmt.Println(mp)
	// map[a:[1 2 3] b:[]]

	// To delete a key from the map
	// mp := map[string][]int{"a": {1, 2, 3}}
	// mp["b"] = []int{4, 5, 6}
	// delete(mp, "b")
	// fmt.Println(mp)
	// map[a:[1 2 3]]

	// How to find a key inside a map
	// mp := map[string][]int{"a": {1, 2, 3}}
	// mp["b"] = []int{4, 5, 6}

	// value, ok := mp["b"]
	// fmt.Println(value, ok)
	// [4 5 6] true

	// Using maps
	mp := map[uint]uint{}
	n := uint(100)

	// For each number 1-100, check which is divisible by 1-5 and add it to the map
	for number := uint(1); number <= n; number++ {
		for d := uint(1); d <= 5; d++ {
			if number%d == 0 {
				mp[d]++
			}
		}
	}
	fmt.Println(mp)
	// map[1:100 2:50 3:33 4:25 5:20]

	// We have 100 values divisibile by 1, 50 values divisible by 2, 33 by 3 etc.

	// 9. Functions

}

// Arrays in functions
func test(arr [3][2]int) {
	arr[0] = [2]int{100, 100}
}

// Passing a slice into a function
func test2(arr []string) {
	arr[0] = "changed this"
}
