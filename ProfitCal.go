package main

import "fmt"

func main() {
	var taxrate, revenue, expenses, ebt, p float64
	fmt.Print("Enter taxrate in %: ")
	fmt.Scanln(&taxrate)
	fmt.Print("Enter the revenue in %: ")
	fmt.Scanln(&revenue)
	fmt.Print("Enter the expenses in %: ")
	fmt.Scanln(&expenses)
	ebt = revenue - expenses
	p = ebt * (1 - taxrate/100)
	fmt.Println("Profit is : ", p)
}
