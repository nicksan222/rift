package tag

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get current file location")
	}

	// Getting the directory of the current file
	dir := filepath.Dir(filename)

	// Constructing the path to tags.json
	jsonPath := filepath.Join(dir, "tags.json")

	// Reading tags from local file
	file, err := os.ReadFile(jsonPath)

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
