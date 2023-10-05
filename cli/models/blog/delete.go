package blog

/**
 * DeleteBlog sets the IsVisible property to false
 * For full deletion, delete the file itself
 */
func (blog *Blog) DeleteBlog() {
	blog.IsVisible = false
}
