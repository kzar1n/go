package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	conta := contas.ContaCorrente{}
	conta.NomeTitular = "Cesar"
	conta.NumeroAgencia = 555
	conta.NumeroConta = 123040
	conta.Depositar(1000)

	fmt.Println("Saldo atual:", conta.GetSaldo())
	conta.Sacar(12)
	fmt.Println("Saldo atual:", conta.GetSaldo())

	fmt.Println(conta.Depositar(1000))

	conta2 := contas.ContaCorrente{}
	conta2.NomeTitular = "Zil"
	conta2.NumeroAgencia = 1555
	conta2.NumeroConta = 133040

	fmt.Println("Conta - Saldo atual:", conta.GetSaldo())
	fmt.Println("Conta 2 - Saldo atual:", conta2.GetSaldo())

	conta.Transferir(1000, &conta2)

	fmt.Println("Conta - Saldo atual:", conta.GetSaldo())
	fmt.Println("Conta 2 - Saldo atual:", conta2.GetSaldo())

}
