package greetings

import (
	"fmt"
	"time"
)

func SayHello() {
	fmt.Println("Hello, Go!")
}

func GreetByYear(year int) string {
	currentYear := time.Now().Year() // Получаем текущий год

	switch {
	case year >= 1946 && year < 1965:
		return "Привет, бумер!"
	case year >= 1965 && year < 1981:
		return "Привет, представитель X!"
	case year >= 1981 && year < 1996:
		return "Привет, миллениал!"
	case year >= 1996 && year < 2012:
		return "Привет, зумер!"
	case year >= 2012 && (year <= currentYear || year < 2030):
		return "Привет, альфа!"
	default:
		return "Привет, путешественник во времени!"
	}
}
