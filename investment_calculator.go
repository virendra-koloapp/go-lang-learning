package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years float64 = 1000, 10
	expectedReturns := 5.5

	futureValue := investmentAmount * math.Pow((1+expectedReturns/100), years)
	futureReadValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Investment Amount  : ", investmentAmount)
	fmt.Println("Future value  : ", futureValue)
	fmt.Println("Future real value  : ", futureReadValue)
}
