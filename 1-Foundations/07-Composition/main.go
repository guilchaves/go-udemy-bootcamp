package main

import "fmt"

type person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type student struct {
	person
	course     string
	university string
}

func main() {
	fmt.Println("Composition")

    p1 := person{"John", "Doe", 20, 178}
    fmt.Println(p1)

    s1 := student{p1, "Engineering", "USP"}
    fmt.Println(s1)
}
