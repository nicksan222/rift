package yaml

import (
	"log"

	"gopkg.in/yaml.v3"
)

type Model[T any] struct {
	Value T
}

func (m *Model[T]) ToYAML() []byte {
	result, err := yaml.Marshal(&m.Value)

	if err != nil {
		log.Fatalf("Error turning model of type %T into YAML: %v", m.Value, err)
	}

	return result
}

func FromYAML[T any](data []byte) Model[T] {
	model := Model[T]{}

	err := yaml.Unmarshal(data, &model.Value)
	if err != nil {
		log.Fatalf("Error reading model of type %T from YAML. Check file format: %v", model.Value, err)
	}

	return model
}
