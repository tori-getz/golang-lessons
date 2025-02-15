package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	var height float64
	var kg float64
	fmt.Println("__ IMT Calculator __")
	fmt.Print("Введите рост:")
	fmt.Scan(&height)
	fmt.Print("Введите вес:")
	fmt.Scan(&kg)
	imt := CalculateIMT(height, kg)
	OutputResult(imt)
}

func OutputResult(imt float64) {
	result := fmt.Sprintf("ИМТ: %.0f", imt)
	fmt.Print(result)
}

func CalculateIMT(userHeight float64, userKg float64) float64 {
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}
