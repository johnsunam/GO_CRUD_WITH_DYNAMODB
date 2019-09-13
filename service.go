package main

import (
	"personApp/bindings"

	granitic_yaml "github.com/graniticio/granitic-yaml/v2"
)

//Change to a non-relative path if you want to use 'go install'

func main() {
	granitic_yaml.StartGraniticWithYaml(bindings.Components())
}
