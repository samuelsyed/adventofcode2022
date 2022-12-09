package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// open file to read

	file, err := os.Open("calories.txt")

	if err != nil {

		log.Fatal(err)
	}

	defer file.Close()

	// Create a scanner to read the file line by line

	scanner := bufio.NewScanner(file)

	// Keep track of the maximum total and current total

	var maxTotal, currentTotal int

	// Read the file line by line

	for scanner.Scan() {

		// Get the next line of the file

		line := scanner.Text()

		// If the line is empty we need to calculate the total hitherto and update max total

		if line == "" {

			if currentTotal > maxTotal {

				maxTotal = currentTotal

			}
			currentTotal = 0
			continue
		}

		// Convert the string to integer

		num, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		// Add the number to the current total

		currentTotal += num

	}
	fmt.Println("Maximum total:", maxTotal)

}
