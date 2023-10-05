package products

type Product struct {
	Code            string   `yaml:"code"`
	Name            string   `yaml:"name"`
	Description     string   `yaml:"description"`
	DefaultPrice    float64  `yaml:"defaultPrice"`
	DiscountedPrice *float64 `yaml:"discountedPrice,omitempty"`
	Currency        *string  `yaml:"currency,omitempty"`
	Image           *string  `yaml:"image,omitempty"`
	Moq             *int     `yaml:"moq,omitempty"`
	Tags            []string `yaml:"tags,omitempty"`
	Stock           *int     `yaml:"stock,omitempty"`
	Availability    *string  `yaml:"availability,omitempty"`
	Weight          *float64 `yaml:"weight,omitempty"`
	Dimensions      *string  `yaml:"dimensions,omitempty"`
	SKU             *string  `yaml:"SKU,omitempty"`
	Manufacturer    *string  `yaml:"manufacturer,omitempty"`
	Warranty        *string  `yaml:"warranty,omitempty"`
	SEOTitle        *string  `yaml:"SEOTitle,omitempty"`
	SEODescription  *string  `yaml:"SEODescription,omitempty"`
	SEOKeywords     []string `yaml:"SEOKeywords,omitempty"`
}
