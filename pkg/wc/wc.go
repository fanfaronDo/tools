package wc

import (
	"bufio"
	"fmt"
	"os"
)

func Reader(file *os.File, filename string, flagW bool, flagL bool, flagM bool, useFile bool) {
	var counter int = 0
	scanner := bufio.NewScanner(file)

	if flagW {
		scanner.Split(bufio.ScanWords)
		counter = amount(scanner)
	} else if flagM {
		scanner.Split(bufio.ScanRunes)
		counter = amount(scanner)
	} else if flagL {
		scanner.Split(bufio.ScanLines)
		counter = amount(scanner)
	}

	if useFile {
		defer file.Close()
		fmt.Println(counter, "\t", filename)
	} else {
		fmt.Println(counter)
	}
}

func amount(scanner *bufio.Scanner) int {
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}
