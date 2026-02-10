package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------------")
		panic("Can't continue, sorry!")
	}

	fmt.Println("Welcome to Go bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	//for i := 0; i < 200; i++ {
	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice:")
		fmt.Scan(&choice)

		//wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Your deposit amount is less than zero")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdrawal amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Your deposit amount is less than zero")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("You can not withdraw more than you have")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for visiting Go Bank!")
			return
		}
	}
}
