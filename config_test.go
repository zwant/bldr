package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestParseConfigFromYaml(t *testing.T) {
  config := parseConfigFromYaml("./.bldr.yml")
  assert.NotNil(t, config, "Config should exist")
  assert.Equal(t, len(config.Config), 2, "Config map should have two entries")
  assert.Equal(t, config.Config["test1"].Cmd, "ls -la")
  assert.Equal(t, config.Config["test2"].Cmd, "kossa")
}
