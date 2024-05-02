package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	// float32, float64

	var number int16 = 100
	fmt.Println(number)

	//alias
	//int32 = rune
	var number3 rune = 123456
	fmt.Println(number3)

	//byte = uint8
	var number4 byte = 123
	fmt.Println(number4)

    // real numbers

    var realNumber float32 = 12342.1
    fmt.Println(realNumber)

    var realNumber2 float64 = 1234123090000.23
    fmt.Println(realNumber2)

    realNumber3 := 12345.42
    fmt.Println(realNumber3)

    // end real numbers

    //Strings
    var str string = "Text"
    fmt.Println(str)

    str2 := "Text 2"
    fmt.Println(str2)

    char := 'B' // print 66
    fmt.Println(char)
    //end strings

    var boolean1 bool = true
    fmt.Println(boolean1)

    var err error = errors.New("Internal error") // <nil>
    fmt.Println(err)
}

