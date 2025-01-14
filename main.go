package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("____Калькулятор ИМТ____")
	userHeight, userKg := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш ИМТ: %.0f\n", IMT)
	fmt.Print(result)
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	IMT := userKg / math.Pow(userHeight, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Print("Введите свой рост (в метрах): ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес (в кг): ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}
