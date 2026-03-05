# Go Course Overview

This document turns `demo.go` into a clean, beginner-friendly Go course summary focused on Go-specific behavior.

## Table of Contents

1. [Go Basics and Data Types](#1-go-basics-and-data-types)
2. [Variables, Constants, and Type Inference](#2-variables-constants-and-type-inference)
3. [Output and Formatting with `fmt`](#3-output-and-formatting-with-fmt)
4. [Arithmetic and Type Conversion](#4-arithmetic-and-type-conversion)
5. [Math Utilities with `math`](#5-math-utilities-with-math)
6. [String/Number Conversion with `strconv`](#6-stringnumber-conversion-with-strconv)
7. [Conditions and Control Flow](#7-conditions-and-control-flow)
8. [Loops and Strings](#8-loops-and-strings)
9. [Arrays](#9-arrays)
10. [Slices](#10-slices)
11. [Maps](#11-maps)
12. [Structs](#12-structs)

## 1. Go Basics and Data Types

Go is statically typed, so every variable has a type at compile time.

Common types from this lesson:

- `uint`, `uint32`, `uint64`: unsigned integers (no negatives)
- `int`, `int32`, `int64`: signed integers
- `float64` (default float), `float32`
- `byte`: alias of `uint8`
- `rune`: alias of `int32` (Unicode code point)
- `bool`: `true` or `false`
- `string`: text in double quotes
- `nil`: zero value for pointers, interfaces, maps, slices, channels, and functions

## 2. Variables, Constants, and Type Inference

Go supports explicit declarations and inferred declarations.

```go
var x string = "Hello World"
var n uint32 = 12
const limit uint8 = 2

y := 2.3 // inferred as float64
```

Type conversion is always explicit:

```go
x := -9
y := uint(x)
fmt.Println(x, y)
// -9 18446744073709551607
```

Beginner note: converting negative signed values to unsigned types wraps to a very large number.

## 3. Output and Formatting with `fmt`

### `Println`

- Prints values separated by spaces
- Adds a newline automatically

```go
x := false
fmt.Println("hello", 2, x)
```

### `Printf`

Format verbs used in the source:

- `%T` type
- `%v` value
- `%b` binary
- `%e` scientific notation
- `%f` float
- `%s` string

```go
x := 2.476332
fmt.Printf("%.2f\n", x)   // 2.48
fmt.Printf("%10.2f\n", x) // padded width
fmt.Printf("%.2f%%\n", x) // 2.48%
```

### `Sprintf`

Use `Sprintf` to build formatted strings without printing.

```go
x := 2.476332
y := fmt.Sprintf("%10.2f%%", x)
fmt.Println(y)
```

## 4. Arithmetic and Type Conversion

Operators covered:

- `+`, `-`, `*`, `/`, `%`, `++`, `--`

Go requires matching numeric types in expressions.

```go
x := uint8(7)
y := 1000
// z := x + y // compile error (mismatched types)
```

If you cast a large value into a small type, it can overflow/wrap:

```go
x := uint8(7)
y := 1000
z := x + uint8(y)
fmt.Println(z) // 239
```

Safer pattern: cast the smaller value upward.

```go
z := int(x) + y
fmt.Println(z) // 1007
```

Integer division returns an integer:

```go
x := 7
y := 1000
fmt.Println(y / x) // 142
```

Use float conversion for a fractional result:

```go
fmt.Println(float64(y) / float64(x))
// 142.85714285714286
```

String concatenation works with `+`, but be careful with `string(int)`:

```go
fmt.Println("hello" + "world") // helloworld

x := "hello"
y := 2
z := x + fmt.Sprint(y)
fmt.Println(z) // hello2
```

## 5. Math Utilities with `math`

The course examples include:

```go
fmt.Println(math.Min(4, 5))    // 4
fmt.Println(math.Max(4, 5))    // 5
fmt.Println(math.Sqrt(256))    // 16
fmt.Println(math.Pow(4, 2))    // 16
fmt.Println(math.Round(4.245)) // 4
```

Also mentioned: `math.Ceil()` and `math.Floor()`.

## 6. String/Number Conversion with `strconv`

### `Atoi`

Convert base-10 string to `int`.

```go
x := "1234"
y, err := strconv.Atoi(x)
fmt.Println(y, err) // 1234 <nil>
```

### `ParseInt`

Convert string with custom base and bit size.

```go
x := "1234"
y, err := strconv.ParseInt(x, 10, 0)
fmt.Println(y, err) // 1234 <nil>
```

Binary example:

```go
x := "1111011"
y, err := strconv.ParseInt(x, 2, 0)
fmt.Println(y, err) // 123 <nil>
```

## 7. Conditions and Control Flow

You already know basic conditionals, so here are the Go-specific points.

Comparison operators:

- `<`, `>`, `<=`, `>=`, `==`, `!=`

Logical operators:

- `||`, `&&`, `!`

Go checks type compatibility in comparisons:

```go
x := uint(8)
y := 10
z := int(x) < y
fmt.Println(z) // true
```

`switch` breaks automatically after a match unless `fallthrough` is used.

```go
a := 1
switch a {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
default:
    fmt.Println("default")
}
```

Naked switch (condition-only):

```go
a := -2
switch {
case a < -1:
    fmt.Println("a is less than -1")
    fallthrough
case a < 0:
    fmt.Println("a is less than 0")
    fallthrough
case a < 1:
    fmt.Println("a is less than 1")
default:
    fmt.Println("default")
}
```

## 8. Loops and Strings

Go uses `for` for all loop forms.

Counter loop:

```go
for idx := 0; idx < 10; idx++ {
    fmt.Println(idx)
}
```

Condition loop:

```go
a := 1
for a < 10 {
    fmt.Println("loop")
    a++
}
```

### Strings, bytes, and UTF-8

- `str[i]` returns a byte (`uint8`)
- some characters use multiple bytes in UTF-8
- `range` iterates by Unicode code points (`rune`)

```go
str := "hello world"
fmt.Println(str[0])         // 104
fmt.Printf("%T\n", str[0]) // uint8
fmt.Println(string(str[0])) // h
```

Unicode-safe iteration:

```go
str := "hello world"
for _, char := range str {
    fmt.Printf("%c", char)
}
```

If a string might contain emoji or non-ASCII text, prefer `range` over byte indexing.

## 9. Arrays

Arrays in Go are **fixed-size collections** of values of the same type.

Key points from the examples:

- The length is part of the type: `[2]int` and `[3]int` are different types.
- Arrays have a fixed size that cannot change after declaration.
- Elements get the zero value of their type by default.

```go
var arr [2]int
fmt.Println(arr) // [0 0]

var flags [2]bool
fmt.Println(flags) // [false false]
```

### Array literals and nested arrays

You can declare and initialize arrays in one statement using literals, including nested arrays.

```go
arr := [2]int{1, 2}
fmt.Println(arr) // [1 2]

matrix := [2][2]int{{1, 2}, {2, 3}}
fmt.Println(matrix) // [[1 2] [2 3]]
```

Go can infer the length from the literal using `...`:

```go
matrix := [...][2]int{{1, 2}, {2, 3}, {3, 2}}
fmt.Printf("%T\n", matrix) // [3][2]int
```

### Indexing, assignment, and length

Use index syntax and `len` with arrays just like with slices and strings.

```go
matrix := [...][2]int{{1, 2}, {2, 3}, {3, 2}}
matrix[0] = [2]int{10, 11}

fmt.Println(matrix)
// [[10 11] [2 3] [3 2]]

fmt.Println(len(matrix)) // 3
```

You can iterate with `range`, including nested loops:

```go
matrix := [...][2]int{{10, 11}, {2, 3}, {3, 2}}

for i, row := range matrix {
    fmt.Println(i, row)
}

for _, row := range matrix {
    for _, v := range row {
        fmt.Println(v)
    }
}
```

### Arrays and functions

Arrays have **value semantics**. Passing an array to a function copies the entire array.

```go
func test(arr [3][2]int) {
    arr[0] = [2]int{100, 100}
}

func main() {
    matrix := [...][2]int{{1, 2}, {2, 3}, {3, 2}}
    test(matrix)
    fmt.Println(matrix)
    // [[1 2] [2 3] [3 2]]
}
```

Beginner note: if you need a function to modify a collection, slices or pointers to arrays are usually a better choice than passing arrays by value.

## 10. Slices

Slices are Go’s primary sequence type. A slice is a **view over an underlying array** and can grow or shrink.

Important properties:

- A slice has a pointer to the first element it can access in the backing array.
- `len` is the number of elements in the slice.
- `cap` is how many elements exist from the pointer to the end of the backing array.

```go
arr := [5]int{1, 2, 3, 4, 5}
sl := arr[:3]

fmt.Println(sl, len(sl), cap(sl))
// [1 2 3] 3 5
```

### Creating and slicing

You can create slices by slicing an array or an existing slice:

```go
arr := [5]int{1, 2, 3, 4, 5}

sl1 := arr[:3]   // [1 2 3]
sl2 := arr[1:3]  // [2 3]
sl3 := arr[1:]   // [2 3 4 5]
```

Slices support indexing and can mutate the underlying array:

```go
sl := arr[1:]
sl[0] = 100

fmt.Println(sl)  // [100 3 4 5]
fmt.Println(arr) // [1 100 3 4 5]
```

Beginner note: multiple slices can share the same backing array; changing one slice can change what another slice sees.

### Length and capacity

Slicing changes the length and sometimes the capacity:

```go
arr := [5]int{1, 2, 3, 4, 5}

sl := arr[1:3]
fmt.Println(sl, len(sl), cap(sl))
// [2 3] 2 4

sl = sl[:4]
fmt.Println(sl, len(sl), cap(sl))
// [2 3 4 5] 4 4
```

### Slice literals and `append`

You can create a slice directly without naming an array:

```go
sl := []string{"hello", "world"}
fmt.Println(sl)        // [hello world]
fmt.Printf("%T\n", sl) // []string
```

Use `append` to grow a slice. When needed, Go allocates a new, larger backing array and copies the data.

```go
sl := []string{"hello", "world"}

for x := 0; x < 3; x++ {
    sl = append(sl, "go")
    fmt.Println(sl, len(sl), cap(sl))
}
```

Beginner note: because `append` may reallocate, two slices that used to share a backing array might stop sharing after enough appends.

You can also create slices with `make`:

```go
nums := make([]int, 10, 20) // len 10, cap 20
fmt.Println(len(nums), cap(nums))
```

### Ranging over slices and passing to functions

Slices work naturally with `range` and are commonly passed into functions that mutate them.

```go
sl := []string{"hello", "world", "from", "go"}

for i, v := range sl {
    fmt.Println(i, v)
}

func test2(arr []string) {
    arr[0] = "changed this"
}
```

Passing a slice to `test2` lets the function modify the underlying data.

## 11. Maps

Maps in Go are built-in hash tables: `map[keyType]valueType`.

### Creating and initializing maps

You can use a literal or `make`:

```go
mp := map[string]int{"a": 1}
fmt.Println(mp) // map[a:1]

empty := make(map[string]int)
fmt.Println(empty) // map[]
```

Maps can have more complex value types:

```go
mp := map[string][]int{"a": {1, 2, 3}}
fmt.Println(mp)        // map[a:[1 2 3]]
fmt.Printf("%T\n", mp) // map[string][]int
```

### Reading, writing, and deleting

Use index syntax to read and write entries:

```go
mp := map[string][]int{"a": {1, 2, 3}}
mp["b"] = []int{4, 5, 6}
fmt.Println(mp) // map[a:[1 2 3] b:[4 5 6]]

mp["b"] = []int{}
fmt.Println(mp) // map[a:[1 2 3] b:[]]

delete(mp, "b")
fmt.Println(mp) // map[a:[1 2 3]]
```

### Checking for key presence

Use the “comma-ok” idiom to distinguish between a missing key and a zero value:

```go
mp := map[string][]int{"a": {1, 2, 3}}
mp["b"] = []int{4, 5, 6}

value, ok := mp["b"]
fmt.Println(value, ok) // [4 5 6] true
```

### Example: counting divisors

The demo includes a small counting example:

```go
mp := map[uint]uint{}
n := uint(100)

for number := uint(1); number <= n; number++ {
    for d := uint(1); d <= 5; d++ {
        if number%d == 0 {
            mp[d]++
        }
    }
}
fmt.Println(mp)
// map[1:100 2:50 3:33 4:25 5:20]
```

Beginner note: Go does not guarantee map iteration order. Never rely on the order of keys when ranging over a map.

## 12. Structs

Structs let you define your own composite types by grouping fields together.

Basic struct definition and usage:

```go
type User struct {
    ID   int
    Name string
}

func main() {
    var u User
    fmt.Println(u) // {0 ""}

    u2 := User{ID: 1, Name: "Alice"}
    fmt.Println(u2) // {1 Alice}
}
```

Fields can be updated using dot notation:

```go
u := User{ID: 1, Name: "Alice"}
u.Name = "Bob"
fmt.Println(u) // {1 Bob}
```

Structs follow value semantics: assigning a struct copies all its fields.

```go
u1 := User{ID: 1, Name: "Alice"}
u2 := u1
u2.Name = "Bob"

fmt.Println(u1) // {1 Alice}
fmt.Println(u2) // {1 Bob}
```

If you need to share and mutate the same instance, use pointers:

```go
u := &User{ID: 1, Name: "Alice"}
u.Name = "Bob"
fmt.Println(u) // &{1 Bob}
```

### Structs with slices and maps

Structs combine naturally with slices and maps for modeling real data:

```go
type Article struct {
    ID      int
    Title   string
    Tags    []string
    Metrics map[string]int
}

articles := []Article{
    {
        ID:    1,
        Title: "Intro to Go",
        Tags:  []string{"go", "basics"},
        Metrics: map[string]int{
            "views": 100,
        },
    },
}
```

Beginner note: combining structs with slices and maps is a common way to represent domain models in Go programs.
