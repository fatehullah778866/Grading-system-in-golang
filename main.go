package main

import (
	"fmt"
	"grading-system/grading"
)

const TotalMarks = 100

func main() {
	fmt.Println("===Simple Grading System===")
	for i := 1; i <= 3; i++ {
		name, marks := grading.GetInput()
		percentage := grading.CalculatePercentage(marks, TotalMarks)
		status := grading.AssignGrade(marks)
		grade := grading.PassOrFail(marks)
		grading.DisplayResult(name, marks, percentage, status, grade)
	}
}
