package main

type BuildConfig struct {
  Config map[string]StageConfig
}

type StageConfig struct {
  Cmd string
}

func parseConfigFromYaml(filename string) *BuildConfig {
  config := BuildConfig{}

  parseFromYamlFile(filename, &config)
  return &config
}
