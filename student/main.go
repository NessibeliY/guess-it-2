package main

import (
	"bufio"
	"fmt"
	"guess-it-2/student/prediction"
	"log"
	"os"
	"strconv"
)

func main() {
	consoleScanner := bufio.NewScanner(os.Stdin)
	var sliceX, sliceY []int
	count := 0
	for consoleScanner.Scan() {

		input := consoleScanner.Text()
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			log.Print(err)
			return
		}
		if inputInt >= 104 {
			sliceY = append(sliceY, inputInt)
			sliceX = append(sliceX, count)
			count++
		}
		predictedRange := prediction.PredictRange(sliceX, sliceY)
		fmt.Printf("%d %d\n", predictedRange[0], predictedRange[1])
	}
}

func ReadFile(sampleFile string) string {
	data, err := os.ReadFile(sampleFile)
	if err != nil {
		log.Print(err)
	}
	dataStr := string(data)
	return dataStr
}
