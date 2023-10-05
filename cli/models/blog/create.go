package blog

import "time"

func CreateBlog(Title string, options ...BlogOption) Blog {
	blog := Blog{
		Title:     Title,
		Date:      time.Now().Round(time.Second),
		Slug:      "/blog/" + Title,
		Content:   "",
		IsVisible: false,
	}

	for _, option := range options {
		option(&blog)
	}

	return blog
}
