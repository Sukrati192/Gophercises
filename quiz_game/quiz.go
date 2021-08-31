package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the csv file containing quiz questions!\n ")
	inputFile, _ := consoleReader.ReadString('\n')
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Coudn't open input file, reading `problems.csv` file...\n")
		err = nil
		file, err = os.Open("/home/sukratiagarwal/Gophercises/quiz_game/problems.csv")
		if err != nil {
			log.Fatalln("Coudn't open csv file!!")
		}
	}
	r := csv.NewReader(file)
	count := 0
	total := 0
	deadline := 30
	fmt.Printf("You have to complete the quiz in 30 seconds.\nWanna set timer yourself? (if yes, press any key. If no, press ENTER) ")
	input, _ := consoleReader.ReadString('\n')
	if input != "" {
		fmt.Printf("Set the timer for the quiz in seconds!")
		fmt.Scanln(&deadline)
	}
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Failed to parse the provided CSV file.")
	}
	total = len(records)
	rows := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Printf("Press ENTER to start the quiz")
	fmt.Scanln()
	fmt.Printf("Your time starts now...\n")
	c1 := make(chan string, 1)
	go func() {
		for _, i := range rows.Perm(total) {
			fmt.Printf("Question: %s\nAnswer: ", records[i][0])
			ans, _ := consoleReader.ReadString('\n')
			ans = strings.ToUpper(strings.TrimSpace(ans))
			records[i][1] = strings.ToUpper(strings.TrimSpace(records[i][1]))
			if ans == records[i][1] || strings.Contains(ans, records[i][1]) {
				count++
			}
		}
		text := "Wow!! you finished it ahead of time :)\n"
		c1 <- text
	}()
	select {
	case res := <-c1:
		fmt.Printf(res)
	case <-time.After(time.Duration(deadline) * time.Second):
		fmt.Println("\nOops!! Time over :(")
	}
	fmt.Printf("\n**********Yay!! You answered %s questions correct out of %s questions**********\n", strconv.Itoa(count), strconv.Itoa(total))
}
