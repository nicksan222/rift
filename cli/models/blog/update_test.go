package blog

import (
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
)

func TestUpdate(t *testing.T) {
	tests := []struct {
		name     string
		initial  Blog
		options  []BlogOption
		expected Blog
	}{
		{
			name: "Update date and slug",
			initial: Blog{
				Title: "Test Blog",
				Date:  time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC),
				Slug:  "/blog/test-blog",
			},
			options: []BlogOption{
				WithDate(time.Date(2023, 10, 6, 0, 0, 0, 0, time.UTC)),
				WithSlug("/blog/updated-test-blog"),
			},
			expected: Blog{
				Title: "Test Blog",
				Date:  time.Date(2023, 10, 6, 0, 0, 0, 0, time.UTC),
				Slug:  "/blog/updated-test-blog",
			},
		},
		// ... other test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blog := tt.initial
			blog.Update(tt.options...)
			if diff := pretty.Compare(blog, tt.expected); diff != "" {
				t.Errorf("Update() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
