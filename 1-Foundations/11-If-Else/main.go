package main

import "fmt"

func main()  {
    fmt.Println("Control Flow")

    number := 10

    if number > 15 {
        fmt.Println("Bigger than 15")
    } else {
        fmt.Println("Lesser than or equal to 15")
    }
    
    if anotherNum := number; anotherNum > 0 {
        fmt.Println("Greater than zero")
    } else {
        fmt.Println("Lesser than zero")
    }
}
