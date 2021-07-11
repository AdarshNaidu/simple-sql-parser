package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"

	"github.com/AdarshNaidu/simple-sql-parser/sql"
)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Please enter an sql query: ")

	for scanner.Scan() {

		parser := sql.NewParser(strings.NewReader(scanner.Text()))

		result, err := parser.Parse()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Result \n%s", result)
		}

		fmt.Println("\nPlease enter an sql query: ")
		
	}
}

