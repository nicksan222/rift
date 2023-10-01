import { z, defineCollection, reference } from "astro:content";

const blogCollection = defineCollection({
  type: "content",
  schema: z.object({
    title: z.string(),
    date: z.date().default(() => new Date()),
    slug: z.string(),
    image: z.string().optional(),
    excerpt: z.string().optional(),  // A short summary or teaser for the post.
    content: z.string(),            // The main content of the blog post.
    tags: z.array(z.string()).optional(),
    author: z.string().optional(),
    authorEmail: z.string().optional(),
    relatedProducts: z.array(reference("products")).optional(),
    relatedBundles: z.array(reference("bundles")).optional(),
    isVisible: z.boolean().default(() => true),   // To control the visibility of the post.
    featureOnHomepage: z.boolean().optional(),    // If the post should be featured on the homepage.

    SEOTitle: z.string().optional(),
  SEODescription: z.string().optional(),
  SEOKeywords: z.array(z.string()).optional(),
  }),
});

export default blogCollection;
