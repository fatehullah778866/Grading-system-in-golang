package grading

import "fmt"

func GetInput() (string, int) {
	var name string
	var marks int
	fmt.Print("Enter student name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter marks: ")
	fmt.Scanln(&marks)
	return name, marks
}
