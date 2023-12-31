package blog

import (
	"time"
)

type Blog struct {
	ID                int64     `yaml:"id"`
	Title             string    `yaml:"title"`
	Date              time.Time `yaml:"date"`
	Slug              string    `yaml:"slug"`
	Image             string    `yaml:"image,omitempty"`
	Excerpt           string    `yaml:"excerpt,omitempty"`
	Content           string    `yaml:"content"`
	Tags              []string  `yaml:"tags,omitempty"`
	Author            string    `yaml:"author,omitempty"`
	AuthorEmail       string    `yaml:"authorEmail,omitempty"`
	RelatedProducts   []int64   `yaml:"relatedProducts,omitempty"` // Assuming related products are referenced by IDs
	RelatedBundles    []int64   `yaml:"relatedBundles,omitempty"`  // Assuming related bundles are referenced by IDs
	IsVisible         bool      `yaml:"isVisible"`
	FeatureOnHomepage bool      `yaml:"featureOnHomepage,omitempty"`
	SEOTitle          string    `yaml:"SEOTitle,omitempty"`
	SEODescription    string    `yaml:"SEODescription,omitempty"`
	SEOKeywords       []string  `yaml:"SEOKeywords,omitempty"`
}
