package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Calling the count function to count the number of words
	// recieved from the standard input and printing it out

	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	//define the scanner slit type to words (defualt is split by lines)
	scanner.Split(bufio.ScanWords)

	//defining a counter
	wc := 0

	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}
	// return the total
	return wc
}
