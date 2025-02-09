package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// defining the a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "count Lines")

	// parsing the flag provided by the user
	flag.Parse()

	// Calling the count function to count the number of words
	// or line recieved from the standard input and printing it out

	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, CountLines bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// if the countlines flag is not set, we want to count words so we
	//define the scanner slit type to words (defualt is split by lines)
	if CountLines {
		scanner.Split(bufio.ScanLines)
	} else {
		scanner.Split(bufio.ScanWords)
	}

	//defining a counter
	wc := 0

	// for every word scanned, increment the counter
	for scanner.Scan() {
		text := scanner.Text()
		if CountLines {
			if strings.TrimSpace(text) != "" {
				wc++
			}
		} else {
			wc++
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	// return the total
	return wc
}
