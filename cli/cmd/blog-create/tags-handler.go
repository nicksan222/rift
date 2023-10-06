package blog_create

import (
	"bytes"
	"errors"
	"rift/models/tag"
	"strings"

	"github.com/manifoldco/promptui"
)

const AddNewLabel = "Aggiungi nuovo"

func inputSingleTag(tags map[string]tag.Tag) (tag.Tag, error) {
	items := make([]string, 0, len(tags))

	for _, tag := range tags {
		items = append(items, tag.Name)
	}

	var result string
	index := -1
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    "Select a tag",
			Items:    items,
			AddLabel: AddNewLabel,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		return tag.Tag{}, err
	}

	if result == AddNewLabel {
		// TODO: add
		return tag.Tag{}, errors.New("new tag addition not implemented")
	}

	return tags[result], nil
}

func InputTags() (string, error) {
	tags, err := tag.GetTags()

	if err != nil {
		return "", err
	}

	selectedTags := make(map[string]tag.Tag, len(tags))

	for {
		tag, err := inputSingleTag(tags)

		if err != nil {
			return "", err
		}

		selectedTags[tag.Name] = tag

		prompt := promptui.Prompt{
			Label:     "Add another tag?",
			IsConfirm: true,
		}

		result, err := prompt.Run()

		if err != nil {
			// Check for the 'n' response
			if err.Error() == "^C" { // The error message for 'n' response is "^C"
				break
			}
			break
		}

		if strings.ToLower(result) == "y" {
			continue // continue the loop if 'y' was pressed
		}
	}

	var buf bytes.Buffer
	for _, tag := range selectedTags {
		buf.WriteString(tag.Name + " ")
	}

	return strings.TrimSpace(buf.String()), nil
}
