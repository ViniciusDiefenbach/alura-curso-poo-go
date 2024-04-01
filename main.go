package main

import (
	"fmt"

	"github.com/viniciusdiefenbach/banco/contas"
	// é possível dar apelidos para os import's da seguinte maneira:
	// c "github.com/viniciusdiefenbach/banco/contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	// Se a importação fosse com o apelido "c" - como visto acima, o código ficaria da seguinte maneira:
	// contaDaSilvia := c.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Transferir(200, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
