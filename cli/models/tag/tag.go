package tag

import (
	"encoding/json"
	"fmt"
	"os"
)

type Tag struct {
	Name string `json:"name"`
}

func CreateTag(name string) (Tag, error) {
	// Does this tag exist?
	// If not, create it

	Tags, err := GetTags()

	if err != nil {
		return Tag{}, err
	}

	Tags[name] = Tag{
		Name: name,
	}

	// Saving locally the tag
	to_save, err := json.MarshalIndent(Tags, "", "	")

	if err != nil {
		return Tag{}, err
	}

	err_write := os.WriteFile("./tags.json", to_save, os.ModePerm)

	if err_write != nil {
		return Tag{}, err_write
	}

	return Tags[name], nil
}

func GetTags() (map[string]Tag, error) {
	// Reading tags from local file
	file, err := os.ReadFile("./tags.json")

	fmt.Println(file)

	if err != nil {
		return nil, err
	}

	Tags := make(map[string]Tag)

	err = json.Unmarshal(file, &Tags)

	if err != nil {
		return nil, err
	}

	return Tags, nil
}
