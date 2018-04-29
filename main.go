package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "", "Path to configuration file")
	flag.Parse()
	if configFile == "" {
		log.Fatal("Usage: flo -config FILE")
	}
	configBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	workingDir, _ := filepath.Split(configFile)

	type Config struct {
		Name  string `yaml:"name"`
		Tasks []struct {
			Name    string `yaml:"name"`
			Command struct {
				Name string   `yaml:"name"`
				Args []string `yaml:"args"`
			} `yaml:"command"`
			Inputs  []string `yaml:"inputs"`
			Outputs []string `yaml:"outputs"`
		} `yaml:"tasks"`
	}

	var config Config
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		panic(err)
	}

	for _, task := range config.Tasks {
		for _, output := range task.Outputs {
			err := os.Mkdir(filepath.Join(workingDir, output), os.ModePerm)
			if err != nil {
				panic(err)
			}
		}

		cmd := exec.Command(task.Command.Name, task.Command.Args...)
		cmd.Dir = workingDir
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}

}
