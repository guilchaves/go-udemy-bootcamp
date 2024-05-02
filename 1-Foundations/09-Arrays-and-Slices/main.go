package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and slices")

    // Array
	var array1 [5]string
	array1[0] = "Position 1"
	fmt.Println(array1)

	array2 := [5]string{"Pos 1", "Pos 2", "Pos 3", "Pos 4", "Pos 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

    // Slice
    slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
    fmt.Println(slice)

    fmt.Println(reflect.TypeOf(slice))
    fmt.Println(reflect.TypeOf(array3))

    slice = append(slice, 18)
    fmt.Println(slice)

    // Internals 
    slice3 := make([]float32, 10, 15)
    fmt.Println(slice3)

    slice4 := make([]float32, 5)
    fmt.Println(slice4)
    fmt.Println(len(slice4))
    fmt.Println(cap(slice4))



}
