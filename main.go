package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var (
	fileLocation string
	number int64
)

func init() {
	flag.StringVar(&fileLocation, "f", "", "location of file to use (Required)")
	flag.Int64Var(&number, "n", 4, "number of entries required to split into new list")
}

func main() {
	flag.Parse()

	if fileLocation == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.ToLower(scanner.Text()))
	}

	nameMap := make(map[string]int64)

	for _, line := range lines {
		if _, ok := nameMap[line]; ok {
			nameMap[line] += 1
		} else {
			nameMap[line] = 1
		}
	}
	var listOfNumber []string
	var listOfRemaining []string
	for name, times := range nameMap {
		if times >= number {
			listOfNumber = append(listOfNumber, strings.Title(name))
		} else {
			listOfRemaining = append(listOfRemaining, strings.Title(name))
		}
	}
	slices.Sort(listOfNumber)
	slices.Sort(listOfRemaining)
	fmt.Printf("Found %d names with %d entries\n", len(listOfNumber), number)
	fmt.Printf("Remaining list size: %d\n", len(listOfRemaining))

	if err := os.WriteFile("ListOfMatches.txt", []byte(strings.Join(listOfNumber, "\n")), 0666); err != nil {
		log.Fatal("failed to write ListOfMatches.txt")
	}
	fmt.Printf("entries with %d matches written to ListOfMatches.txt\n", number)
	if err := os.WriteFile("ListOfRemaining.txt", []byte(strings.Join(listOfRemaining, "\n")), 0666); err != nil {
		log.Fatal("failed to write ListOfRemaining.txt")
	}
	fmt.Printf("entries remaing written to ListOfRemaining.txt")
}
