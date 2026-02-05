package grading

func AssignGrade(marks int) string {
	switch {
	case marks >= 80:
		return "A"
	case marks >= 70:
		return "B"
	case marks > 60:
		return "C"
	case marks >= 50:
		return "D"
	default:
		return "F"
	}
}
