package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"os"
	"os/exec"
	"time"
)

func cronBody() {
	log.Println("tick")
}

func execBackup() {
	cmd := exec.Command("ls", "-la")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Unable to run backup command: %s", err)
	}
}

func initCron() {
	c := cron.New()
	err := c.AddFunc("* * * * * *", cronBody)

	if err != nil {
		log.Printf("Failed to register cron: %s\n", err)
	}

	c.Start()
}

func main() {
	fmt.Println("influx-cron-backup")

	//initCron()

	execBackup()

	for true {
		time.Sleep(10 * time.Second)
	}
}
