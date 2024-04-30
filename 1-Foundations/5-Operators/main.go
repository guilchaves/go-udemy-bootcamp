package main

import "fmt"

func main() {
    // Math operators
    sum := 1 + 2
    difference := 1 - 2
    division := 10 / 4
    multiplication := 10 * 5
    modulus := 10 % 2
    
    fmt.Println(sum, difference, division, multiplication, modulus)

    // Attribution operators
    var variable1 string = "String"
    variable2 := "String2"
    fmt.Println(variable1, variable2)

    // Relation operators
    fmt.Println(1 > 2)
    fmt.Println(1 >= 2)
    fmt.Println(1 == 2)
    fmt.Println(1 <= 2)
    fmt.Println(1 > 2)
    fmt.Println(1 < 2)
    fmt.Println(1 != 2)
    
    // Logical operators
    truthy, falsy := true, false
    fmt.Println(truthy && falsy)
    fmt.Println(truthy || falsy)
    fmt.Println(!truthy) // false

    // Unary operators
    n := 10
    n++ // 11
    n += 15 // 26
    n-- // 25
    n -= 15 // 10
    
    fmt.Println(n) 

    // No ternary operators in Go
}
