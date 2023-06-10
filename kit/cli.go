package kit

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	exitOption        = "Para salir presione cualquier tecla"
	optionMessage     = "Ingrese una opcion:"
	errorInputMessage = "Ocurrió un error, intente nuevamente"
)

func PrintOptions(title string, options ...string) {
	fmt.Println(title)
	time.Sleep(1 * time.Second)
	for i := 0; i < len(options); i++ {
		fmt.Println(options[i])
	}
	fmt.Println(optionMessage)
	fmt.Println(exitOption)
}

func ScanMonthAndYear() (time.Time, error) {
	yearStr, err := ScanOptionSelectedWithMessage("Indique el año")
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	monthStr, err := ScanOptionSelectedWithMessage("Indique el mes")
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC), nil
}

func ScanDate() (time.Time, error) {
	yearStr, err := ScanOptionSelectedWithMessage("Indique el año")
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	monthStr, err := ScanOptionSelectedWithMessage("Indique el mes")
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	dayStr, err := ScanOptionSelectedWithMessage("Indique el dia")
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

func ScanDateAndHour() (time.Time, error) {
	yearStr, err := ScanOptionSelectedWithMessage("Indique el año")
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	monthStr, err := ScanOptionSelectedWithMessage("Indique el mes")
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	dayStr, err := ScanOptionSelectedWithMessage("Indique el dia")
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	hourStr, err := ScanOptionSelectedWithMessage("Indique la hora")
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	hour, err := strconv.Atoi(hourStr)
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, err
	}

	return time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.UTC), nil
}
func ScanOptionSelected() (string, error) {
	return ScanOptionSelectedWithMessage("")
}

func ScanOptionSelectedWithMessage(message string) (string, error) {
	fmt.Println(message)
	var optionSelected string
	_, err := fmt.Scanln(&optionSelected)

	if err != nil {
		fmt.Println(errorInputMessage)
		return "", err
	}

	return optionSelected, nil
}
