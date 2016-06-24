package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const configFile = "/Users/ilja/go/config.json"

type config struct {
	Version   string
	Mtime     string
	Instances []instancesConfig
}

type instancesConfig struct {
	InstanceID string
	CronRule   string
	Notify     string
	Comment    string
	KeepDay    int
	KeepWeek   int
	KeepMonth  int
}

func (c *config) readConfig() []byte {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Unable to read config file ", configFile)
	}
	return data
}

func (c *config) writeConfig(data []byte) bool {
	status := false
	err := ioutil.WriteFile(configFile, data, 0644)
	if err != nil {
		fmt.Println("Unable to write config file ", configFile, err)
	} else {
		status = true
	}
	return status
}

func (c *config) statConfig() string {
	info, err := os.Stat(configFile)
	if err != nil {
		fmt.Println("Unable to stat() config file, check permissions ", configFile, err)
	}
	statInfo := info.ModTime()
	return statInfo.Format("2016-01-02 15:04:05")
}

// func validateConfig

func main() {

	cfg := new(config)

	err := json.Unmarshal(cfg.readConfig(), &cfg)
	if err != nil {
		fmt.Println("error:", err)
	}

	newItem := instancesConfig{
		InstanceID: "i-b113a8733",
		CronRule:   "11 1 6 * * * *",
		Notify:     "me@me.com",
		Comment:    "awdawdonfig",
		KeepDay:    3,
		KeepWeek:   4,
		KeepMonth:  1}

	cfg.Instances = append(cfg.Instances, newItem)
	cfg.Mtime = string(time.Now().Format(time.RFC3339))
	fmt.Println(cfg.Mtime)
	b, err := json.Marshal(cfg)
	if err != nil {
		fmt.Println("error:", err)
	}
	//	os.Stdout.Write(b)
	cfg.writeConfig(b)
}

