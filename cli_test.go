package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCreateCliAppEmptyStructs(t *testing.T) {
  app := createCliApp(&StageList{}, &BuildConfig{})
  assert.NotNil(t, app, "App should not be nil")
}

func TestCreateCliAppStagesAdded(t *testing.T) {
  app := createCliApp(&StageList{Stages: []string{"1", "2"}}, &BuildConfig{})
  assert.Equal(t, len(app.Commands), 3, "App should have two commands registered")
  assert.Equal(t, app.Commands[0].Name, "list", "List should be the first command")
  assert.Equal(t, app.Commands[1].Name, "1", "Second command should be called 1")
  assert.Equal(t, app.Commands[2].Name, "2", "Third command should be called 2")
  assert.NotNil(t, app.Commands[1].Action, "Second command has an action")
  assert.NotNil(t, app.Commands[2].Action, "Third command has an action")
}
