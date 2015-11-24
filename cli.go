package main

import (
  "fmt"
  "github.com/codegangsta/cli"
)

func buildListCommand(stageList *StageList) cli.Command {
  return cli.Command{
    Name:  "list",
    Usage: "List available stages",
    Action: func(c *cli.Context) {
         fmt.Printf("Basic stages:\n")
         for _, stageName := range stageList.Stages {
           fmt.Printf("- %v\n", stageName)
         }
         fmt.Printf("\nCompound stages:\n")
         for k, v := range stageList.CompoundStages {
           fmt.Printf("- %v\n", k)
           for _, stageName := range v {
             fmt.Printf("\t %v\n", stageName)
           }
           fmt.Printf("\n")
         }
    },
  }
}

func buildStageCommand(stageName string, config *BuildConfig) cli.Command {
  return cli.Command{
    Name:  stageName,
    Usage: stageName,
    Action: func(c *cli.Context) {
         runCommand(c.Command.Name, config)
    },
  }
}

func createCliApp(stageList *StageList, config *BuildConfig) *cli.App {
  app := cli.NewApp()
	app.Name = "bldr"
	app.Usage = "Build stuff"
	app.Commands = []cli.Command{}

  app.Commands = append(app.Commands, buildListCommand(stageList))

	for _, stageName := range stageList.Stages {
		app.Commands = append(app.Commands, buildStageCommand(stageName, config))
	}

  return app
}
