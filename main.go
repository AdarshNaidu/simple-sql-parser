package main

import (
	"fmt"
	"strings"
	"github.com/AdarshNaidu/simple-sql-parser/sql"
)

func main() {
	fmt.Println("Hello Parser")

	var query string = "Hello"

	scanner := sql.NewScanner(strings.NewReader(query))

	fmt.Println(scanner)

	// print all characters of the input.
	for {
		ch, _, err := scanner.R.ReadRune()

		if err != nil {
			break
		}

		fmt.Println(string(ch))
	}
}

