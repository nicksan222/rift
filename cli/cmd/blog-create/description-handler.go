package blog_create

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func InputDescription() (string, error) {

	prompt := promptui.Select{
		Label: "Add a description?",
		Items: []string{"Yes", "No", "Generate with AI"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	switch result {
	case "Yes":
		prompt_desc := promptui.Prompt{
			Label: "Description",
			Validate: func(input string) error {
				if len(input) < 3 {
					return errors.New("description must be at least 3 characters long")
				}
				return nil
			},
		}

		result, err := prompt_desc.Run()

		if err != nil {
			return "", err
		}

		return result, nil

	case "No":
		return "", nil

	case "Generate with AI":
		return "TODO", nil
	}

	return "", nil
}
