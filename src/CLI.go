package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func cliNewProject(path string, name string) {
	prj := CLIPP{Op: 0, Name: name, Location: path}
	jprj, err := json.Marshal(prj)
	if err != nil {
		log.Fatalf("CLI Error: %v\n", err)
	}

	res, err := http.Post("http://localhost:8000/cli/project", "application/json", bytes.NewReader(jprj))
	if err != nil {
		log.Fatalf("CLI Request failed: %v\n", err)
	}
	if res.StatusCode == 201 {
		log.Printf("Success")
	}
}
