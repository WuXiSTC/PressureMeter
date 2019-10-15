package util

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func GetConf(yamlPath string, out interface{}) {
	yamlFile, err := ioutil.ReadFile(yamlPath)
	LogE(err)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, out)
		LogE(err)
	}
}
