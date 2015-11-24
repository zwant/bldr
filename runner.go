package main

import (
  "fmt"
  "log"
	"os/exec"
  "strings"
)

func runCommand(cmd string, config *BuildConfig) {
  cmdToRun, existsInMapping := config.Config[cmd]
  if(!existsInMapping) {
    log.Fatal("Command has no mapping: ", cmd)
  }

  var (
		cmdOut []byte
		err    error
	)

  splitCmd := strings.Split(cmdToRun.Cmd, " ")

  if cmdOut, err = exec.Command(splitCmd[0], splitCmd[1:]...).Output(); err != nil {
    log.Fatal("Failed to execute: ", err)
	}

  fmt.Printf("Success! Output: %v", string(cmdOut))
}
