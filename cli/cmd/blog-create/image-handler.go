package blog_create

import (
	"errors"
	"net/url"

	"github.com/manifoldco/promptui"
)

func InputImage() (string, error) {
	prompt := promptui.Select{
		Label: "Insert an image?",
		Items: []string{"Insert from local file", "From url"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	switch result {
	case "Insert from local file":
		prompt_desc := promptui.Prompt{
			Label: "Local path to image",
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

	case "From url":
		prompt_desc := promptui.Prompt{
			Label: "Url to image",
			Validate: func(input string) error {
				// Should include a regex to validate url
				_, err := url.Parse(input)

				if err != nil {
					return errors.New("invalid url")
				}

				return nil
			},
		}

		result, err := prompt_desc.Run()

		if err != nil {
			return "", err
		}

		return result, nil

	}

	return "", nil
}
