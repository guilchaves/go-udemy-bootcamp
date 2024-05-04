package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementing i", i)
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementing j", j)
		time.Sleep(time.Second)
	}

	names := [3]string{"Guilherme", "Íris", "Shiro"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	// print without the index
	for _, name := range names {
		fmt.Println(name)
	}

	// never forget to convert char to string
	for index, letter := range "WORD" {
		fmt.Println(index, string(letter))
	}

	user := map[string]string{
		"name":    "Íris",
		"surname": "Schardt",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

}
