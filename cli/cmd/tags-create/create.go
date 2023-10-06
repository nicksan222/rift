package tags_create

import (
	"errors"

	"rift/models/tag"

	"github.com/manifoldco/promptui"
)

func inputTitle() (string, error) {
	validateTitle := func(input string) error {
		if len(input) < 3 {
			return errors.New("name must be at least 3 characters long")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Choose blog title",
		Validate: validateTitle,
	}

	result, err := prompt.Run()

	if err != nil {
		return (""), err
	}

	return result, nil
}

func InputTitleAndCreate() (string, error) {
	Tag, err := inputTitle()

	if err != nil {
		return "", err
	}

	tagCreated, err := tag.CreateTag(Tag)

	if err != nil {
		return "", err
	}

	return tagCreated.Name, nil
}
