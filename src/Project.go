package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

type CError struct{ msg string }

func (err *CError) Error() string {
	return err.msg
}

func jsDate() int64 {
	return time.Now().UnixMilli()
}

func newProject(path string, name string) error {
	for n := range conf.Projects {
		prj := conf.Projects[n]
		if prj.Name == name {
			return &CError{msg: "A project with this name already exists"}
		}
	}

	confPath := filepath.Join(path, "cig.json")
	prjconf, err := os.Create(confPath)
	if err != nil {
		log.Fatalf("Error creating config file in project directory: %v\n", err)
	}

	_ = prjconf
	//prjconf.Write([]byte{"{auth: \"off\", username: \"\", password: \"\"}"})

	conf.Projects = append(conf.Projects, PRJ{Name: name, Location: path, Refresh: jsDate()})
	return saveFile(cfgp)
}
