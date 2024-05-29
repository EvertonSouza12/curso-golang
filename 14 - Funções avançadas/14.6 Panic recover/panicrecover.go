package main

import "fmt"


func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar execução")
	if r := recover(); r != nil {
		fmt.Println("Recuperado com sucesso")
	}
}

func aluneEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(aluneEstaAprovado(6,6))
	fmt.Println("Pos execução")
}