package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years float64
	var expectedReturns float64

	fmt.Print("Enter Investment amount : ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter years to invest : ")
	fmt.Scan(&years)

	fmt.Print("Enter expected return rate : ")
	fmt.Scan(&expectedReturns)

	futureValue := investmentAmount * math.Pow((1+expectedReturns/100), years)
	futureReadValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Investment Amount  : ", investmentAmount)
	fmt.Println("Future value  : ", futureValue)
	fmt.Println("Future real value  : ", futureReadValue)
}
