package models

import "fmt"

type Telfone struct {
	DDD            string
	NumeroTelefone string
}

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Telfone   Telfone
}

func (p Pessoa) GetFullName() {
	fmt.Printf("%s %s.", p.Nome, p.Sobrenome)
}

func (p Pessoa) GetFirstName() {
	fmt.Printf("%s.", p.Nome)
}
func (p Pessoa) GetLastName() {
	fmt.Printf("%s.", p.Sobrenome)
}
