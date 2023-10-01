import { z, defineCollection, reference } from "astro:content";

const bundlesCollection = defineCollection({
  type: "content",
  schema: z.object({
    name: z.string(),
    description: z.string(),
    image: z.string().optional(),
    price: z.number().optional(),
    moq: z.number().optional(),
    products: z.array(reference("products")),
    category: z.array(reference("categories")).optional(),
    tags: z.array(z.string()).optional(),
    availability: z.string().optional(), // e.g., "in stock", "out of stock", "pre-order"
    weight: z.number().optional(),
    dimensions: z.string().optional(), // e.g., "10x5x2"
    SKU: z.string().optional(),
    expirationDate: z.date().optional(),
    manufacturer: z.string().optional(),
    warranty: z.string().optional(),
    featured: z.boolean().optional(),

    SEOTitle: z.string().optional(),
    SEODescription: z.string().optional(),
    SEOKeywords: z.array(z.string()).optional(),
  }),
});

export default bundlesCollection;
