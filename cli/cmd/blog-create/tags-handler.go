package blog_create

import (
	"fmt"
	cmd "rift/cmd/utils"
	"rift/models/tag"
	"strings"
)

func InputTags() (string, error) {
	tags, err := tag.GetTags()

	fmt.Println("Select tags for this blog post:")
	fmt.Println("")
	fmt.Println(tags)

	if err != nil {
		return "", nil
	}

	var tagsElements []*cmd.ItemMultipleSelect

	for _, tag := range tags {
		tagsElements = append(tagsElements, &cmd.ItemMultipleSelect{
			ID:         tag.Name,
			IsSelected: false,
		})
	}

	tagsElements, err = cmd.SelectItems(0, tagsElements)
	if err != nil {
		return "", err
	}

	var selectedTags []string
	for _, tag := range tagsElements {
		if tag.IsSelected {
			selectedTags = append(selectedTags, tag.ID)
		}
	}

	// TODO: Save this tags in the blog yaml

	return strings.Join(selectedTags, ","), nil
}
