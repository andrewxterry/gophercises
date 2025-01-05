package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"time"
) 

type Quiz struct  {
	addends		string 
	sums 		string 
}

func splitCSV (quizfile string) []Quiz  {
//open csv

	file, err := os.Open(quizfile)
	if err != nil {
		fmt.Println("Error opening file", err)
		
	}

	defer file.Close() 

	//create reader
	reader := csv.NewReader(file)
	// read records

	var newQuiz []Quiz;

	for { 

		question, err := reader.Read()
		if err != nil{
			break //
		}
		addend := question[0]
		sum := question[1]
		quiz := Quiz {
			addends:	addend, 
			sums: 		sum,
		}
		newQuiz = append(newQuiz, quiz)
		
	//	fmt.Println(quiz)
		//time.Sleep(2 * time.Second)
		
		
	}
	
	return newQuiz

}


func main() {
	file := ("problem.csv")
	
	quizFile := splitCSV(file)

/*	var score = 0
	for i, question := range quizFile {
		var ans string
		addend := question.addends 
		sum := question.sums
		fmt.Printf("%d): %s= __\n ", i+1, addend)
		fmt.Scanln(&ans)
		if ans != sum {
			fmt.Println("Incorrect")
		} else {
			fmt.Println("Correct")
			score += 1
		}
	}
fmt.Printf("Your score was: %d\n", score)*/	
	score := 0
	timerDuration := 29 * time.Second
	timer := time.NewTimer(timerDuration)
	fmt.Println("You have 30 seconds to complete the quiz!")

		
	for i, q := range quizFile {
		var ans string
		fmt.Printf("%d: %s = ", i+1, q.addends)

		answerCh := make(chan string)

		go func() {
			fmt.Scanln(&ans)
			answerCh <- ans
		}()
		select {
		case <- timer.C:
		//timer channel
			fmt.Println("\nTimes up!")
			fmt.Printf("Your final score is: %d\n", score)
			return
		case userAnswer := <- answerCh:
		// user answer channel
			if userAnswer == q.sums {
				fmt.Println("Correct!")
				score ++
			} else {
				fmt.Println("Incorrect")
			}
		}
	}
	fmt.Printf("\nYou completed this very difficult quiz!\nYour score is: %d\n", score)
}
	
	


//create math problem 
//math logic
//scoring logic
