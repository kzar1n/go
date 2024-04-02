package contas

type ContaCorrente struct {
	NomeTitular   string
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor > 0 && c.saldo >= valor
	if podeSacar {
		c.saldo -= valor
		return "Saque de realizado!"
	} else {
		return "Saque indisponível!"
	}

}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado! saldo atual:", c.saldo
	} else {
		return "Deposito indisponível! saldo atual:", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor > 0 {
		c.Sacar(valor)
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
