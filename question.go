package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Question() {
	data, err := os.ReadFile("question.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	maps := make(map[int][]string)
	line := strings.Split(string(data), "\n")

	count := 1
	for i := 0; i+7 < len(line); i += 8 {
		maps[count] = line[i : i+8]
		count++
	}
	question := 0
	reader := bufio.NewReader(os.Stdin)
	usedLifeline := ""

	for i := 1; i <= 20; i++ {
		ClearScreen()
		fmt.Println(Color("          ====================================================================", "\033[1;34m"))
		fmt.Println(Color("                                    RUNTIME RICHES", "\033[1;34m"))
		fmt.Println(Color("          ====================================================================", "\033[1;34m"))
		fmt.Println()
		fmt.Println(Color("---------------------------------", "\033[1;37m"))
		fmt.Printf(Color("        QUESTION %v OF 20\n", "\033[1;37m"), question+1)
		fmt.Println(Color("---------------------------------", "\033[1;37m"))
		TypeEffectFast([]string{
			(Color("FOR THIS QUESTION, YOUR PRIZE MONEY IS; ", "\033[1;36m")),
			fmt.Sprintf((Color("₦%v", "\033[1;32m")), maps[i][6]),
		})
		fmt.Println(Color(maps[i][0], "\033[1;37m"))
		TypeEffectFast([]string{
			fmt.Sprintf("%-20s %-s", maps[i][1], maps[i][2]),
			fmt.Sprintf("%-20s %-s", maps[i][3], maps[i][4]),
		})

	start:
		fmt.Print(Color(" ANSWER: ", "\033[1;36;4m"))
		answer, _ := reader.ReadString('\n')
		answer = strings.ToUpper(strings.TrimSpace(answer))

		if answer == "HELP" {
			AskForHelp()
			goto start
		}

		if answer == "WALKAWAY" {
			if question > 0 {
				TypeEffectFast([]string{
					"",
					fmt.Sprintf(Color("YOU ARE ABOUT TO WALK AWAY WITH YOUR CURRENT PRIZE MONEY ₦%v", "\033[1;31m"), maps[i-1][6]),
					Color("DO YOU STILL WANT TO PROCEED?", "\033[1;31m"),
					"",
					Color("IF YES?", "\033[1;33m"),
				})
				fmt.Print(Color("PRESS ENTER TO CONTINUE...", "\033[1;33m"))
				reader.ReadString('\n')
				AccessingPrize()
				return
			} else {
				TypeEffectFast([]string{
					"",
					Color("YOU HAVE NO CURRENT PRIZE THEREFORE YOU ARE WALKING AWAY EMPTY", "\033[1;31m"),
					Color("DO YOU STILL WANT TO PROCEED?", "\033[1;31m"),
					"",
					Color("IF YES?", "\033[1;33m"),
				})
				fmt.Println(Color("PRESS ENTER TO CONTINUE...", "\033[1;33m"))
				reader.ReadString('\n')
				return
			}
		}
		if answer == "LIFELINE" {
			fmt.Println()
			if (strings.Contains(usedLifeline, "50:50") || strings.Contains(usedLifeline, "1")) &&
				(strings.Contains(usedLifeline, "hint") || strings.Contains(usedLifeline, "2")) &&
				(strings.Contains(usedLifeline, "suggestion") || strings.Contains(usedLifeline, "3")) {

				TypeEffectFast([]string{
					Color("YOU ARE OUT OF LIFELINES", "\033[1;31m"),
					Color("PROCEED TO ANSWER THE QUESTION BY YOURSELF CHOOSE WALKAWAY", "\033[31m"),
					"",
				})

				goto start
			}

			fmt.Print(Color("LIFELINES ARE;\n 1. 50:50\n 2. HINT\n 3. SUGGESTION\n", "\033[1;35m"))
			fmt.Println()
		star:
			fmt.Print(Color("OPTION: ", "\033[1;36m"))

			lifeline, _ := reader.ReadString('\n')
			lifeline = strings.ToLower(strings.TrimSpace(lifeline))
			switch lifeline {
			case "hint", "2":
				err := CheckIfLifelineHasBeenUsed(strings.Contains(usedLifeline, "hint"))
				if err != nil {
					fmt.Println(err)
					fmt.Println()
					goto star
				} else {
					Hints(i)
					usedLifeline += "hint"
					fmt.Println()
					goto start
				}

			case "suggestion", "3":
				err := CheckIfLifelineHasBeenUsed(strings.Contains(usedLifeline, "suggestion"))
				if err != nil {
					fmt.Println(err)
					fmt.Println()
					goto star
				} else {
					CreatorsSuggestions(i)
					usedLifeline += "suggestion"
					fmt.Println()
					goto start
				}

			case "50:50", "1":
				err := CheckIfLifelineHasBeenUsed(strings.Contains(usedLifeline, "50:50"))
				if err != nil {
					fmt.Println(err)
					fmt.Println()
					goto star
				} else {
					FiftyFifty(maps[i], i)
					fmt.Println(strings.Join(maps[i][:1], "\n"))
					TypeEffectFast(maps[i][1:5])
					usedLifeline += "50:50"
					goto start
				}

			default:
				TypeEffectFast([]string{
					Color("NOT A LIFELINE", "\033[1;31m"),
					Color("TYPE <HELP> TO SEE AVAILBLE LIFLINES", "\033[31m"),
					"",
				})
				goto start
			}
		}
		if answer != "A" && answer != "B" && answer != "C" && answer != "D" {
			TypeEffectFast([]string{
				"",
				Color("INVALID ANSWER FORMAT ❌", "\033[1;31m"),
				Color("PLEASE ENTER THE OPTION LETTER (A, B, C OR D).", "\033[1;31m"),
				"",
			})
			goto start
		}

		if answer == maps[i][len(maps[i])-3] {
			TypeEffectFast([]string{
				" ",
				Color("CHECKING ANSWER...", "\033[1;33m"),
			})
			time.Sleep(2 * time.Second)

			TypeEffectFast([]string{
				Color("THE CORRECT ANSWER IS...", "\033[1;33m")})
			time.Sleep(2 * time.Second)
			fmt.Printf(Color("%v\n", "\033[1m"), maps[i][len(maps[i])-3])

			fmt.Println(Color("YOU PASSED✅", "\033[1;32m"))
			fmt.Printf(Color("YOU HAVE WON ₦%v\n", "\033[1;32m"), maps[i][6])
			fmt.Println()
			question++
			time.Sleep(3 * time.Second)
		} else {
			TypeEffectFast([]string{
				" ",
				Color("CHECKING ANSWER...", "\033[1;33m")})
			time.Sleep(3 * time.Second)

			TypeEffectFast([]string{
				Color("THE CORRECT ANSWER IS...", "\033[1;33m")})
			time.Sleep(2 * time.Second)
			fmt.Printf(Color("%v\n", "\033[1m"), maps[i][len(maps[i])-3])

			fmt.Println(Color("YOU FAILED ❌", "\033[1;31m"))

			if question < 5 {
				fmt.Println(Color("YOU HAVE NOT ANSWERED ANY GUARANTEED QUESTION THEREFORE YOU CANNOT WIN A CASH PRIZE", "\033[1;33m"))
				fmt.Println(Color("THANK YOU FOR YOUR TIME. WE HOPE TO SEE YOU AGAIN", "\033[1;33m"))
			}
			if question < 10 && question >= 5 {
				fmt.Println(Color("THE PRIZE FOR YOUR PREVIOUS GUARANTEED LEVEL IS ₦20,000", "\033[1;32m"))
				time.Sleep(3 * time.Second)
				AccessingPrize()
			}
			if question < 15 && question >= 10 {
				fmt.Println(Color("THE PRIZE FOR YOUR PREVIOUS GUARANTEED LEVEL IS ₦500,000", "\033[1;32m"))
				time.Sleep(3 * time.Second)
				AccessingPrize()
			}
			if question < 20 && question >= 15 {
				fmt.Println(Color("THE PRIZE FOR YOUR PREVIOUS GUARANTEED LEVEL IS ₦3,000,000", "\033[1;32m"))
				time.Sleep(3 * time.Second)
				AccessingPrize()
			}

			fmt.Println()
			return
		}
	}

	var words = []string{
		Color("CONGRATULATIONS YOU HAVE WON ₦20,000,000", "\033[1;33m"),
	}
	TypeEffectFast(words)
	AccessingPrize()
}