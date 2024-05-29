package main

import "fmt"

func main() {

	retorno := func(texto string) string{
		return fmt.Sprint("Recebido -> %s", texto)
	}("Passando parametro")
	fmt.Println(retorno)
}