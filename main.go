package main

import (
	"fmt"
	"meuprojeto/models" // Substitua "meuprojeto" pelo nome do seu módulo
)

func main() {
	p := models.Pessoa{
		Nome:      "Guilherme",
		Sobrenome: "Lisboa",
		Idade:     20,
		Telfone: models.Telfone{
			DDD:            "11",
			NumeroTelefone: "989724605",
		},
	}

	p.GetFullName()
	fmt.Println()
	p.GetFirstName()
	fmt.Println()
	p.GetLastName()
}
