---
title: Getting Started
description: Quick start guide for setting up and configuring a Penbot documentation site.
---

# Getting Started

## Clone the starter

```bash
pnpx degit penbot_eco/penbot/start
```

## Navigation setup

Edit `src/lib/navigation.ts`:

```ts
import { defineNavigation } from "@penbot/core";
export const navigation = defineNavigation({ /* customize */ });
```

## Site config

Edit `src/lib/site-config.ts`:

```ts
import { defineSiteConfig } from "@penbot/core";
export const siteConfig = defineSiteConfig({
  title: "Bladocs",
  description: "A SvelteKit docs starter template",
  keywords: "sveltekit, docs, starter, template",
  ogImage: { url: "https://...", height: 630, width: 1200 },
});
```

Override per-route with `useSiteConfig` hook.

## Theme

In `src/app.css`:

```css
@import "@penbot/core/themes/emerald.css";
@import "@penbot/core/globals.css";
```

## Logo

Edit sidebar logo in `src/routes/(docs)/+layout.svelte` via the `logo` snippet:

```svelte
{#snippet logo()}
  <LogoDark class="hidden h-7 dark:block" />
  <LogoLight class="block h-7 dark:hidden" />
  <span class="sr-only">Project name</span>
{/snippet}
```

## Markdown frontmatter

Each markdown document supports the following frontmatter fields defined in `docs/velite.config.js`:

| Field | Type | Description |
|---|---|---|
| `title` | `string` | Page title (rendered as heading) |
| `description` | `string` | Page description (rendered below title) |
| `section` | `string` | Section grouping (must match one of `docs/velite.config.js` sections) |
| `currentVersion` | `string` | Version string for the overall doc version (e.g. `1.0.0`) |
| `availableSinceVersion` | `string` | (optional) When this feature/component was introduced — renders a badge like `Available since: 1.0.0` below the title |
| `navLabel` | `string` | (optional) Custom label for sidebar navigation |
| `order` | `number` | (optional) Sort order within a section |

### Using `availableSinceVersion`

Add it to the frontmatter of any markdown file to show when a feature became available:

```md
---
title: Checkbox
description: A toggle input component.
section: Components
availableSinceVersion: 1.0.3
---
```

This renders as an "Available since" badge next to the page title. It is useful for documenting features added in later versions of your project.
