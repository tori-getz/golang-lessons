package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	Loading()

	for {
		CalculateLoop()

		fmt.Print("Хотите повторно измерить ИМТ? (yes|no): ")

		var answer string
		fmt.Scan(&answer)

		if answer == "yes" {
			continue
		} else {
			fmt.Println("Have a nice day!")
			break
		}
	}
}

func CalculateLoop() {
	fmt.Println("__ IMT Calculator __")
	height, kg := GetUserProperties()
	imt, err := CalculateIMT(height, kg)

	if err != nil {
		fmt.Println(err)
		return
	}

	switch {
	case imt < 16:
		fmt.Printf("У вас сильный дефицит массы тела!")
	case imt < 18.5:
		fmt.Printf("У вас дефицит массы тела!")
	case imt < 25:
		fmt.Printf("У вас нормальный вес!")
	case imt < 30:
		fmt.Printf("У вас избыточный вес!")
	case imt < 35:
		fmt.Printf("У вас 1 степень ожирения!")
	case imt < 40:
		fmt.Printf("У вас 2 степень ожирения!")
	default:
		fmt.Printf("У вас 3 степень ожирения!")
	}

	// if imt < 16 {
	// 	fmt.Printf("У вас сильный дефицит массы тела!")
	// } else if imt < 18.5 {
	// 	fmt.Printf("У вас дефицит массы тела!")
	// } else if imt < 25 {
	// 	fmt.Printf("У вас нормальный вес!")
	// } else if imt < 30 {
	// 	fmt.Printf("У вас избыточный вес!")
	// } else if imt < 35 {
	// 	fmt.Printf("У вас 1 степень ожирения!")
	// } else if imt < 40 {
	// 	fmt.Printf("У вас 2 степень ожирения!")
	// } else if imt > 40 {
	// 	fmt.Printf("У вас 3 степень ожирения!")
	// }

	OutputResult(imt)
}

func OutputResult(imt float64) {
	result := fmt.Sprintf("\nИМТ: %.0f", imt)
	fmt.Println(result)
}

func CalculateIMT(userHeight float64, userKg float64) (float64, error) {
	if userHeight <= 0 || userKg <= 0 {
		return 0, errors.New("Введен некорректный вес или рост")
	}
	imt := userKg / math.Pow(userHeight/100, IMTPower)
	return imt, nil
}

func GetUserProperties() (float64, float64) {
	var height float64
	var kg float64
	fmt.Print("Введите рост:")
	fmt.Scan(&height)
	fmt.Print("Введите вес:")
	fmt.Scan(&kg)
	return height, kg
}

func Loading() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Printf("Loading: %d\n", i)
	// }

	// альтернативная запись цикла (weird, но запомним)
	i := 0
	for i <= 12 {
		if i > 10 {
			continue
		}
		if i == 10 {
			fmt.Println("Загрузка завершена!")
			break
		}
		fmt.Printf("Loading: %d\n", i)
		i++
	}
}
