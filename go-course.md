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
9. [Arrays (Next Topic)](#9-arrays-next-topic)

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

## 9. Arrays (Next Topic)

The source file introduces arrays as the next section, but array examples are not yet included.

---

Course status: this overview reflects all material currently present in `demo.go`.
