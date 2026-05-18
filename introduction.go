package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Introduction() {
	ClearScreen()

	fmt.Println("PLEASE INPUT YOUR NAME TO SIGN UP")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("FULLNAME: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println()
	fmt.Printf("HELLO, %v\n\n", Capitalize(name))
	time.Sleep(1 * time.Second)

	lines := []string{
		Color("WELCOME TO RUNTIME RICHES - LEEF CENTER EDITION SERIES 1", "\033[1;36m"),
		Color("...", "\033[1;36m"),
		Color("THIS IS NOT JUST A QUIZ GAME.", "\033[1;36m"),
		Color("IT IS A TEST OF KNOWLEDGE, MEMORY, AND COURAGE.", "\033[1;36m"),
		"",
		Color("OBJECTIVE:", "\033[1;36m"),
		Color("ANSWER 20 QUESTIONS CORRECTLY AND WIN ₦20,000,000", "\033[1;36m"),
		"",
		Color("BUT BE CAREFUL...", "\033[1;36m"),
		Color("ONE WRONG ANSWER CAN END YOUR JOURNEY.", "\033[1;36m"),
		"",
		Color("YOU ALSO HAVE THE OPTION OF WALKING AWAY WITH YOUR CURRENT PRIZE", "\033[1;36m"),
		"",
		Color("GUARANTEED LEVELS:", "\033[1;36m"),
		Color("QUESTION 5  → SAFE POINT", "\033[1;36m"),
		Color("QUESTION 10 → HIGHER SAFE POINT", "\033[1;36m"),
		Color("QUESTION 15 → FINAL SAFE POINT", "\033[1;36m"),
		"",
		Color("LIFELINES AVAILABLE:", "\033[1;36m"),
		Color("1. 50:50 - REMOVE TWO WRONG OPTIONS", "\033[1;36m"),
		Color("2. HINT - SMALL CLUE", "\033[1;36m"),
		Color("3. SUGGESTION - GUIDED HELP", "\033[1;36m"),
		"",
		Color("ONCE USED... THEY CANNOT BE USED AGAIN.", "\033[1;36m"),
		Color("...", "\033[1;36m"),
		Color("TYPE <help> FOR HELP", "\033[36m"),
		Color("ARE YOU READY?", "\033[1;36m"),
	}

	TypeEffectFast(lines)

	time.Sleep(2 * time.Second)

	TypeEffectFast([]string{
		Color("STARTING GAME IN...", "\033[1;33m"),
	})

	time.Sleep(1 * time.Second)
	fmt.Println(Color("3...", "\033[1;33m"))
	time.Sleep(1 * time.Second)
	fmt.Println(Color("2...", "\033[1;33m"))
	time.Sleep(1 * time.Second)
	fmt.Println(Color("1...", "\033[1;33m"))

	time.Sleep(1 * time.Second)

	fmt.Println()
	TypeEffectFast([]string{
		Color("LET THE GAME BEGIN...", "\033[1;33m"),
	})

	time.Sleep(1 * time.Second)
	fmt.Print(Color("PRESS ENTER TO CONTINUE...", "\033[1;36m"))
	reader.ReadString('\n')
}