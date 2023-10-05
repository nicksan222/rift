package products

import (
	yaml "rift/utils"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestToYAML(t *testing.T) {
	product := Product{
		Code:            "P12345",
		Name:            "Test Product",
		Description:     "Test Description",
		DefaultPrice:    19.99,
		DiscountedPrice: float64Ptr(17.99),
		Currency:        strPtr("USD"),
		Image:           strPtr("image.jpg"),
		Moq:             intPtr(10),
		Tags:            []string{"tag1", "tag2"},
		Stock:           intPtr(100),
		Availability:    strPtr("In Stock"),
		Weight:          float64Ptr(2.5),
		Dimensions:      strPtr("10x10x10"),
		SKU:             strPtr("SKU12345"),
		Manufacturer:    strPtr("Test Manufacturer"),
		Warranty:        strPtr("1 year"),
		SEOTitle:        strPtr("SEO Test Title"),
		SEODescription:  strPtr("SEO Test Description"),
		SEOKeywords:     []string{"keyword1", "keyword2"},
	}

	yamlData := yaml.ToYAML(product)
	var newProduct Product
	err := yaml.FromYAML(yamlData, &newProduct)

	if err != nil {
		t.Error(err)
	}

	// Check if both products are the same
	if !(pretty.Compare(product, newProduct) == "") {
		t.Logf("Differences: %+v", pretty.Compare(product, newProduct))
		t.Error("The converted Product struct does not match the original.")
	}
}

func TestFromYAML(t *testing.T) {
	yamlData := []byte(`
code: P12345
name: Test Product
description: Test Description
defaultPrice: 19.99
discountedPrice: 17.99
currency: USD
image: image.jpg
moq: 10
tags:
- tag1
- tag2
stock: 100
availability: In Stock
weight: 2.5
dimensions: 10x10x10
SKU: SKU12345
manufacturer: Test Manufacturer
warranty: 1 year
SEOTitle: SEO Test Title
SEODescription: SEO Test Description
SEOKeywords:
- keyword1
- keyword2
`)
	var product Product

	err := yaml.FromYAML(yamlData, &product)

	if err != nil {
		t.Error(err)
	}

	if product.Name != "Test Product" {
		t.Errorf("Expected Name to be 'Test Product', but got %s", product.Name)
	}

	// ... Continue for other fields ...
}

func float64Ptr(f float64) *float64 { return &f }
func strPtr(s string) *string       { return &s }
func intPtr(i int) *int             { return &i }
