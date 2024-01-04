package main

import (
	"fmt"

	"github.com/bank-web-go/accounts"
	"github.com/bank-web-go/clients"
)

func main() {
	myClient := clients.Owner{Name: "Nariana", CPF: "08731393993", Occupation: "Desenvolvdora"}
	myAccount := accounts.SavingsAccount{Owner: myClient, Agency: 0001,
		Account: 12321323, Operation: 10}

	fmt.Println(myAccount)

}
