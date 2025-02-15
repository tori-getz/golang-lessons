package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var height float64
	var kg float64
	fmt.Println("__ IMT Calculator __")
	fmt.Print("Введите рост:")
	fmt.Scan(&height)
	fmt.Print("Введите вес:")
	fmt.Scan(&kg)
	IMT := kg / math.Pow(height, IMTPower)
	OutputResult(IMT)
}

func OutputResult(imt float64) {
	result := fmt.Sprintf("ИМТ: %.0f", imt)
	fmt.Print(result)
}
