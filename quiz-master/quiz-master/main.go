package main

import(
	"encoding/csv"
	"fmt"
	"os"
	"log"
	"bufio"
	"time"
)

func readCsvFile(filePath string) [][]string{
	f, err := os.Open(filePath)
	if err != nil{
		log.Fatal("Unable to read file "+filePath, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err!=nil{
		log.Fatal(err)
	}

	return records
}

func runQuizGame(records [][]string, isGameOn chan string){
	fmt.Println("Let's start the quiz! Get Set Go...!")
	// Loop through reach question
	n := 1
	points := 0
	totalQuestion := len(records)
	for _, element := range records{
		isGameOn <-"Yes"
		que, correctAns := element[0], element[1]
		fmt.Println("Question ", n, ": " , que)
		fmt.Println("Your Answer : ")
		answer := bufio.NewScanner(os.Stdin)
		answer.Scan()
		if answer.Text() == correctAns {
			points++;
			fmt.Println("Hurray, Correct Answer!")
		}else{
			fmt.Println("Sorry, Incorrect answer!")
		}	
	}

	fmt.Println("Your score is ", points, " out of total ", totalQuestion, " questions.")
}

func main(){
	records := readCsvFile("problems.csv")
	// timer := time.Tick(10*time.Second)
	isGameOn := make(chan string)
	go runQuizGame(records, isGameOn)
	gameOver:
		
		for{
			select{
			case <-time.After(5*time.Second):
				fmt.Println("Game Over! You ran out of time.")
				break gameOver

			case <-isGameOn:

			}
		}
	


	
}