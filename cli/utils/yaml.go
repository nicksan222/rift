package yaml

import (
	"log"

	"gopkg.in/yaml.v3"
)

func ToYAML[T any](value T) []byte {
	result, err := yaml.Marshal(&value)

	if err != nil {
		log.Fatalf("Error turning value of type %T into YAML: %v", value, err)
	}

	return result
}

func FromYAML[T any](data []byte, value T) error {
	err := yaml.Unmarshal(data, &value)
	if err != nil {
		log.Fatalf("Error reading value of type %T from YAML. Check file format: %v", value, err)
		return err
	}

	return nil
}
