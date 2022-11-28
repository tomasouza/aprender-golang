package main

import "fmt"

func main() {
	a := 20
	b := 10
	resultadoSoma := somar(a, b)
	fmt.Println("Resultado da soma:", resultadoSoma)
	resultadoSubtracao := subtrair(a, b)
	fmt.Println("Resultado da subtracao:", resultadoSubtracao)
	resultadoMultiplicacao := multiplicar(a, b)
	fmt.Println("Resultado da multiplicacao:", resultadoMultiplicacao)
	resultadoDivisao := dividir(a, b)
	fmt.Println("Resultado da divisao:", resultadoDivisao)
}

func somar(a, b int) int {
	return a + b
}

func subtrair(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}
func dividir(a, b int) int {
	return a / b
}
