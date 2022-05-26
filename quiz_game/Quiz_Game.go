/*
QUIZ GAME:
Features:
	1. There will be addition/subtraction/multiplication/division questions that the user
	must answer.
	2. Quiz will not move onto next question if user does not answer
	3. Quiz moves onto next question if user answers, regardless of whether they answered
	correctly or not.
	4. After all questions have been asked, show the user how many questions they got right
	out of the total # of questions.


Requirements:
- Will need something to parse csv file (to read the questions/answers from the quiz)
- Will need to figure out how to take in user input
*/

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	//"strconv"
)

// structure to store questions and answers
type QA struct {
	question string
	answer   string
}

func createQAList(data [][]string) []QA {
	var qaList []QA
	for _, line := range data {
		var qa QA
		for j, field := range line {
			if j == 0 {
				qa.question = field
			} else {
				qa.answer = field
			}
		}
		qaList = append(qaList, qa)
	}

	return qaList
}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil { // this if statement runs if the err is not null, meaning that there is an error
		log.Fatal(err) // fatal is similar to print, follows a call to os.exit and the program terminates
	}

	// closing the file after opening
	defer file.Close() // defer means that the file is closed after everything else in main runs

	// read the csv files
	// NewReader just returns a new Reader object
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// converting the read data to a slice of structs
	qaList := createQAList(data)
	//fmt.Println(qaList)
	points := 0
	for i := 0; i < len(qaList); i++ {
		var ans string
		fmt.Printf("%v: ", qaList[i].question)
		//fmt.Println("Scanning")
		fmt.Scanf("%s\n", &ans)
		if ans == qaList[i].answer {
			points++
		}
	}
	fmt.Printf("Game Over! Your score is: %v/%v", points, len(qaList))
	//fmt.Println(qaList)

}
