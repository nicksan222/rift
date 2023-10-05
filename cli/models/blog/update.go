package blog

func (blog *Blog) Update(options ...BlogOption) {
	for _, option := range options {
		option(blog)
	}
}
