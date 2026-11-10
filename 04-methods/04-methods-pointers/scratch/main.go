package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

func (a Account) DepositCopy(amount float64) {
	a.Balance = a.Balance + amount
}

func (a *Account) Deposit(amount float64) {
	a.Balance = a.Balance + amount
}

func (a *Account) Withdraw(amount float64) bool {
	if a.Balance >= amount {
		a.Balance = a.Balance - amount
		return true
	}
	return false
}

func main() {
	acc := Account{"An", 100}
	acc.DepositCopy(5)
	fmt.Println(acc)
	acc.Deposit(20)
	fmt.Println(acc)

	accounts := []Account{{"An", 100}, {"Binh", 200}}

	for _, a := range accounts {
		a.Deposit(10) // không thay đổi balance gốc thì dùng range thì a chỉ là bản sao
	}
	fmt.Println(accounts)

	for i := range accounts {
		accounts[i].Deposit(10) // balance thay đổi vì đang dùng đến chỉ số đến slice gốc
	}
	fmt.Println(accounts)

	acc2 := Account{"Duy", 200}
	acc2.Withdraw(30)
	fmt.Println(acc2.Balance)
	acc2.Withdraw(1000)
	fmt.Println(acc2.Balance)
}