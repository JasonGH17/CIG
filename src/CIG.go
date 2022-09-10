package main

import (
	"log"

	"github.com/go-git/go-git/v5"
)

func main() {
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalln("Git repository not found")
	}

	_ = repo
}
