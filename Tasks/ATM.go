package main

import (
	"fmt"
)

type BankAccount struct {
	Owner  string
	Amount float64
}

// Pointer receiver (important)
func (b *BankAccount) deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Deposit amount must be positive")
		return
	}
	b.Amount += amount
	fmt.Println("Deposit successful")
}

func (b *BankAccount) withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be positive")
		return
	}

	if amount > b.Amount {
		fmt.Println("Insufficient balance")
		return
	}

	b.Amount -= amount
	fmt.Println("Withdrawal successful")
}

func main() {
	var owner string
	fmt.Print("Enter account owner name: ")
	fmt.Scanln(&owner)

	account := BankAccount{
		Owner:  owner,
		Amount: 0,
	}

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var amount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scanln(&amount)
			account.deposit(amount)

		case 2:
			var amount float64
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scanln(&amount)
			account.withdraw(amount)

		case 3:
			fmt.Printf("Current balance: %.2f\n", account.Amount)

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
