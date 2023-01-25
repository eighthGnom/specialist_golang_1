package main

import (
	"fmt"
	"strings"
)

func printDate(date string) string {
	newDate := strings.ReplaceAll(date, "-", "-0")

	switch len(newDate) {
	case 7:
		return fmt.Sprintf("000" + newDate)
	case 8:
		return fmt.Sprintf("00" + newDate)
	case 9:
		return fmt.Sprintf("0" + newDate)
	default:
		return fmt.Sprintf(newDate)
	}
}
