package main

import "fmt"

func main() {
    var variable1 string = "Variable 1"
    variable2 := "Variable 2"

    fmt.Println(variable1)
    fmt.Println(variable2)

    var (
        variable3 string = "variable 3"
        variable4 string = "variable 4"
    )

    fmt.Println(variable3, variable4)

    variable5, variable6 := "Variable 5", "Variable 6"
    fmt.Println(variable5, variable6)
}
