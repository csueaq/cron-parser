package main

import (
	"fmt"
	"github.com/csueaq/cron-parser/pkg/cronParser"
	"os"
)

func main() {

	parser := cronParser.CreateHandler(cronParser.CreateService())

	err := parser.ProcessUserInput(os.Args[1:])

	if err != nil {
		fmt.Println(fmt.Sprintf("failed to process input: %s", err.Error()))
	}

}
