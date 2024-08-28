package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRENCIA != PARALELISMO
	go escreer("Olá, mundo")
	go escreer("Programando em GO")
	escreer("Teste com 2 goroutine")
}

func escreer(texto string) {

	for {
		fmt.Println(texto)
		time.Sleep((time.Second))
	}
}