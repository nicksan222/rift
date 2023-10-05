package blog

import (
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
)

func TestCreateBlog(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		options  []BlogOption
		expected Blog
	}{
		{
			name:  "Default values",
			title: "Test Blog",
			expected: Blog{
				Title:     "Test Blog",
				Date:      time.Now().Round(time.Second),
				Slug:      "/blog/Test Blog",
				IsVisible: false,
			},
		},
		{
			name:  "With date and slug",
			title: "Test Blog",
			options: []BlogOption{
				WithDate(time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC)),
				WithSlug("/blog/test-blog"),
			},
			expected: Blog{
				Title: "Test Blog",
				Date:  time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC),
				Slug:  "/blog/test-blog",
			},
		},
		// ... other test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blog := CreateBlog(tt.title, tt.options...)
			blog.Date = blog.Date.Round(time.Second) // Round to second to avoid failing due to nanosecond differences
			if diff := pretty.Compare(blog, tt.expected); diff != "" {
				t.Errorf("CreateBlog() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
