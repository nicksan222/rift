import { z, defineCollection, reference } from "astro:content";

const categoryCollection = defineCollection({
  type: "content",
  schema: z.object({
    title: z.string(),
    date: z.date().default(() => new Date()),
    slug: z.string(),
    image: z.string().optional(),
    parentCategory: z.string().optional(),
    displayOrder: z.number().optional(),
    icon: z.string().optional(),
    tags: z.array(z.string()).optional(),
    isVisible: z.boolean().default(() => true),

    SEOTitle: z.string().optional(),
    SEODescription: z.string().optional(),
    SEOKeywords: z.array(z.string()).optional(),
  }),
});

export default categoryCollection;
