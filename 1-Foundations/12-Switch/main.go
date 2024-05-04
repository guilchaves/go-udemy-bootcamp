package main

import "fmt"

func weekDay(number int) string {
	switch number {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Not a valid day"
	}
}

func month(number int) string {
	var month string
	switch {
	case number == 1:
		month = "January"
	case number == 2:
		month = "February"
	case number == 3:
		month = "March"
	case number == 4:
		month = "April"
	case number == 5:
		month = "May"
	case number == 6:
		month = "June"
	case number == 7:
		month = "July"
	case number == 8:
		month = "August"
	case number == 9:
		month = "September"
	case number == 10:
		month = "October"
	case number == 11:
		month = "November"
	case number == 12:
		month = "December"
	default:
		month = "Not a valid number"
	}

	return month
}

func main() {
	fmt.Println("Switch")

	day := weekDay(3)
	fmt.Println(day)

	month := month(4)
	fmt.Println(month)
}
