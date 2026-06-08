# How to Create a Brand New Documentation Site from Bladocs

This guide documents the minimal, working steps to take the Bladocs starter and turn it into documentation for your own product (e.g., NoteBox, a CLI tool, an API, etc.).

---

## Step 1: Remove all existing versions

```bash
# Remove existing documentation versions (v1, v2, etc.)
bun run docs:version:remove -- v2
bun run docs:version:remove -- v1
```

This deletes route scaffolding, API search endpoints, and markdown files for each version.

---

## Step 2: (Optional) Clean up stale build artifacts

```bash
rm -rf docs/src/mds_final
```

---

## Step 3: Add your fresh v1

```bash
bun run docs:version:add -- v1
```

This creates:
- Route scaffolding: `docs/src/routes/(docs)/docs/v1/` (with `+page.svelte`, `+page.ts`, `[...slug]/`)
- API search endpoint: `docs/src/routes/api/v1.search.json/` (with `+server.ts` and empty `search.json`)
- Markdown directory: `docs/src/mds/v1/` (with starter `index.md`)

---

## Step 4: Update the `sections` array

First, evaluate which sections your documentation needs and which defaults are unneeded.

```bash
# See the current default sections
bun run docs:sections list
```

Remove default sections you don't need, then add your own:

```bash
# Remove unneeded defaults (example)
bun run docs:sections remove -- Components
bun run docs:sections remove -- Utilities

# Add your own sections (example)
bun run docs:sections add -- "CLI Reference"
bun run docs:sections add -- Features

# Verify the final list
bun run docs:sections list
```

> `Overview` and `Configuration` are commonly useful defaults. Adjust the list to match your documentation structure.

---

## Step 5: Update the site config

Edit `docs/src/lib/site-config.json` — change the name, description, URL, keywords, and license to match your product.

---

## Step 6: Write your markdown files

Create markdown files under `docs/src/mds/v1/`. At minimum you need:
- `index.md` — The introduction/landing page for your docs
- Any other pages you need (e.g., `getting-started.md`, `configuration.md`, etc.)

Each file must have frontmatter with **at least** `title`, `description`, and `section`. The `section` value must match one of the sections defined in `velite.config.js`.

```markdown
---
title: Getting Started
description: A quick guide to get started.
section: Overview
---

Content here using Markdown and Svelte components (Callout, Steps, Tabs, Card, etc.).
```

The `index.md` file should also include `currentVersion: 1.0.0`.

---

## Step 7: Update the search index

Edit `docs/src/routes/api/v1.search.json/search.json` with search data for all your pages. Each entry needs:

```json
{
  "title": "Page Title",
  "href": "/docs/v1/page-slug",
  "description": "Short description",
  "content": "Searchable plain text content...",
  "category": "SectionName"
}
```

---

## Step 8: Sync everything

```bash
bun run docs:version:sync
```

This ensures all routes, mds folders, and search endpoints are consistent and updates the core search.json.

---

## One-shot Prompt (Copy & Paste)

> I want to create documentation for **{ProductName}**, a {one-line description}.
>
> Using the Bladocs codebase:
>
> 1. Remove all existing documentation versions (v1, v2, etc.) using `bun run docs:version:remove -- v2` then `-- v1`.
> 2. Add a fresh v1 using `bun run docs:version:add -- v1`.
> 3. First evaluate all required sections and identify unneeded defaults by running `bun run docs:sections list`. Then remove unneeded sections with `bun run docs:sections remove -- <name>` and add required ones with `bun run docs:sections add -- <name>` to match my documentation structure.
> 4. Update `docs/src/lib/site-config.json` with my product's name, URL, description, and keywords.
> 5. Write markdown files for all pages under `docs/src/mds/v1/`, covering:
>    - Index/introduction page
>    - Getting started guide
>    - Configuration reference
>    - CLI reference (if applicable)
>    - Feature-specific pages
> 6. Update `docs/src/routes/api/v1.search.json/search.json` with search data for all pages.
> 7. Run `bun run docs:version:sync` to sync everything.
