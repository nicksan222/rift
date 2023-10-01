import { z, defineCollection } from "astro:content";
import blogCollection from "./blog";
import bundlesCollection from "./bundles";
import categoryCollection from "./category";
import productCollection from "./product";

export const collections = {
  blog: blogCollection,
  bundles: bundlesCollection,
  category: categoryCollection,
  product: productCollection,
};
