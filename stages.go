package main

type StageList struct {
	Stages []string
  CompoundStages map[string][]string
}

func parseStagesFromYaml(filename string) *StageList {
  stages := StageList{}

  parseFromYamlFile(filename, &stages)

  return &stages
}
