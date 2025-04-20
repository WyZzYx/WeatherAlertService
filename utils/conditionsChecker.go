package utils

import (
	"fmt"
	"strings"
)

func CheckCondition(temp float64, condition string) bool {
	var op string
	var threshold float64

	if strings.Contains(condition, "<=") {
		op = "<="
	} else if strings.Contains(condition, ">=") {
		op = ">="
	} else if strings.Contains(condition, "<") {
		op = "<"
	} else if strings.Contains(condition, ">") {
		op = ">"
	} else if strings.Contains(condition, "==") {
		op = "=="
	} else {
		fmt.Println("Unsupported condition:", condition)
		return false
	}

	_, err := fmt.Sscanf(condition, "temperature "+op+" %f", &threshold)
	if err != nil {
		fmt.Println("Error parsing condition:", err)
		return false
	}

	switch op {
	case "<":
		return temp < threshold
	case ">":
		return temp > threshold
	case "<=":
		return temp <= threshold
	case ">=":
		return temp >= threshold
	case "==":
		return temp == threshold
	}

	return false
}
