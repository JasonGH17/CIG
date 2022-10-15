package main

import (
	"log"
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
					err := Pull(project.Location)
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

func Pull(path string) error {
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
	}

	return nil
}
