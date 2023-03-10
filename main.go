package main

import (
	"log"

	"github.com/geekr-dev/go-cli-app/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
