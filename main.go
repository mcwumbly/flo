package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "", "Path to configuration file")
	flag.Parse()
	if configFile == "" {
		log.Fatal("Usage: flo -config FILE")
	}
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config: %s", err)
	}
	fmt.Println(string(bytes))
}
