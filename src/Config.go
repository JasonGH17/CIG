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

var conf CFG

func parsefile(path string) *CFG {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalf("Error on config file parsing: %v\n", err)
	}

	return &conf
}

func savefile(path string) error {
	output, err := json.MarshalIndent(&conf, "", "\t")
	if err != nil {
		return err
	}
	os.WriteFile(path, output, os.ModePerm)
	return nil
}

func getPrjsJson() []byte {
	output, err := json.Marshal(&conf.Projects)
	if err != nil {
		log.Fatalf("Error on config projects json marshalling: %v\n", err)
	}
	return output
}
