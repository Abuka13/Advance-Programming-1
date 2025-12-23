package Wallet

import "fmt"

type Transaction struct {
	Type   string
	Amount float64
}

type Wallet struct {
	Balance      float64
	Transactions []Transaction
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
	w.Transactions = append(w.Transactions, Transaction{
		Type:   "add",
		Amount: amount,
	})
	fmt.Println("Money added")
}

func (w *Wallet) SpendMoney(amount float64) {
	if amount > w.Balance {
		fmt.Println("Not enough balance")
		return
	}

	w.Balance -= amount
	w.Transactions = append(w.Transactions, Transaction{
		Type:   "spend",
		Amount: amount,
	})
	fmt.Println("Money spent")
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}
