package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type rating float32

const (
	extraPositive rating = 1.2
	positive      rating = 0.6
	negative      rating = -0.6
	initial       rating = 5.0
	extraNegative rating = -1.2
)

type result struct {
	feedbackDate     string
	feedbackTotal    int
	feedbackPositive int
	feedbackNegative int
	feedbackNeutral  int
}

var customer []rating

func main() {
	f, err := os.Open("./switch/customer-sentiment/data.csv")
	if err != nil {
		exitProgram(err.Error())
	}

	defer f.Close()

	r := bufio.NewReader(f)
	str, err := r.ReadString('\n')
	if err != nil {
		exitProgram(err.Error())
	}

	var feedback result
	feedback.feedbackDate = str

	for {
		str, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if len(str) > 10 {
			feedback.feedbackTotal++

			var customerRating rating
			customerRating = 5.0
			text := strings.Split(str, " ")

			for _, word := range text {
				switch s := strings.Trim(strings.ToLower(word), " ,.,!,?,\t,\n,\r"); s {
				case "pleasure", "impressed", "wonderful", "fantastic", "splendid":
					customerRating += extraPositive
				case "help", "helpful", "thanks", "happy":
					customerRating += positive
				case "not helpful", "sad", "angry", "improve", "annoy":
					customerRating += negative
				case "pathetic", "bad", "worse", "unfortunately", "agitated", "frustrated":
					customerRating += extraNegative
				}
			}

			switch {
			case customerRating > 8.0:
				feedback.feedbackPositive++
			case customerRating >= 4.0 && customerRating <= 8.0:
				feedback.feedbackNeutral++
			case customerRating < 4.0:
				feedback.feedbackNegative++
			}
			customer = append(customer, customerRating)
		}
	}
	feedbackTable(feedback, customer)
}

func exitProgram(s string) {
	log.Fatal("Exiting the program with: ", s)
}

func feedbackTable(f result, c []rating) {
	fmt.Printf("Report for : %s", f.feedbackDate)
	fmt.Printf("Total Customer Review: %d\n", f.feedbackTotal)
	fmt.Printf("Positive Reviews: %d\n", f.feedbackPositive)
	fmt.Printf("Negative Reviews: %d\n", f.feedbackNegative)
	fmt.Printf("Neutral Reviews: %d\n", f.feedbackNeutral)
	fmt.Println("Customer Ratings: ", c)
}
