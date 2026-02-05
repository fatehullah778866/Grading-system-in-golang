package grading

func CalculatePercentage(marks int, total int) float64 {
	return (float64(marks) / float64(total)) * 100
}

func PassOrFail(marks int) string {
	if marks >= 50 {
		return "Pass"
	}
	return "Fail"
}
