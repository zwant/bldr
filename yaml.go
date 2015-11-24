package main

import (
	"gopkg.in/yaml.v2"
  "io/ioutil"
  "path/filepath"
)

func parseFromYamlFile(filename string, out interface{}) {
  absFilename, _ := filepath.Abs(filename)
  yamlFile, err := ioutil.ReadFile(absFilename)

  if err != nil {
    panic(err)
  }

  err = yaml.Unmarshal(yamlFile, out)
	if err != nil {
		panic(err)
	}
}
