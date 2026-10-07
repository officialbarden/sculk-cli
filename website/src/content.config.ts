import { defineCollection } from 'astro:content'
import { glob } from 'astro/loaders'

const docsPath = "./src/text/docs"

const docs = defineCollection({
    loader: glob({ base: docsPath, pattern: "**/*.md" })
})

export const collections = { docs }