package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kylelemons/go-gypsy/yaml"
	"gopkg.in/gcfg.v1"
)

type configuration struct {
	Enabled bool
	Path    string
}

func readJsonConfig() {
	file, _ := os.Open("application.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
}

func readYamlConfig() {
	config, err := yaml.ReadFile("application.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}

func readIniConfig() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}
	err := gcfg.ReadFileInto(&config, "application.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}
