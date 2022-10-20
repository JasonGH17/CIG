package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
)

func GitLoop() chan struct{} {
	ticker := time.NewTicker(time.Hour)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				for i := 0; i < len(conf.Projects); i++ {
					project := conf.Projects[i]
					log.Printf("Pulling changes - %s", project.Name)
					err := pull(project.Location, project.Name)
					if err != nil {
						log.Printf("Error pulling (%s): %v", project.Name, err)
						continue
					}
					log.Printf("Pulled changes successfully")
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}

func pull(path string, name string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		log.Printf("Error opening git repository in directory: %s\nError msg: %v", path, err)
		return err
	}

	wt, err := repo.Worktree()
	if err != nil {
		log.Printf("Error getting git repository's base worktree: %v", err)
		return err
	}

	err = wt.Pull(&git.PullOptions{})
	if err != git.NoErrAlreadyUpToDate && err != nil {
		log.Printf("Error pulling: %v", err)
		return err
	} else {
		runPullScripts(path, name)
	}

	return nil
}

func runPullScripts(path string, name string) error {
	log.Printf("%s - Running post pull scripts", name)

	confPath := filepath.Join(path, ".cig")
	scriptsPath := filepath.Join(confPath, "scripts.json")

	if _, err := os.Stat(confPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
			return err
		}

		fscripts, err := os.Create(scriptsPath)
		if err != nil {
			log.Println(err)
			return err
		}

		fscripts.Write([]byte("[]"))
	}

	bScripts, err := os.ReadFile(scriptsPath)
	if err != nil {
		log.Println(err)
		return err
	}

	scripts := []string{}
	err = json.Unmarshal(bScripts, &scripts)
	if err != nil {
		log.Println(err)
		return err
	}

	for i := 0; i < len(scripts); i++ {
		script := scripts[i]
		cmd := runScript(path, script)
		err := cmd.Run()
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func runScript(path string, script string) *exec.Cmd {
	switch build_os {
	case "windows":
		cmd := exec.Command("cmd", "/C", script)
		cmd.Dir = path
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd

	case "linux":
		cmd := exec.Command("bash", "-c", script)
		cmd.Dir = path
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd

	default:
		log.Fatalln("Can't run scripts on this machine (Incompatible OS)")
	}

	return nil
}
