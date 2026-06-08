# What is Bladocs?
- This repo is Bladocs
- This repo exhibits an example of using Sveltekit as means to create static documentation sites
- This repo is not a framework, just clone it and start changing stuff to make it yours
- Once you make it yours, make approriate changes to this README.md file
- Its not for everyone, but it may very well be for you

# Prereq
- Bun >= v1.3.9

# Required manual prep
- adjust images in docs/static folder 
  - favicon and product images
  - adjust `logo-small.png` & `logo-small-dark.png` files to suit the product

# Commands 
- Install deps - `bun i` from root workspace
- Build required artifacts for dev server (only first time) - `bun run build` from root workspace
- Run dev server - `bun run dev` from root workspace
- Build static website - `bun run build` from root workspace

## Version Management

Versions are managed via `@bladocs/mds` Go tool. Available commands:

```bash
# Add a new version from scratch (e.g. v3)
bun run docs:version:add -- v3

# Add a new version by copying content from the previous version (e.g. v3 from v2)
bun run docs:version:add -- --copy v3

# Remove a version (v1 is protected from removal)
bun run docs:version:remove -- v3

# Sync all versions — ensures every versioned route has its
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

# Monorepo
 - Uses bun to organize private packages
 - See `workspaces` && `script` properties `package.json` file from root workspace

## Monorepo Structure

### docs folder
  - docs contains the main markdown files inside `src/mds` directory
  - docs also contains the sanitized version of the above inside `src/mds_final` directory (Used for clipboard)

### packages/core contains the scaffolding code for the bladocs framework. This folder should not be touched

### packages/mds is a Go program that provides:
  - **Markdown sanitization**: strips HTML/Svelte component tags from markdown to render clean text (used for clipboard). See `runProcessMds` function in `process.ts` file.
  - **Version management**: add, remove, and sync documentation versions (see [Version Management](#version-management))
  - **Section listing**: list documentation sections from the velite config (see [sections](#sections))

### packages/mdsx is the markdown processor sveltekit plugin that converts svelte snippets to html code for bladocs. This folder should not be touched

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
        description: A quick guide to get started using Bladocs
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
    - See NEW-DOC.md

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

# TODO
- Make Bladocs more easy for everyone to just get started with
