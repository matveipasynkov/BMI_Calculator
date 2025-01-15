package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	for {
		fmt.Println("____Калькулятор ИМТ____")
		userHeight, userKg := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			fmt.Println("Вес и/или рост некорректны.")
		} else {
			outputResult(IMT)
		}
		if !checkRepeatCalculation() {
			break
		}
	}
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш ИМТ: %.0f\n", IMT)
	fmt.Print(result)
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела.")
	case IMT <= 18.5:
		fmt.Println("У вас дефицит массы тела.")
	case IMT <= 25:
		fmt.Println("У вас норма массы тела.")
	case IMT <= 30:
		fmt.Println("У вас избыточная масса тела.")
	case IMT <= 35:
		fmt.Println("У вас 1-я степень ожирения.")
	case IMT <= 40:
		fmt.Println("У вас 2-я степень ожирения.")
	default:
		fmt.Println("У вас 3-я степень ожирения.")
	}
}

func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userHeight <= 0 || userKg <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userKg / math.Pow(userHeight, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Print("Введите свой рост (в метрах): ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес (в кг): ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}

func checkRepeatCalculation() bool {
	var input string
	for {
		fmt.Print("Хотите ли вы продолжить работу программы (да / нет)?: ")
		fmt.Scan(&input)
		switch input {
		case "да":
			return true
		case "нет":
			return false
		default:
			fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		}
	}
}
