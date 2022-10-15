package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var cfgp string

func main() {
	// Setup
	var prj string

	flag.StringVar(&cfgp, "cfg", "config.json", "Specify the \"config.json\" file path")
	flag.StringVar(&prj, "start", "", "Add the cwd as a new project")
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
	parseFile(cfgp)

	// CLI
	if prj != "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting the cwd: %v\n", err)
		}
		cliNewProject(cwd, prj)

		os.Exit(0)
	}

	// Git
	quit := GitLoop()
	_ = quit

	// Web API
	GUI()
}
