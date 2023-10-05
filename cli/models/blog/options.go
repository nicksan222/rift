package blog

import "time"

type BlogOption func(*Blog)

func WithTitle(title string) BlogOption {
	return func(b *Blog) {
		b.Title = title
	}
}

func WithDate(date time.Time) BlogOption {
	return func(b *Blog) {
		b.Date = date
	}
}

func WithSlug(slug string) BlogOption {
	return func(b *Blog) {
		b.Slug = slug
	}
}

func WithImage(image string) BlogOption {
	return func(b *Blog) {
		b.Image = image
	}
}

func WithExcerpt(excerpt string) BlogOption {
	return func(b *Blog) {
		b.Excerpt = excerpt
	}
}

func WithTags(tags []string) BlogOption {
	return func(b *Blog) {
		b.Tags = tags
	}
}

func WithAuthor(author string) BlogOption {
	return func(b *Blog) {
		b.Author = author
	}
}

func WithAuthorEmail(authorEmail string) BlogOption {
	return func(b *Blog) {
		b.AuthorEmail = authorEmail
	}
}

func WithRelatedProducts(relatedProducts []int64) BlogOption {
	return func(b *Blog) {
		b.RelatedProducts = relatedProducts
	}
}

func WithRelatedBundles(relatedBundles []int64) BlogOption {
	return func(b *Blog) {
		b.RelatedBundles = relatedBundles
	}
}

func WithFeatureOnHomepage(featureOnHomepage bool) BlogOption {
	return func(b *Blog) {
		b.FeatureOnHomepage = featureOnHomepage
	}
}

func WithSEOTitle(SEOTitle string) BlogOption {
	return func(b *Blog) {
		b.SEOTitle = SEOTitle
	}
}

func WithSEODescription(SEODescription string) BlogOption {
	return func(b *Blog) {
		b.SEODescription = SEODescription
	}
}

func WithSEOKeywords(SEOKeywords []string) BlogOption {
	return func(b *Blog) {
		b.SEOKeywords = SEOKeywords
	}
}

func WithContent(content string) BlogOption {
	return func(b *Blog) {
		b.Content = content
	}
}

func WithVisibility(isVisible bool) BlogOption {
	return func(b *Blog) {
		b.IsVisible = isVisible
	}
}
