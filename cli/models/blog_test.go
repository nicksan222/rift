package blog

import (
	"reflect"
	"testing"
	"time"
)

func TestToYAML(t *testing.T) {
	blog := Blog{
		ID:                1,
		Title:             "Test Title",
		Date:              time.Now().Round(time.Second),
		Slug:              "test-slug",
		Image:             "test-image.jpg",
		Excerpt:           "Test excerpt",
		Content:           "Test content",
		Tags:              []string{"tag1", "tag2"},
		Author:            "Test Author",
		AuthorEmail:       "test@email.com",
		RelatedProducts:   []int64{1, 2},
		RelatedBundles:    []int64{1, 2},
		IsVisible:         true,
		FeatureOnHomepage: true,
		SEOTitle:          "SEO Test Title",
		SEODescription:    "SEO Test Description",
		SEOKeywords:       []string{"keyword1", "keyword2"},
	}

	yamlData := blog.ToYAML()
	newBlog := FromYAML(yamlData)

	// Check if both blogs are the same
	if !reflect.DeepEqual(blog, newBlog) {
		t.Log(blog)
		t.Log(newBlog)
		t.Error("The converted Blog struct does not match the original.")
	}
}

func TestFromYAML(t *testing.T) {
	yamlData := []byte(`
id: 1
title: Test Title
date: ` + time.Now().Round(time.Second).Format(time.RFC3339) + `
slug: test-slug
image: test-image.jpg
excerpt: Test excerpt
content: Test content
tags:
- tag1
- tag2
author: Test Author
authorEmail: test@email.com
relatedProducts:
- 1
- 2
relatedBundles:
- 1
- 2
isVisible: true
featureOnHomepage: true
SEOTitle: SEO Test Title
SEODescription: SEO Test Description
SEOKeywords:
- keyword1
- keyword2
`)
	blog := FromYAML(yamlData)

	if blog.Title != "Test Title" {
		t.Errorf("Expected Title to be 'Test Title', but got %s", blog.Title)
	}

	// ... Continue for other fields ...
}
