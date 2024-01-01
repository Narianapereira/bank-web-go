package accounts

type BankAccount struct {
	Owner   string
	Agency  int
	Account int
	Balance float64
}

func (c *BankAccount) GetCash(cashValue float64) string {
	canGetCash := cashValue > 0 && cashValue <= c.Balance
	if canGetCash {
		c.Balance -= cashValue
		return "Cashed success"
	} else {
		return "Not enough balance"
	}
}

func (c *BankAccount) DepositCash(cashValue float64) (string, float64) {
	if cashValue > 0 {
		c.Balance += cashValue
		return "Deposit success", c.Balance
	} else {
		return "Ivalid deposit", c.Balance
	}

}

func (c *BankAccount) Transfer(cashValue float64, accountToTransfer *BankAccount) (string, float64) {
	if c.Balance >= cashValue && cashValue > 0 {
		c.Balance -= cashValue
		accountToTransfer.DepositCash(cashValue)
		return "Transfer success", c.Balance
	} else {
		return "Invalid balance", c.Balance
	}

}
