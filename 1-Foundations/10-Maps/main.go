package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":    "Guilherme",
		"surname": "Chaves",
	}

	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"fist":   "John",
			"second": "Doe",
		},
		"major": {
			"name":   "Law",
			"campus": "Campus 1",
		},
	}

	fmt.Println(user2)
	delete(user2, "name")
	fmt.Println(user2)

    user2["minor"] = map[string]string{

    }

}
