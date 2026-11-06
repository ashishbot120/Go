package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

// func (a Account) Describe() string {
// 	return a.Owner + " has " + fmt.Sprint(a.Balance)
// }

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit: amount must be positive, got %.2f", amount)
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	if a.Balance < amount {
		return fmt.Errorf("withdraw: insufficient funds, have %.2f, need %.2f", a.Balance, amount)
	}
	a.Balance -= amount
	return nil
}

// func reset(a *Account) {
// 	a.Balance = 0
// }

func main() {
	//a := Account{Owner: "Ravi"}
	//fmt.Println(a.Owner, a.Balance)
	b := Account{Owner: "Asha", Balance: 100}
	depositErr := b.Deposit(50)
	if depositErr != nil {
		fmt.Println("Deposit error:", depositErr)
	} else {
		fmt.Println("Deposit successful")
	}
	fmt.Println("Current balance:", b.Balance)
	withdrawErr := b.Withdraw(30)
	if withdrawErr != nil {
		fmt.Println("Withdraw error:", withdrawErr)
	} else {
		fmt.Println("Withdraw successful")
	}
	fmt.Println("Current balance:", b.Balance)
	withdrawErr = b.Withdraw(500)
	if withdrawErr != nil {
		fmt.Println("Withdraw error:", withdrawErr)
	} else {
		fmt.Println("Withdraw successful")
	}
	fmt.Println("Current balance:", b.Balance)
	depositErr = b.Deposit(-5)
	if depositErr != nil {
		fmt.Println("Deposit error:", depositErr)
	} else {
		fmt.Println("Deposit successful")
	}
	fmt.Println("Current balance:", b.Balance)
	// fmt.Println(b.Balance)
	// reset(&b)
	// fmt.Println(b.Balance)

}
