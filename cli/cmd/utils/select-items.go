package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type ItemMultipleSelect struct {
	ID         string
	IsSelected bool
}

func SelectItems(selectedPos int, allItems []*ItemMultipleSelect) ([]*ItemMultipleSelect, error) {
	// Always prepend a "Done" ItemMultipleSelect to the slice if it doesn't
	// already exist.
	const doneID = "Done"
	if len(allItems) > 0 && allItems[0].ID != doneID {
		var items = []*ItemMultipleSelect{
			{
				ID: doneID,
			},
		}
		allItems = append(items, allItems...)
	}

	// Define promptui template
	templates := &promptui.SelectTemplates{
		Label: `{{if .IsSelected}}
                    ✔
                {{end}} {{ .ID }} - label`,
		Active:   "→ {{if .IsSelected}}✔ {{end}}{{ .ID | cyan }}",
		Inactive: "{{if .IsSelected}}✔ {{end}}{{ .ID | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "ItemMultipleSelect",
		Items:     allItems,
		Templates: templates,
		Size:      5,
		// Start the cursor at the currently selected index
		CursorPos:    selectedPos,
		HideSelected: true,
	}

	selectionIdx, _, err := prompt.Run()
	if err != nil {
		return nil, fmt.Errorf("prompt failed: %w", err)
	}

	chosenItem := allItems[selectionIdx]

	if chosenItem.ID != doneID {
		// If the user selected something other than "Done",
		// toggle selection on this ItemMultipleSelect and run the function again.
		chosenItem.IsSelected = !chosenItem.IsSelected
		return SelectItems(selectionIdx, allItems)
	}

	// If the user selected the "Done" ItemMultipleSelect, return
	// all selected items.
	var selectedItems []*ItemMultipleSelect
	for _, i := range allItems {
		if i.IsSelected {
			selectedItems = append(selectedItems, i)
		}
	}
	return selectedItems, nil
}
