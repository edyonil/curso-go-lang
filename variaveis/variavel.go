package main

//Quando dois tipos são iguais eu posso declarar na mesma linha
//Exemplo

var nome, sobrenome string

// Escopo de variavel

// Curso Nome do curso do sistema
var Curso string     // Escopo publico, visualizado por todos os pacotes
var faculdade string //Escopo privado, visualizado apenas pelo pacote

//Outro exemplo de declaração eliminando os diversos var
var (
	cidade string
	estado string
)

var endereco string //valor default ""
var cep int         //valor default 0
var ehAtivo bool    //valor default false
var valor float64   //valor default 0.00

func main() {

}
