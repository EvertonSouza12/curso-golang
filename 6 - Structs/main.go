package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint32
}

func main() {
	var u usuario
	u.nome = "Everton"
	u.idade = 26
	fmt.Println(u)

	enderecoExemplo := endereco{"Nome da rua", 23}
	usuario2 := usuario{"Everton", 26, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Everton"}
	fmt.Println(usuario3)
}