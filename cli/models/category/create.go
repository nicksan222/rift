package category

func CreateCategory(Name string, options ...CategoryOption) Category {
	category := Category{
		Name: Name,
	}

	for _, option := range options {
		option(&category)
	}

	return category
}
