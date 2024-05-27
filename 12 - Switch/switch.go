package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Numero invalido"
	}
}

func diaDaSemana2(numero int) string{
	var diadasemana string

	switch {
	case numero == 1:
		diadasemana = "Domingo"
	case numero == 2:
		diadasemana = "Segunda"
	case numero == 3:
		diadasemana = "Terça"
	case numero == 4:
		diadasemana = "Quarta"
	case numero == 5:
		diadasemana = "Quinta"
	case numero == 6:
		diadasemana = "Sexta"
	case numero == 7:
		diadasemana = "Sabado"
	default:
		diadasemana = "Dia invalido"
	}

	return diadasemana
}

func main() {
	dia := diaDaSemana(200)
	fmt.Println(dia)
	
	fmt.Println("--------------")
	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)
}