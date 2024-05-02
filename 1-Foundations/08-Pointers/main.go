package main

import "fmt"

func main()  {
    fmt.Println("Pointers")

    var variable1 int = 10
    var variable2 int = variable1

    fmt.Println(variable1, variable2)

    variable1++
    fmt.Println(variable1, variable2)

    var variable3 int = 100
    var pointer *int = &variable3 

    fmt.Println(variable3, *pointer)
}
