package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var expectedReturns = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturns/100), float64(years))

	fmt.Println("Investment Amount  : ", investmentAmount)
	fmt.Println("Future value  : ", futureValue)
}
