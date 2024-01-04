package accounts

import "github.com/bank-web-go/clients"

type SavingsAccount struct {
	Owner     clients.Owner
	Agency    int
	Account   int
	Operation int
	balance   float64
}

func (c *SavingsAccount) GetCash(cashValue float64) string {
	canGetCash := cashValue > 0 && cashValue <= c.balance
	if canGetCash {
		c.balance -= cashValue
		return "Cashed success"
	} else {
		return "Not enough balance"
	}
}

func (c *SavingsAccount) DepositCash(cashValue float64) (string, float64) {
	if cashValue > 0 {
		c.balance += cashValue
		return "Deposit success", c.balance
	} else {
		return "Ivalid deposit", c.balance
	}

}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}
