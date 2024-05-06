package main

import "fmt"

// Named return function
func calculations(n1, n2 int) (sum int, diff int) {
	sum = n1 + n2
	diff = n1 - n2
	return
}

// Variadic functions
func adding(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

// Recursive function
func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

// Defer function
func gradeStudent(n1, n2 float32) bool {
	defer fmt.Println("Finished calculating grade. Result will be shown")
	fmt.Println("Checking if student is approved...")

	grade := (n1 + n2) / 2

	if grade >= 6 {
		return true
	}

	return false
}

// Panic and Recover
func gradeStudent2(n1, n2 float32) bool {
    defer recoverOperation()
    grade := (n1 + n2) / 2

    if grade > 6 {
        return true
    } else if grade < 6 {
        return false
    }

    panic("GRADE IS EXACTLY 6!")
}

func recoverOperation() {
    if r := recover(); r != nil {
        fmt.Println("Operation recovered succesfully!")
    }
}

// Closure functions
func closure() func() {
    text := "Inside closure func"

    function := func() {
        fmt.Println(text)
    }

    return function
}

// Pointer functions
func invertSign(number int) int {
    return number * -1
}

func invertSignWithPointer(number *int) {
    *number = *number * - 1
}

// init function
func init(){ 
    fmt.Println("Init is executed before main func")
}


func main() {
	fmt.Println("Named return")
	sum, diff := calculations(10, 20)
	fmt.Println(sum, diff)
	fmt.Println("------")

	fmt.Println("Variadic functions")
	sumTotal := adding(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sumTotal)

	write("Hello World", 10, 2, 3, 4, 5, 6)

	fmt.Println("------")
	fmt.Println("Anonymous functions")

	// Anonymous functions
	anom := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("Param")

	fmt.Println(anom)

	fmt.Println("------")
	fmt.Println("Recursive functions")

	position := uint(10)

	for i := uint(1); i <= position; i++ {
		fmt.Println(fibonacci(i))
	}

	fmt.Println("------")
	fmt.Println("Defer")

	fmt.Println(gradeStudent(7, 6))

	fmt.Println("------")
	fmt.Println("Panic and Recover")

    fmt.Println(gradeStudent2(6, 6))

	fmt.Println("------")
	fmt.Println("Closure function")

    text := "Inside main func"
    fmt.Println(text)

    newFunction := closure()
    newFunction()

	fmt.Println("------")
	fmt.Println("Pointer functions")

    number := 10
    invertedNumber := invertSign(number)
    fmt.Println(number, invertedNumber)

    newNumber := 40
    fmt.Println(newNumber)
    invertSignWithPointer(&newNumber)
    fmt.Println(newNumber)

}
