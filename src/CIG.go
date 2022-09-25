package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func main() {
	// Setup
	var cfgp string
	flag.StringVar(&cfgp, "cfg", "config.json", "Specify the \".cig\" file path")
	help := flag.Bool("help", false, "Shows this menu")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	cfgp, err := filepath.Abs(cfgp)
	if err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}
	parsefile(cfgp)

	GUI()

	// Git functions
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalln("Git repository not found")
	}

	_ = repo
}
