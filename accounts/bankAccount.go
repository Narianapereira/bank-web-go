package accounts

import "github.com/bank-web-go/clients"

type BankAccount struct {
	Owner   clients.Owner
	Agency  int
	Account int
	balance float64
}

func (c *BankAccount) GetCash(cashValue float64) string {
	canGetCash := cashValue > 0 && cashValue <= c.balance
	if canGetCash {
		c.balance -= cashValue
		return "Cashed success"
	} else {
		return "Not enough balance"
	}
}

func (c *BankAccount) DepositCash(cashValue float64) (string, float64) {
	if cashValue > 0 {
		c.balance += cashValue
		return "Deposit success", c.balance
	} else {
		return "Ivalid deposit", c.balance
	}

}

func (c *BankAccount) Transfer(cashValue float64, accountToTransfer *BankAccount) (string, float64) {
	if c.balance >= cashValue && cashValue > 0 {
		c.balance -= cashValue
		accountToTransfer.DepositCash(cashValue)
		return "Transfer success", c.balance
	} else {
		return "Invalid balance", c.balance
	}

}

func (c *BankAccount) GetBalance() float64 {
	return c.balance
}
