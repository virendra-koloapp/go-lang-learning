package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax * (1 - taxRate/100)

	fmt.Println("Earning before tax : ", earningBeforeTax)
	fmt.Println("Profit : ", profit)

}
