package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

func (a Account) Describe() string {
	return a.Owner + " has " + fmt.Sprint(a.Balance)
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func reset(a *Account) {
	a.Balance = 0
}

func main() {
	//a := Account{Owner: "Ravi"}
	//fmt.Println(a.Owner, a.Balance)
	b := Account{Owner: "Asha", Balance: 100}
	b.Deposit(50)
	fmt.Println(b.Balance)
	reset(&b)
	fmt.Println(b.Balance)

}
