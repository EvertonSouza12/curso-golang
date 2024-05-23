package main

import "fmt"

func main() {

	fmt.Println("------------aritmeticos------------------")
	soma := 1 + 2
	subtracao := 1 + 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	resto := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)


	// ATRIBUIÇÂO
	fmt.Println("------------ATRIBUIÇÂO------------------")
	var variavel1 string = "String"
	variavel2 := "String"
	fmt.Println(variavel1, variavel2)

	// Relacionais
	fmt.Println("------------Relacionais------------------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// Logicos
	fmt.Println("------------Logicos------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// Unarios
	fmt.Println("------------Unarios------------------")
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)
}