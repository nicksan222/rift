package category

import (
	yaml "rift/utils"
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
)

func TestToYAML(t *testing.T) {
	category := Category{
		Title:          "Test Category",
		Date:           time.Now().Round(time.Second),
		Slug:           "test-slug",
		IsVisible:      true,
		DisplayOrder:   intPtr(1),
		SEOTitle:       strPtr("SEO Test Title"),
		SEODescription: strPtr("SEO Test Description"),
		SEOKeywords:    []string{"keyword1", "keyword2"},
	}

	yamlData := yaml.ToYAML(category)
	var newCategory Category
	err := yaml.FromYAML(yamlData, &newCategory)

	if err != nil {
		t.Error(err)
	}

	// Check if both categories are the same
	if !(pretty.Compare(category, newCategory) == "") {
		t.Logf("Differences: %+v", pretty.Compare(category, newCategory))
		t.Error("The converted Category struct does not match the original.")
	}
}

func TestFromYAML(t *testing.T) {
	yamlData := []byte(`
title: Test Category
date: ` + time.Now().Round(time.Second).Format(time.RFC3339) + `
slug: test-slug
isVisible: true
displayOrder: 1
SEOTitle: SEO Test Title
SEODescription: SEO Test Description
SEOKeywords:
- keyword1
- keyword2
`)
	var category Category

	err := yaml.FromYAML(yamlData, &category)

	if err != nil {
		t.Error(err)
	}

	if category.Title != "Test Category" {
		t.Errorf("Expected Title to be 'Test Category', but got %s", category.Title)
	}

	// ... Continue for other fields ...
}

// Helper functions to get pointers
func intPtr(i int) *int       { return &i }
func strPtr(s string) *string { return &s }
