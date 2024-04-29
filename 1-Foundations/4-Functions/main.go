package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathCalc(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	difference := n1 - n2
	return sum, difference
}

func main() {
	total := sum(10, 20)
	fmt.Println(total)

	var f = func(txt string) string {
		fmt.Println(txt)
        return txt
	}
	//type function
    result := f("Function f")
	fmt.Println(result)

    resultSum, resultDifference := mathCalc(10, 15)
    fmt.Println(resultSum, resultDifference)
}
