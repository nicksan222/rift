package category

import (
	"time"
)

type Category struct {
	Title          string    `yaml:"title"`
	Date           time.Time `yaml:"date,omitempty"`
	Slug           string    `yaml:"slug"`
	Image          *string   `yaml:"image,omitempty"`
	ParentCategory *string   `yaml:"parentCategory,omitempty"`
	DisplayOrder   *int      `yaml:"displayOrder,omitempty"`
	Icon           *string   `yaml:"icon,omitempty"`
	Tags           []string  `yaml:"tags,omitempty"`
	IsVisible      bool      `yaml:"isVisible"`
	SEOTitle       *string   `yaml:"SEOTitle,omitempty"`
	SEODescription *string   `yaml:"SEODescription,omitempty"`
	SEOKeywords    []string  `yaml:"SEOKeywords,omitempty"`
}
