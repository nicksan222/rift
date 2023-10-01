import { z, defineCollection } from "astro:content";

const productCollection = defineCollection({
  type: "content",
  schema: z.object({
    code: z.string(),
    name: z.string(),
    description: z.string(),
    defaultPrice: z.number(),
    discountedPrice: z.number().optional(),
    currency: z.string().optional(),
    image: z.string().optional(),
    moq: z.number().optional(),
    tags: z.array(z.string()).optional(),
    stock: z.number().optional(),
    availability: z.string().optional(),
    weight: z.number().optional(),
    dimensions: z.string().optional(),
    SKU: z.string().optional(),
    manufacturer: z.string().optional(),
    warranty: z.string().optional(),

    SEOTitle: z.string().optional(),
    SEODescription: z.string().optional(),
    SEOKeywords: z.array(z.string()).optional(),
  }),
});

export default productCollection;
