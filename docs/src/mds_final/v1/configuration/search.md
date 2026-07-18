# Search
> Learn how to customize the search in your Penbot project.


The search functionality provides fast, client-side search across your documentation using FlexSearch for indexing and fuzzy matching.

## How It Works

The search system consists of three main parts:

1. **Build-time indexing** - Documents are processed and indexed during the build via a script
2. **Static API endpoints** - Versioned search data is served from per-version static JSON files
3. **Client-side search** - Fast, interactive search interface using FlexSearch

## Build Process

During the build process, the `docs/scripts/build-search-data.js` script generates versioned search indexes. This script handles multiple documentation versions automatically:

```js
import { fileURLToPath } from "node:url";
import { writeFileSync, mkdirSync } from "node:fs";
import { resolve, join } from "node:path";
import { docs } from "../.velite/index.js";
import { defineSearchContent, cleanMarkdown } from "@penbot/core/search";

const __dirname = fileURLToPath(new URL(".", import.meta.url));
const API_BASE_DIR = resolve(__dirname, "../src/routes/api");
const KIT_DIR = resolve(__dirname, "../../packages/core/src/lib/components/layout");

const versions = [...new Set(docs.map((doc) => doc.slug.split("/")[0]))];
const baseURL = "/docs";

function generateSearchData(filteredDocs) {
	return defineSearchContent(
		filteredDocs.map((doc) => ({
			title: doc.title,
			href: `${baseURL}/${doc.slug}`,
			description: doc.description,
			content: cleanMarkdown(doc.raw),
			category: doc.section,
		})),
	);
}

versions.forEach((version) => {
	// Create the directory: e.g., src/routes/api/v1.search.json/
	const versionDir = join(API_BASE_DIR, `${version}.search.json`);
	mkdirSync(versionDir, { recursive: true });

	// Filter docs belonging to this version
	const versionDocs = docs.filter((doc) => doc.slug.startsWith(`${version}/`));
	const searchIndex = generateSearchData(versionDocs);

	// Write the versioned search file
	writeFileSync(join(versionDir, "search.json"), JSON.stringify(searchIndex), { flag: "w" });

	console.log(`✅ Generated: api/${version}.search.json`);
});
```

The script:

- Imports processed documentation data from Velite
- Detects all available doc versions from the slug structure (e.g., `v1`, `v2`)
- Maps each document to a search entry with `title`, `href`, `description`, `content`, and `category`
- Uses `cleanMarkdown()` from `@penbot/core` to remove Markdown syntax for better search
- Creates a separate search JSON file per version (e.g., `api/v1.search.json/search.json`)

## API Endpoints

Each version gets its own static API route. For example, `docs/src/routes/api/v1.search.json/+server.ts`:

```ts
import type { RequestHandler } from "@sveltejs/kit";
import search from "./search.json" with { type: "json" };

export const prerender = true;

export const GET: RequestHandler = () => {
	return Response.json(search);
};
```

These routes are prerendered for optimal performance. The API is structured as:

- `/api/v1.search.json` - Search index for v1 docs
- `/api/v2.search.json` - Search index for v2 docs

## Client-Side Search

The search interface is implemented as a reusable Svelte component using `bits-ui` (Command + Dialog primitives):

```svelte
<script lang="ts">
	import MagnifyingGlass from "phosphor-svelte/lib/MagnifyingGlass";
	import { Command, Dialog } from "bits-ui";
	import { type SearchResult, createContentIndex, searchContentIndex } from "./search-utils.js";
	import { getVersionManager } from "$lib/version.svelte.js";

	const vm = getVersionManager();

	let searchState = $state<"loading" | "ready">("loading");
	let searchQuery = $state("");

	$effect(() => {
		const cb = async () => {
			const content = await fetch(`/api/${vm.current}.search.json`).then((res) => res.json());
			if (!content) return;
			createContentIndex(content);
			searchState = "ready";
		};
		cb()
	});

	const results: SearchResult[] = $derived(
		searchState === "ready" ? searchContentIndex(searchQuery) : []
	);

	let dialogOpen = $state(false);

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === "k" && (e.metaKey || e.ctrlKey)) {
			e.preventDefault();
			dialogOpen = true;
		}
	}

	let { version }: { version: string } = $props();
</script>

<svelte:document onkeydown={handleKeydown} />

<Dialog.Root bind:open={dialogOpen}>
	<Dialog.Trigger>
		<MagnifyingGlass />Search Docs ...
		<kbd>⌘</kbd>
		<kbd>K</kbd>
	</Dialog.Trigger>
	<Dialog.Portal>
		<Dialog.Overlay />
		<Dialog.Content>
			<Command.Root shouldFilter={false}>
				<Command.Input bind:value={searchQuery} placeholder="Search for something..." />
				{#if searchQuery !== "" && results.length === 0}
					<Command.Empty>No results found.</Command.Empty>
				{/if}
				{#if searchQuery !== "" && results.length > 0}
					<Command.List>
						{#each results as { title, href, snippet, category }}
							<Command.LinkItem {href} onSelect={() => { searchQuery = ""; dialogOpen = false; }}>
								<span>{title}</span>
								{#if category}
									<span>{category}</span>
								{/if}
								{#if snippet}
								{/if}
							</Command.LinkItem>
						{/each}
					</Command.List>
				{/if}
			</Command.Root>
		</Dialog.Content>
	</Dialog.Portal>
</Dialog.Root>
```

The component:

- **Fetches versioned search data** on mount using `$effect` - requests the correct index based on the current docs version via `getVersionManager()`
- **Creates FlexSearch indexes** for fast querying of titles and content
- **Uses `bits-ui` Dialog** with `Command.Root` for a search-as-you-type experience
- **Supports keyboard shortcut** (Cmd/Ctrl + K) to open the search dialog
- **Displays results with snippets** showing the matching context with highlighted terms

## Version-Aware Search

Because Penbot supports multiple documentation versions, the search component is version-aware:

- The API endpoint includes the version: `/api/v1.search.json` or `/api/v2.search.json`
- The `getVersionManager()` provides the current version context from the URL
- When a user switches versions, the search index is re-fetched for the new version
- Each version's search data only contains docs relevant to that version

## Search Result Structure

Each search result contains:

```ts
type SearchContent = {
	title: string;
	content: string;
	description: string;
	href: string;
	category?: string;
};

type SearchResult = SearchContent & {
	snippet?: string; // Content snippet with highlighted matches
	highlights?: string[]; // Array of highlighted terms
	category?: string; // Document category/section
};
```

## Search Algorithm

The search uses a multi-tiered approach in `searchContentIndex()`:

1. **Primary search**: FlexSearch indexes for title and content with different weights
   - Title matches: **10 points**
   - Content matches: **5 points**
2. **Fallback search**: Fuzzy matching when no exact results found
   - Fuzzy title matches: **8 points**
   - Fuzzy content matches: **3 points**

Results are ranked by score, limited to 10 items, and each result includes a contextual snippet with highlighted matching terms.

## Usage

### Basic Integration

The search component is already included in the default layout's header. To use it in your own layout:

```svelte
<script>
	import Search from "@penbot/core/components/search/search.svelte";
</script>

<Search version="v1" />
```

### Search Utilities

The search utilities (`createContentIndex`, `searchContentIndex`, `cleanMarkdown`) are exported from `@penbot/core`:

```ts
import { createContentIndex, searchContentIndex, cleanMarkdown, defineSearchContent } from "@penbot/core";
```

These can be used to build custom search interfaces or pre-process content for indexing.
