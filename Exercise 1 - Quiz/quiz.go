package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func readCsvFile(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file " + filePath)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + filePath)
	}

	return records
}

func main() {
	//File Path Setup
	filePath := flag.String("path", "problems.csv", "defines quiz csv path")
	timeLimit := flag.Int("time", 10, "timer for quiz")
	flag.Parse()

	quizFiles := readCsvFile(*filePath)

	//Quiz Greeting
	fmt.Println(fmt.Sprintf("Welcome to the Quiz. Hit the Enter key to start the time. You have %d seconds", *timeLimit))

	//User Start Quiz
	inputRaw := bufio.NewReader((os.Stdin))
	input, err := inputRaw.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	input = strings.TrimSuffix(input, "\r\n")
	if input != "" {
		fmt.Println("Incorrect button to start quiz. Please restart")
		os.Exit(3)
	}

	ticker := time.NewTicker(time.Second * time.Duration(*timeLimit))
	done := make(chan bool)

	//Quiz Start
	correctAnswer := 0

	go func() {
		for i := 0; i < len(quizFiles); i++ {

			//Ask Question
			fmt.Println("Quiz Question " + quizFiles[i][0])

			//User Response
			inputRaw := bufio.NewReader((os.Stdin))
			input, err := inputRaw.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
				return
			}
			input = strings.TrimSuffix(input, "\r\n")

			//Answer Check
			if input == quizFiles[i][1] {
				fmt.Println("Correct! The Answer is " + quizFiles[i][1])
				correctAnswer++
			} else {
				fmt.Println("Incorrect! The Answer is " + quizFiles[i][1])
			}
		}
		done <- true
	}()

	select {
	case <-done:
	case <-ticker.C:
		fmt.Println("Time's Up!")
	}

	fmt.Println(fmt.Sprintf("Quiz Over. Mark is %d of %d", correctAnswer, len(quizFiles)))
}
