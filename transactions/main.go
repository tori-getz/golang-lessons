package main

import "fmt"

func main() {
	transactions := [3]int{15}

	transactions[1] = -31

	fmt.Println(transactions)
}
