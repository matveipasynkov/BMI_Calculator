package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	userHeight, userKg := 1.8, 100.0
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Println(IMT)
}
