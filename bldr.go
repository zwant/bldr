package main

import (
	"os"
)

func main() {
	stages := parseStagesFromYaml("./stages.yml")
  config := parseConfigFromYaml("./.bldr.yml")

  app := createCliApp(stages, config)
	app.Run(os.Args)
}
