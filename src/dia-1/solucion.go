package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func removeNonNumeric(line string) string {
	// Use a regular expression to match only numeric characters
	numericRegex := regexp.MustCompile("[^0-9]")
	result := numericRegex.ReplaceAllString(line, "")

	return result
}

func processLine(line string) int {
	lineWithoutLetters := removeNonNumeric(line)
	number := lineWithoutLetters[0:1] + lineWithoutLetters[len(lineWithoutLetters)-1:]
	num, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	return num
}

func main() {
	total := 0

	// abro el archivo
	file, err := os.Open("src/dia-1/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// Creo el channel por donde mando los resultados

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// por cada linea tiro una task
		line := scanner.Text()
		fmt.Println(line)

		total += processLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Result:", total)

}
