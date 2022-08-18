package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var configuration = &Configuration{}

type Configuration struct {
	Datasource  string              `yaml:"datasource"`
	Lang        string              `yaml:"lang"`
	OrmTag      string              `yaml:"orm-tag"`
	TargetType  string              `yaml:"target-type"`
	TplName     string              `yaml:"tpl-name"`
	TplFile     string              `yaml:"tpl-file"`
	Readonly    bool                `yaml:"readonly"`
	Output      string              `yaml:"output"`
	SkipTables  []string            `yaml:"skip-tables"`
	SkipColumns []SkipColumn        `yaml:"skip-columns"`
	TypeMap     map[string]string   `yaml:"type-map"`
	PackageName string              `yaml:"package-name"`
	TypeMapping map[string]LangType `yaml:"type-mapping"`
}

func init() {
	data, err := os.ReadFile("config.yml")

	if err != nil {
		fmt.Println(err)

		return
	}

	if err = yaml.Unmarshal(data, &configuration); err != nil {
		fmt.Println(err)

		return
	}
}

func GetConfiguration() Configuration {
	return *configuration
}
