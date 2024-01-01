package main

import (
	"fmt"

	"github.com/bank-web-go/accounts"
)

func main() {
	myAccount := accounts.BankAccount{"Nariana", 0000, 111111, 900}
	accountToTransfer := accounts.BankAccount{"Kleber", 0001, 1111112, 500}

	fmt.Println(myAccount.Transfer(950, &accountToTransfer))
	fmt.Println(accountToTransfer.Balance)
}
