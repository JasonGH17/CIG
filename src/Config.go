package main

import (
	"encoding/json"
	"log"
	"os"
)

type PRJ struct {
	Name     string `json:"name"`
	Refresh  int    `json:"refresh"`
	Location string `json:"location"`
}
type CFG struct {
	Projects []PRJ `json:"projects"`
}

func parsefile(path string) CFG {
	config := CFG{}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Error on config file parsing: %v\n", err)
	}

	return config
}

func savefile(path string, config *CFG) error {
	output, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}
	os.WriteFile(path, output, os.ModePerm)
	return nil
}

func getPrjsJson(config *CFG) string {
	output, err := json.Marshal(&config.Projects)
	if err != nil {
		log.Fatalf("Error on config projects json marshalling: %v\n", err)
	}
	return string(output)
}
