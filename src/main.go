package main

import "fmt"

func main() {
	resultado := somar(5, 5);

	fmt.Println(resultado);
}

func somar(number1 float32, number2 float32) float32 {
	soma := number1 + number2

	return soma
}