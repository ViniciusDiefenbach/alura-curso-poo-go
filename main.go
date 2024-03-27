package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// var contaDoGuilherme ContaCorrente = ContaCorrente{} // Zero-Initialization

	contaDoGuilherme := ContaCorrente{
		titular:       "Guilherme",
		numeroAgencia: 589,
		numeroConta:   1234256,
		saldo:         125.5,
	}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDaCris)
}
