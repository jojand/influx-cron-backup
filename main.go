package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type DbConfigRecord struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Cron string `json:"cron"`
}

type Configuration struct {
	Database DbConfigRecord `json:"database"`
}

func parseConfig() {
	jsonFile, err := os.Open("resources/test-config.json")
	if err != nil {
		log.Fatalf("Unable to open configuration file: %s\n", err)
	}
	defer jsonFile.Close()

	jsonFileContents, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Unable to read the configuration file: %s\n", err)
	}

	var config Configuration

	err = json.Unmarshal(jsonFileContents, &config)
	if err != nil {
		log.Fatalf("Unable to parse config JSON: %s", err)
	}

	fmt.Println(config.Database.Name)
	fmt.Println(config.Database.Host)
	fmt.Println(config.Database.Cron)
}

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

	parseConfig()
	//initCron()
	//execBackup()

	//for true {
	//	time.Sleep(10 * time.Second)
	//}
}
