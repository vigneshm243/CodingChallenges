package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileNames := os.Args

	if len(fileNames) == 1 {
		log.Fatal("No filenames provided")
	}
	// fmt.Println(fileNames[1])

	tokens := GetTokenList(fileNames[1])
	// for _, token := range tokens {
	// 	fmt.Print(token.Literal + " ")
	// 	fmt.Println(tokenNames[token.Type])
	// }

	parsedJSON, err := parse(tokens)
	if err != nil {
		panic(err)
	}
	if parsedJSON == nil {
		fmt.Println("Invalid JSON. Empty file")
		os.Exit(1)
	}
	// fmt.Println(parsedJSON)
	// fmt.Println(tokens)
}
