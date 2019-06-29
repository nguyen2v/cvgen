package cv

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// ReadFromYamlFile reads YAML file and processes it
func ReadFromYamlFile(filepath string) *CV {
	if data, err := ioutil.ReadFile(filepath); err != nil {
		panic(err)
	} else {
		return processData(data)
	}
}

func processData(data []byte) *CV {
	cv := CV{}

	err := yaml.Unmarshal(data, &cv)

	if err != nil {
		panic(err)
	}

	return &cv
}