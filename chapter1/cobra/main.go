package main

import (
	"log"

	"github.com/willshS/go-tour/chapter1/cobra/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute error: %v", err)
	}
}
