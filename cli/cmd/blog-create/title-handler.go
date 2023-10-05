package blog_create

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func InputTitle() (string, error) {
	validate_title := func(input string) error {
		if len(input) < 3 {
			return errors.New("name must be at least 3 characters long")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Choose blog title",
		Validate: validate_title,
	}

	result, err := prompt.Run()

	if err != nil {
		return (""), err
	}

	return result, nil
}
