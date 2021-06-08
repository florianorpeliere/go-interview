package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

var Configuration struct {
	ServerBindAddress string   `yaml:"serverBindAddress"`
	DatabaseURI       string   `yaml:"databaseURI"`
	ReloadInterval    string   `yaml:"reloadInterval"`
	Currencies        []string `yaml:"currencies"`
}

func init() {
	var name string
	flag.StringVar(&name, "config", "", "Put the name of a .yaml file within the config package.")
	flag.Parse()

	var err error

	if len(name) > 0 {
		log.Println("Config Loaded", name)
		err = populateConfigYAMLFromFile(name)
	}

	if err != nil {
		panic("Could not load config file. Error: " + err.Error() + " Current Path: " + fmt.Sprint(filepath.Abs(name)))
	}
}

func populateConfigYAMLFromFile(path string) (err error) {
	raw, err := getConfigFileRaw(path)
	if err != nil {
		return
	}
	yaml.Unmarshal(raw, &Configuration)
	return
}

func getConfigFileRaw(path string) (raw []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	raw, err = ioutil.ReadAll(file)
	return
}
