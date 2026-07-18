# Getting Started
> A quick guide to get started using Penbot.



The following guide will walk you through the process of getting a Penbot project up and running.

## Clone the repo

Clone the Penbot repo from GitHub:

```bash
git clone https://github.com/beyondgrayzone/penbot.git my-docs
cd my-docs
```


## 

> Replace `my-docs` with your project name.

## Install dependencies and build

```bash
bun i
bun run build
```

The first build generates the required Velite artifacts. You only need to run `bun run build` once before starting the dev server.

## Start the dev server

```bash
bun run dev
```

Open [http://localhost:5173](http://localhost:5173)  you'll see the default Penbot documentation site.

## Navigation

The starter template comes with a navigation structure that is split across two files:

### 1. Static Config (`docs/src/lib/navigation.json`)

Defines the structure of anchors and sections:

```json
{
	"anchors": [
		{ "title": "Introduction", "href": "/", "icon": "ChalkboardTeacher" },
		{ "title": "Getting Started", "href": "/getting-started", "icon": "RocketLaunch" }
	],
	"sections": [
		{ "title": "Configuration" },
		{ "title": "Components" }
	]
}
```

### 2. Dynamic Builder (`docs/src/lib/navigation.ts`)

Imports the config and dynamically populates section items from Velite-processed docs:

```ts
import { defineNavigation } from "@penbot/core";
import navConfig from "./navigation.json";
import { getAllDocs } from "./utils.js";

export const navigation = defineNavigation({
	anchors: dynamicAnchors("v1"),
	sections: dynamicSections("v1"),
});
```

### Structure

- **Anchors**  Links at the top of the sidebar (for important/highlight links, can include icons)
- **Header**  Links in the top navigation bar
- **Sections**  Group related pages under categories (e.g., "Components", "Configuration")
- **Items**  Flat list of sidebar items without a section header

## Site config

The site config is used to configure site-wide settings, such as the title, description, keywords, ogImage, and other metadata. The config is located in the `docs/src/lib/site-config.json` file.

```json
{
	"name": "Your Product",
	"siteLink": "/",
	"url": "https://yourproduct.com",
	"ogImage": {
		"url": "https://yourproduct.com/og.png",
		"height": "630",
		"width": "1200"
	},
	"description": "Your product description.",
	"keywords": ["your", "keywords", "here"],
	"license": {
		"name": "MIT",
		"urll": "https://opensource.org/licenses/MIT"
	},
	"links": {}
}
```

### Per-Route Site Config

You can override any part of the site config on a per-route basis using the `useSiteConfig` hook.


## Under Development

> This feature is still being worked on.

## Theme

Penbot uses the `mode-watcher` library to manage themes. To change the default theme, edit `docs/src/routes/+layout.svelte`:

```svelte
<script>
	import { ModeWatcher } from "mode-watcher";
</script>

<ModeWatcher defaultTheme="oceanic" />   <!-- change this -->
```

Also update the static fallback in `docs/src/app.html`:

```html
<html lang="en" data-theme="oceanic">   <!-- match the defaultTheme -->
```

### CSS imports

In `docs/src/app.css`, import the `globals.css` **first**, then import any theme files **after**:

```css
@import "@penbot/core/globals.css";
@import "@penbot/core/theme-oceanic.css";
@import "@penbot/core/theme-forest.css";
@import "@penbot/core/theme-sober-1.css";
@import "@penbot/core/theme-amber.css";
```

### Available themes

| Theme    | Import path                          |
|----------|--------------------------------------|
| oceanic  | `@penbot/core/theme-oceanic.css`   |
| forest   | `@penbot/core/theme-forest.css`    |
| sober-1  | `@penbot/core/theme-sober-1.css`   |
| orange   | `@penbot/core/theme-orange.css`    |
| green    | `@penbot/core/theme-green.css`     |
| blue     | `@penbot/core/theme-blue.css`      |
| purple   | `@penbot/core/theme-purple.css`    |
| pink     | `@penbot/core/theme-pink.css`      |
| lime     | `@penbot/core/theme-lime.css`      |
| yellow   | `@penbot/core/theme-yellow.css`    |
| cyan     | `@penbot/core/theme-cyan.css`      |
| teal     | `@penbot/core/theme-teal.css`      |
| violet   | `@penbot/core/theme-violet.css`    |
| amber    | `@penbot/core/theme-amber.css`     |
| red      | `@penbot/core/theme-red.css`       |
| sky      | `@penbot/core/theme-sky.css`       |
| emerald  | `@penbot/core/theme-emerald.css`   |
| fuchsia  | `@penbot/core/theme-fuchsia.css`   |
| indigo   | `@penbot/core/theme-indigo.css`    |
| rose     | `@penbot/core/theme-rose.css`      |

## Logo

To customize the logo displayed in the sidebar header, place `logo-small.png` (light mode) and `logo-small-dark.png` (dark mode) in `docs/static/`. These are auto-detected by the core layout.

For a custom Svelte snippet, edit the logo in `docs/src/routes/(docs)/+layout.svelte`:

```svelte
{#snippet logo()}
	<LogoDark class="hidden h-7 dark:block" />
	<LogoLight class="block h-7 dark:hidden" />
	<span class="sr-only">Your project name</span>
{/snippet}
```
