package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestParseStagesFromYaml(t *testing.T) {
  stages := parseStagesFromYaml("./stages.yml")
  assert.NotNil(t, stages, "stages should exist")
  assert.Equal(t, len(stages.Stages), 2, "Stages should have two entries")
  assert.Equal(t, stages.Stages[0], "test1")
  assert.Equal(t, stages.Stages[1], "test2")
}
