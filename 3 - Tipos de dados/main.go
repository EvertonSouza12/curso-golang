package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 10000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 1234
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 173
	fmt.Println(numero4)

	var numeroReal float32 = 123.35
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 123000000000000.35
	fmt.Println(numeroReal2)

	numeroReal3 := 123000000000000.35
	fmt.Println(numeroReal3)

	var str string = "Qualquer texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var texto int
	fmt.Println(texto)

	var booleano1 bool = false
	fmt.Println(booleano1)

	// Criar um erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}