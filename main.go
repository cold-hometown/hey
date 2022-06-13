package main

import (
	"log"
	"os"

	"github.com/cold-hometown/hey/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
