package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var hasError = false

func main() {
	if err := _main(os.Args[1:]); err != nil {
		log.Panic(err)
	}
}

func _main(args []string) error {
	if len(args) > 1 {
		return errors.New("Usage: glox [script]")
	} else if len(args) == 1 {
		return runFile(args[0])
	} else {
		return runPrompt(os.Stdin)
	}

	return nil
}

func runFile(fileName string) error {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return run(string(fileContents))
}

func runPrompt(stdin io.Reader) error {
	scanner := bufio.NewScanner(stdin)

	fmt.Printf("> ") // for the 1st line
	for scanner.Scan() {
		fmt.Printf("> ")
		run(scanner.Text())
		hasError = false
	}

	return nil
}

func run(input string) error {
	if hasError {
		return errors.New("ERROR")
	}

	return nil
}

func report(lineNumber int, where string, message string) {
	fmt.Println("[line %v] Error %s: %s", lineNumber, where, message)
	hasError = true
}
