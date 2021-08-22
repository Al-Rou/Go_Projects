package main

import (
	c "TwoFactorVerification/CheckPassword"
	s "TwoFactorVerification/SMS"
	"fmt"
)

func main() {
	/*
		An infinite loop to ask for username and password.
		The loop ends as soon as a correct set of username/password is entered or EXIT is requested.
	*/
	for {
		fmt.Println("Enter your username OR type EXIT to exit: ")
		var inputUser string
		fmt.Scanln(&inputUser)
		if inputUser == "EXIT" {
			return
		} else {
			fmt.Println("Enter your password: ")
			var inputPass string
			fmt.Scanln(&inputPass)
			if c.CheckPass(inputUser, inputPass) {
				break
			} else {
				fmt.Println("The entered username or password is wrong!")
			}
		}
	}
	/*
		Coming out of the above loop means username and password are correctly entered.
		Now, a random code has to be created and sent via SMS to the user.
		Since Twilio account hasn't been created, the code is printed for the program to go on!
	*/
	var sentCode string = s.CreateCode()
	fmt.Println(sentCode)
	/*
		In reality, when there is a Twilio account, this if-block has to check whether s.Send() is
		true, but now it checks the other way around, just for the program to go on!
	*/
	if !s.Send() {
		/*
			An infinite loop to ask for the code.
			The loop breaks in case the correct code is entered or the user types EXIT to get out of
			the program!
		*/
		for {
			fmt.Println("Enter the code OR type EXIT to exit: ")
			var inputCode string
			fmt.Scanln(&inputCode)
			if inputCode == "EXIT" {
				return
			} else {
				if inputCode == sentCode {
					//This block means the user is verified successfully!
					fmt.Println("You are successfully verified")
					break
				} else {
					fmt.Println("The entered code is not correct!")
				}
			}
		}
	} else {
		fmt.Println("The SMS could not be sent!")
	}
}
