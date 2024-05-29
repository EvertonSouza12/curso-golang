package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	posicao := uint(10)
	fmt.Println(fibonacci(int(posicao)))

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(int(i)))
	}
}