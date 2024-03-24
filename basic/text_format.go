package main

import "fmt"

func main() {
	fmt.Printf("This is a value %v \nThis is another value %v \n", 100, 200)
	fmt.Printf("This is a float value %.2f \nThis is another float value %.3f", 100.343, 200.4343)

	formattedValue := fmt.Sprintf("\nyour return is %.2f", 1994.334)

	fmt.Println(formattedValue)
}
