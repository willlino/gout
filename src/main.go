package main

import "fmt"

func main() {
	resultado := soma(5, 5);

	fmt.Println(resultado);
}

func soma(number1 float32, number2 float32) float32 {
	soma := number1 + number2

	return soma
}