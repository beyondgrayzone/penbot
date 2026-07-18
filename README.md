# What is Penbot?

[![GitHub](https://img.shields.io/badge/GitHub-beyondgrayzone/penbot-181717?logo=github)](https://github.com/penbot/penbot)

Penbot is a starter kit for building **static documentation sites** using SvelteKit. It's not a framework  just clone it, strip out the demo content, and make it yours.

- Uses [SvelteKit](https://kit.svelte.dev/) to generate fully static HTML sites
- Markdown-powered documentation with [Velite](https://velite.js.org/)
- Built-in version management (v1, v2, v3...)
- Searchable, section-organized docs out of the box
- Ready-made Svelte components: Callouts, Tabs, Steps, Cards, and more

> Not for everyone, but it may very well be for you.

---

# Getting Started

> **Repo**: [https://github.com/penbot/penbot](https://github.com/penbot/penbot)

```bash
# 1. Clone the repo
git clone https://github.com/penbot/penbot.git my-docs
cd my-docs

# 2. Install dependencies (requires Bun >= v1.3.9)
bun i

# 3. Build required artifacts (first time only)
bun run build

# 4. Start the dev server
bun run dev
```

Open `http://localhost:5173` and you'll see the default Penbot documentation site.

---

## Turning it into your own docs

Follow these steps to replace the demo content with documentation for your own product:

### 1. Remove the existing versions

```bash
bun run docs:version:remove -- v2
bun run docs:version:remove -- v1
```

### 2. (Optional) Clean up stale build artifacts

```bash
rm -rf docs/src/mds_final
```

### 3. Add your fresh v1

```bash
bun run docs:version:add -- v1
```

This creates blank scaffolding:
- Route: `docs/src/routes/(docs)/docs/v1/`
- API search: `docs/src/routes/api/v1.search.json/`
- Markdown dir: `docs/src/mds/v1/` (with a starter `index.md`)

### 4. Configure your sections

```bash
# See the current defaults
bun run docs:sections list

# Remove unneeded ones
bun run docs:sections remove -- Components

# Add your own
bun run docs:sections add -- "CLI Reference"
bun run docs:sections add -- Features
```

### 5. Update the site config

Edit `docs/src/lib/site-config.json`  change the **name**, **description**, **URL**, **keywords**, and **license** to match your product.

### 6. Update branding assets

Replace the images in `docs/static/`:
- `favicon.png`  your product's favicon
- `logo-small.png` & `logo-small-dark.png`  your product's logo

### 7. Write your markdown

Create `.md` files under `docs/src/mds/v1/`. At minimum you need:
- `index.md`  landing page (must have `currentVersion: 1.0.0` in frontmatter)
- Any other pages (e.g., `getting-started.md`, `configuration.md`)

Each file needs frontmatter:

```markdown
---
title: Getting Started
description: A quick guide to get started.
section: Overview
---

Content here using Markdown and Svelte components.
```

### 8. Update the search index

Edit `docs/src/routes/api/v1.search.json/search.json` with search data for all your pages.

### 9. Sync everything

```bash
bun run docs:version:sync
```

### 10. Preview and build

```bash
# Run the dev server
bun run dev

# Build the static site
bun run build
```

The static output goes to `docs/build/` and can be deployed anywhere (GitHub Pages, Netlify, etc.).

---

# Prereq

- **Bun** >= v1.3.9
- **Go** >= 1.21 (required for the `@penbot/mds` tool  version management, sections, theme CLI, and markdown sanitization)

# Commands (from root workspace)

| Command | Description |
|---------|-------------|
| `bun i` | Install all dependencies |
| `bun run build` | Build artifacts (first time) or static site |
| `bun run dev` | Start the dev server |
| `bun run docs:version:add -- v<N>` | Add a new version |
| `bun run docs:version:remove -- v<N>` | Remove a version |
| `bun run docs:version:sync` | Sync all versions |
| `bun run docs:sections list` | List all sections |
| `bun run docs:sections add -- <name>` | Add a section |
| `bun run docs:sections remove -- <name>` | Remove a section |
| `bun run docs:theme:list` | List all available themes |
| `bun run docs:theme:apply -- <name>` | Set the default theme and update the theme timestamp |
| `bun run lint` | Run Biome linter |
| `bun run format` | Run Biome formatter |

## Version Management

Versions are managed via `@penbot/mds` Go tool. Available commands:

```bash
# Add a new version from scratch (e.g. v3)
bun run docs:version:add -- v3

# Add a new version by copying content from the previous version (e.g. v3 from v2)
bun run docs:version:add -- --copy v3

# Remove a version (v1 is protected from removal)
bun run docs:version:remove -- v3

# Sync all versions  ensures every versioned route has its
# corresponding mds folder with index.md and search.json endpoint
bun run docs:version:sync

# List all sections defined in velite config
bun run docs:sections
```

### `add` (from scratch)
Creates empty scaffolding for a new version:
- Route scaffolding: `docs/src/routes/(docs)/docs/{version}/` with `+page.svelte`, `+page.ts`, and `[...slug]/` route
- API search endpoint: `docs/src/routes/api/{version}.search.json/` with `+server.ts` and `search.json`
- Markdown directory: `docs/src/mds/{version}/` with a starter `index.md` (`currentVersion: X.0.0`)

### `add --copy` (from previous version)
Copies the entire markdown content from the previous version (e.g., v2 → v3), then updates only the `currentVersion` in the copied `index.md`. All other files, frontmatter, and body content are preserved. This is useful when creating a new patch version with mostly the same documentation.

### `remove`
Deletes all files created by `add` (routes, API search, mds folder). v1 is protected from removal as it is the minimum required version.

### `sync`
Scans all versioned route directories and creates any missing mds folders or search.json endpoints. It also warns about orphan mds folders or search.json directories that exist without a corresponding route. The core `search.json` in `packages/core/` is rebuilt from the actual `currentVersion` values in each version's `index.md`.

### `sections`
Lists all documentation sections defined in `docs/velite.config.js`. This reads the `sections` array from the config file and prints each value on a separate line.

### `sections add <name>`
Appends a new section to the `sections` array in `docs/velite.config.js`.

### `sections remove <name>`
Finds and removes a section from the `sections` array in `docs/velite.config.js`.

## Theme CLI

The `@penbot/mds` Go tool also provides theme management. The default theme and theme timestamp are stored in `docs/src/lib/site-config.json`. Available commands:

```bash
# List all available themes (current default marked with *)
bun run docs:theme:list

# Set the default theme and auto-update the timestamp to now
bun run docs:theme:apply -- oceanic
```

### `list`
Scans `packages/core/src/lib/styles/theme-*.css` and prints every theme name. The currently configured default is marked with `*`.

### `apply <name>`
Sets the `defaultTheme` and updates `themeTimestamp` to `Date.now()` in `docs/src/lib/site-config.json`. Validates that the theme exists before writing.

On the next deploy, returning visitors will have their theme overridden to the new default if the server timestamp is greater than their saved client timestamp. See [.skills/theme/SKILL.md](.skills/theme/SKILL.md) for details.

### Prerequisite
The `mds` tool requires **Go** >= 1.21.

# Monorepo
 - Uses bun to organize private packages
 - See `workspaces` && `script` properties `package.json` file from root workspace

## Monorepo Structure

### docs folder
  - docs contains the main markdown files inside `src/mds` directory
  - docs also contains the sanitized version of the above inside `src/mds_final` directory (Used for clipboard)

### packages/core contains the scaffolding code for the penbot framework. This folder should not be touched

### packages/mds is a Go program that provides:
  - **Markdown sanitization**: strips HTML/Svelte component tags from markdown to render clean text (used for clipboard). See `runProcessMds` function in `process.ts` file.
  - **Version management**: add, remove, and sync documentation versions (see [Version Management](#version-management))
  - **Section listing**: list documentation sections from the velite config (see [sections](#sections))
  - **Theme management**: list available themes and apply a new default theme with timestamp (see [Theme CLI](#theme-cli))

### packages/mdsx is the markdown processor sveltekit plugin that converts svelte snippets to html code for penbot. This folder should not be touched

# How it works?

- `docs` folder contains these items for the actual documentation
    * markdown files
    * versioned routes
    * landing page route (currently not implemeted)
    * server side logic for markdown preprocessing
  - It uses velite for markdown processing, whose config can be found in `docs/.velite.config.js` file
    - This file defines all sections in the above file 
    - This file defines frontmatter properties in the `baseSchema` in the above file as well
      Ex: 
        ```
        ---
        title: Getting Started
        description: A quick guide to get started using Penbot
        section: Overview
        ---
        ```
- Markdown files
  - `docs/src/mds` folder is where you put all documentation related markdown files
  - Docs are versioned by default Ex: `docs/src/mds/v1/getting-started.md`
  - You need atleast one version called `v1`
  - Each doc must have `index.md` file, Ex: `docs/src/mds/v2/index.md`
  - Use `currentVersion` propery to define the version Ex: `currentVersion: 2.0.1` means this is available since version `2.0.1`
  - `index.md` files, inside `src/mds` folder, must have `currentVersion` set to some value

# Config 
  - All config can be found at `docs/src/lib/site-config.json`
  - When creating doc for a particular product, use this config
  
# How to create brand new doc using this codebase?

Follow the [Turning it into your own docs](#turning-it-into-your-own-docs) section above. For a condensed reference, see [NEW-DOC.md](./NEW-DOC.md).

# Linting & formatting
- Biome for linting and formatting. Its WIP, so PRs are welcome

# Current downsides
- Its very simple, so for complex usecase this may not be for you
- Since this is not a framework, it may be weird for some to make direct edits on a cloned repo
- Configurations are currently manual and some care needs to be taken
- There is no way to navigate a specific version, example: if the currentVersion is `2.1.3` but previously it was `2.0.9`, you can only see the latest version since the url would be `/v2/`, which would always point to the latest version. You can however, individually see specific versioned doc by manual navigation of pages

# Skills / Tags

All documentation items from `docs/src/mds/v1` have been extracted into `.skills/` folder with quick-reference `SKILL.md` files.

## Components

![Button](https://img.shields.io/badge/Button-blue)
![Callout](https://img.shields.io/badge/Callout-blue)
![Card](https://img.shields.io/badge/Card-blue)
![CardContainer](https://img.shields.io/badge/CardContainer-blue)
![CardGrid](https://img.shields.io/badge/CardGrid-blue)
![Checkbox](https://img.shields.io/badge/Checkbox-blue)
![Collapsible](https://img.shields.io/badge/Collapsible-blue)
![Input](https://img.shields.io/badge/Input-blue)
![NativeSelect](https://img.shields.io/badge/NativeSelect-blue)
![PropField](https://img.shields.io/badge/PropField-blue)
![Select](https://img.shields.io/badge/Select-blue)
![Steps](https://img.shields.io/badge/Steps-blue)
![Switch](https://img.shields.io/badge/Switch-blue)
![Tabs](https://img.shields.io/badge/Tabs-blue)
![Textarea](https://img.shields.io/badge/Textarea-blue)

## Configuration

![Navigation](https://img.shields.io/badge/Navigation-green)
![Search](https://img.shields.io/badge/Search-green)
![Theme](https://img.shields.io/badge/Theme-green)

## Overview

![Getting Started](https://img.shields.io/badge/Getting_Started-purple)
![Introduction](https://img.shields.io/badge/Introduction-purple)

## Assets

![Images & Assets](https://img.shields.io/badge/Images_%26_Assets-yellow)

## Plain Text Tags

```
button, callout, card, card-container, card-grid, checkbox, collapsible, input, native-select, prop-field, select, steps, switch, tabs, textarea, navigation, search, theme, getting-started, introduction, images-and-assets
```
