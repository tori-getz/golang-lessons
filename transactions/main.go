package main

import "fmt"

func main() {
	fmt.Println("__ transactions app __")

	transactions := []int{}

	for {
		input := GetUserTransaction()

		if input == 0 {
			break
		}

		transactions = append(transactions, input)
	}

	fmt.Printf("Баланс: %v", SummarizeTransactions(transactions))
}

func GetUserTransaction() int {
	var userInput int

	fmt.Print("Введите транзакцию (n для выхода): ")
	fmt.Scan(&userInput)

	return userInput
}

func SummarizeTransactions(transactions []int) int {
	sum := 0

	for _, transaction := range transactions {
		sum = sum + transaction
	}

	return sum
}
