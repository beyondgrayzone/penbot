# Search

**Import (client component):** `import Search from "@penbot/core/components/search/search.svelte";`

## How it works

1. **Build-time indexing** — Documents are indexed during build
2. **Static API endpoint** — Search data served from a JSON file
3. **Client-side search** — Fast FlexSearch-based interface

## Build script

In your `build-search-data.js`:

```ts
import { defineSearchContent, cleanMarkdown } from "@penbot/core";

export function buildDocsSearchIndex() {
  return defineSearchContent(
    docs.map((doc) => ({
      title: doc.title,
      href: `/docs/v1/${doc.slug}`,
      description: doc.description,
      content: cleanMarkdown(doc.raw),
      category: doc.section,
    }))
  );
}
```

## API Endpoint

```ts
// src/routes/api/search.json/+server.ts
import search from "./search.json" assert { type: "json" };
export const prerender = true;
export const GET = () => Response.json(search);
```

## Client Integration

```svelte
<Search />
```

Supports keyboard shortcut (Cmd/Ctrl + K).

## Search algorithm

- **Primary:** FlexSearch on titles (10pts) and content (5pts)
- **Fallback:** Fuzzy matching (8pts title, 3pts content)
- Results capped at 10, ranked by score
