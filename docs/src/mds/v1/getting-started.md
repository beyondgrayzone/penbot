---
title: Getting Started
description: A quick guide to get started using Penbot.
section: Overview
---

<script>
	import { Callout } from "@penbot/core";
</script>

The following guide will walk you through the process of getting a Penbot project up and running.

## Clone the repo

Clone the Penbot repo from GitHub:

```bash
git clone https://github.com/beyondgrayzone/penbot.git my-docs
cd my-docs
```

<Callout type="note">
Replace `my-docs` with your project name.
</Callout>

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

The starter template comes with a basic navigation structure to get you started. To customize the navigation, adjust the `docs/src/lib/navigation.ts` file.

```ts
import { defineNavigation } from "@penbot/core";

export const navigation = defineNavigation({
	// Customize the navigation here
});
```

### Structure

- **Anchors**  Links at the top of the sidebar (for important/highlight links)
- **Sections**  Group related pages under categories (e.g., "Components", "Configuration")

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

<Callout type="warning" title="Under Development">
This feature is still being worked on.
</Callout>

## Theme

To customize the theme, edit the `docs/src/app.css` file. Import a theme CSS file **before** the globals:

```css
/* @import "@penbot/core/themes/orange.css"; */
@import "@penbot/core/themes/emerald.css";
@import "@penbot/core/globals.css";
```

### Available themes

| Theme    | Import path                          |
|----------|--------------------------------------|
| orange   | `@penbot/core/themes/orange.css`   |
| green    | `@penbot/core/themes/green.css`    |
| blue     | `@penbot/core/themes/blue.css`     |
| purple   | `@penbot/core/themes/purple.css`   |
| pink     | `@penbot/core/themes/pink.css`     |
| lime     | `@penbot/core/themes/lime.css`     |
| yellow   | `@penbot/core/themes/yellow.css`   |
| cyan     | `@penbot/core/themes/cyan.css`     |
| teal     | `@penbot/core/themes/teal.css`     |
| violet   | `@penbot/core/themes/violet.css`   |
| amber    | `@penbot/core/themes/amber.css`    |
| red      | `@penbot/core/themes/red.css`      |
| sky      | `@penbot/core/themes/sky.css`      |
| emerald  | `@penbot/core/themes/emerald.css`  |
| fuchsia  | `@penbot/core/themes/fuchsia.css`  |
| rose     | `@penbot/core/themes/rose.css`     |

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
