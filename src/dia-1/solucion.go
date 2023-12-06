package main

import (
	"fmt"
	"regexp"
)

func removeNonNumeric(input string) string {
	// Use a regular expression to match only numeric characters
	numericRegex := regexp.MustCompile("[^0-9]")
	result := numericRegex.ReplaceAllString(input, "")

	return result
}

func main() {
	fmt.Println("Hello world !!")

	// abro el archivo

	// por cada linea tiro una task

	// voy sumando por channel las sumas parciales

	// cierro el channel y devuelvo el resultado por pantalla
}
