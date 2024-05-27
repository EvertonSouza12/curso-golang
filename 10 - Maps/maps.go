package main

import "fmt"

func main() {

	usuario := map[string]string {
		"Nome":      "Everton",
		"Sobrenome": "Souza",
	}

	// fmt.Println(usuario)
	// fmt.Println(usuario["Nome"])
	fmt.Println(usuario["Sobrenome"])

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro" : "Joao",
			"segundo" : "Pedro",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string {
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)
}