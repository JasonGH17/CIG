package main

import (
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

	conf.Projects = append(conf.Projects, PRJ{Name: name, Location: path, Refresh: jsDate()})
	return saveFile(cfgp)
}
