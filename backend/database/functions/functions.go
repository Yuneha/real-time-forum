package functions

import (
	"fmt"
	"time"
)

func FormattedDate(date string) string {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}
	formattedDate := parsedDate.Format("02/01/2006")
	return formattedDate
}

func GetAge(date string) int {
	birthTime, err := time.Parse("02/01/2006", date)
	if err != nil {
		fmt.Println("Error parsing birthdate:", err)
		return 0
	}
	currentTime := time.Now()
	age := currentTime.Year() - birthTime.Year()
	if currentTime.YearDay() < birthTime.YearDay() {
		age--
	}
	return age
}
