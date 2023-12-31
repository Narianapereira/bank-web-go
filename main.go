package main

import "fmt"

type bankAccount struct {
	owner   string
	agency  int
	account int
	balance float64
}

func main() {
	myAccount := bankAccount{owner: "Nariana", balance: 1200.00}
	fmt.Println(myAccount)
}
