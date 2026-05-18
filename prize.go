package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AccessingPrize() {

	reader := bufio.NewReader(os.Stdin)
	ClearScreen()
	var words = []string{
		Color("TO ACCESS YOUR REWARD YOU WILL HAVE TO DROP YOUR BANK DETAILS HERE", "\033[1;33m"),
		Color("FORMAT; <account number> <bankname>", "\033[1;33m"),
		"",
	}
	TypeEffectFast(words)
input:
	fmt.Print(Color("INPUT HERE: ", "\033[1;30m"))

	bankDetails, _ := reader.ReadString('\n')
	bankDetails = strings.TrimSpace(strings.ToLower(bankDetails))
	details := strings.Split(bankDetails, " ")
	if len(details) != 2 {
		fmt.Println(Color("YOU DID NOT FOLLOW THE FORMAT <account number> <bankname>", "\033[1;31m"))
		fmt.Println(Color("TRY AGAIN", "\033[1;31m"))
		fmt.Println()
		goto input
	}

	for _, char := range details[0] {
		if char < '0' || char > '9' {
			fmt.Println(Color("YOUR ACCOUNT NUMBER SHOULD CONTAIN JUST NUMBERS", "\033[1;31m"))
			fmt.Println(Color("TRY AGAIN", "\033[1;31m"))
			fmt.Println()
			goto input
		}
	}

	for _, char := range details[1] {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			fmt.Println(Color("YOUR BANK NAME SHOULD CONTAIN JUST LETTERS", "\033[1;31m"))
			fmt.Println(Color("TRY AGAIN", "\033[1;31m"))
			fmt.Println()
			goto input
		}
	}

	fmt.Printf(Color("OK WE'VE GOT YOUR DETAILS: %v\n", "\033[1;32m"), bankDetails)
	fmt.Println(Color("FOR ANY COMPLAINT CONTACT US AT; 07036702434 or 09133440339", "\033[1;32m"))
	fmt.Println(Color("THANK YOU FOR YOUR TIME", "\033[1;32m"))
	fmt.Println()

	var word = []string{
		Color("SEE YOU NEXT TIME ON RUNTIME RICHES - LEEF CENTER EDITION SERIES 1", "\033[1;96m"),
		"\n",
	}
	TypeEffectFast(word)
}