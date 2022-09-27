package main

import (
	"encoding/json"
	"log"
	"os"
)

type PRJ struct {
	Name     string `json:"name"`
	Refresh  int64  `json:"refresh"`
	Location string `json:"location"`
}
type CFG struct {
	Projects []PRJ `json:"projects"`
}

var conf CFG

func parseFile(path string) *CFG {
	if _, err := os.Stat(path); err != nil {
		f, _ := os.Create(path)
		f.Close()
	}
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}

	if len(file) == 0 {
		conf = CFG{Projects: []PRJ{}}
		err = saveFile(path)
		if err != nil {
			log.Fatalf("Error on default value config file save: %v\n", err)
		}

		return &conf
	}

	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalf("Error on config file parsing: %v\n", err)
	}

	return &conf
}

func saveFile(path string) error {
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
