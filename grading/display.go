package grading

import "fmt"

func DisplayResult(name string, marks int, percentage float64, status string, grade string) {
	fmt.Println("\n---Result---")
	fmt.Println("Name: ", name)
	fmt.Println("Marks: ", marks)
	fmt.Printf("Percentage: %.2f%%\n", percentage)

	fmt.Println("Status: ", status)
	fmt.Println("Grade: ", grade)
}
