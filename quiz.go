package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Quiz game!")

	fmt.Printf("Enter you name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, Welcome to the game!", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 15 {
		fmt.Println("Yay you can play!")
	}else {
		fmt.Println("Sorry, You cannot play!")
		return
	}

	score := 0
	num_questions := 2

	fmt.Printf("What is better, Macbook or Windows ?")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer + " " + answer2 == "Macbook" {
		fmt.Println("Correct!")
		score ++
	} else if answer + " " + answer2 == "macbook" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many minutes does a full week have? ")
	var minutes uint
	fmt.Scan(&minutes)

	if minutes == 10080 {
		fmt.Println("Correct!")
		score++
	} else { 
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v%%", percent)
	}