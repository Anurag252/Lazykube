package cmd

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const configfile = "cheatsheet.yaml"

type KubeCommands struct {
	Commands []KubeCommand `yaml:"commands"`
}
type KubeCommand struct {
	Command Command `yaml:"command"`
}

type Command struct {
	Description string `yaml:"description,omitempty"`
	Cmd         string `yaml:"cmd,omitempty"`
}

func readYaml() *KubeCommands {
	yamlFile, err := os.ReadFile(configfile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := &KubeCommands{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
