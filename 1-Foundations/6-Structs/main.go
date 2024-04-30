package main

import "fmt"

type user struct {
	name string
	age  uint8
    address address
}

type address struct {
    street string
    number uint8
}

func main() {
	fmt.Println("File structs")

	var u user
	u.name = "Íris"
	u.age = 32
	fmt.Println(u)

    mockAddress := address{"Milwaukee Ave", 0}

	user2 := user{"Guilherme", 32, mockAddress}

	fmt.Println(user2)

    user3 := user{age: 32}
	fmt.Println(user3)

}
