---
title: Site Config
description: Learn how to configure site-wide settings in your Penbot project.
section: Configuration
---

The site config centralizes all site-wide metadata, branding, and social links. It's the single source of truth for your project's identity across the landing page, sidebar, footer, and SEO metadata.

## Configuration File

Site-wide settings are defined in `docs/src/lib/site-config.json`:

```json
{
  "name": "Penbot",
  "siteLink": "/",
  "url": "https://penbot.codelinter.com",
  "description": "Documentation toolkit for Penbot.",
  "keywords": ["penbot", "sveltekit", "documentation", "docs"],
  "ogImage": {
    "url": "https://docs.penbot.dev/penbot.png",
    "height": "630",
    "width": "1200"
  },
  "license": {
    "name": "MIT",
    "url": "https://yourwebsite.com/LICENSE"
  },
  "links": {
    "github": "https://github.com/beyondgrayzone/penbot",
    "x": "https://x.com/penbot"
  },
  "defaultTheme": "oceanic",
  "themeTimestamp": 1721300000000
}
```

## Full Type Definition

The full `SiteConfig` type is defined in `@penbot/core`:

```ts
type SiteConfig = {
  name: string;
  url: string;
  siteLink: string;
  description: string;
  links?: {
    logo?: string;
    x?: string;
    github?: string;
  };
  author?: string;
  keywords?: string[];
  ogImage?: {
    url: string;
    width: string;
    height: string;
  };
  license?: {
    name: string;
    url: string;
  };
  themeTimestamp?: number;
  defaultTheme?: string;
};
```

## Field Reference

### Branding

| Field | Required | Description |
|-------|----------|-------------|
| `name` | Yes | Site name. Used in the sidebar, page titles (`Title - Name`), footer copyright, and meta tags |
| `siteLink` | Yes | Root URL for the site logo/link in the sidebar |
| `url` | Yes | Canonical URL of the site. Used in Open Graph and Twitter card meta tags |
| `description` | Yes | Default meta description. Used as fallback when a page doesn't provide its own |

### SEO & Metadata

| Field | Required | Description |
|-------|----------|-------------|
| `keywords` | No | Array of keywords for the `<meta name="keywords">` tag |
| `author` | No | Author name for the `<meta name="author">` tag |
| `ogImage` | No | Open Graph image configuration with `url`, `width`, and `height` |

### Social Links

| Field | Required | Description |
|-------|----------|-------------|
| `links.github` | No | GitHub URL. Renders a GitHub icon in the header and footer |
| `links.x` | No | X (Twitter) URL. Renders an X icon in the header and footer |

### Legal

| Field | Required | Description |
|-------|----------|-------------|
| `license.name` | No | License name displayed in the footer |
| `license.url` | No | Link to the full license text |

### Theme

| Field | Required | Description |
|-------|----------|-------------|
| `defaultTheme` | No | Default theme for first-time visitors. Must match a theme CSS file imported in `app.css` (e.g. `"oceanic"`, `"forest"`) |
| `themeTimestamp` | No | Unix timestamp (ms) for forcing theme resets on returning visitors after a redeploy. See [Theme](./theme.md) for details |

## How to Use

### 1. Define the Config

Import and type-check your config using `defineSiteConfig` from `@penbot/core`:

```ts
// docs/src/lib/site-config.ts
import { defineSiteConfig } from "@penbot/core";
import siteConfigData from "./site-config.json";

export const siteConfig = defineSiteConfig(siteConfigData);
```

### 2. Initialize in Root Layout

Call `useSiteConfig` in your root layout to make the config available to all child components:

```svelte
<script lang="ts">
  import { useSiteConfig } from "@penbot/core";
  import { siteConfig } from "$lib/site-config";

  useSiteConfig(() => siteConfig);
</script>
```

### 3. Access Anywhere

Once initialized, any component can access the config via `useSiteConfig()`:

```svelte
<script lang="ts">
  import { useSiteConfig } from "@penbot/core";

  const siteConfig = useSiteConfig();
  let name = $derived(siteConfig.current.name);
  let url = $derived(siteConfig.current.url);
  let github = $derived(siteConfig.current.links?.github);
</script>

<footer>
  <span>&copy; {new Date().getFullYear()} {name}</span>
  {#if github}
    <a href={github}>GitHub</a>
  {/if}
</footer>
```

The `siteConfig.current` property is reactive and automatically reflects the latest config values.

## Where the Config Is Used

The site config powers the following built-in components:

| Component | What it uses |
|-----------|-------------|
| **Metadata** (`<svelte:head>`) | `name`, `description`, `keywords`, `ogImage`, `url` |
| **Footer** | `name`, `url`, `links.github`, `links.x` |
| **Social Icons** (header/footer) | `links.github`, `links.x` |
| **Sidebar** | `name` (default logo text), `siteLink` |
| **Theme System** | `defaultTheme`, `themeTimestamp` |
| **Version Switcher** | `name` |

## Per-Route Overrides

You can override site config values on a per-route basis. For example, the `Metadata` component accepts props that take priority over the site config defaults:

```svelte
<Metadata
  title="Custom Page Title"
  description="A custom description for this page only"
  ogImage={{ url: "/custom-og.png", width: "1200", height: "630" }}
/>
```

This pattern is used in `DocPage` to let individual documentation pages provide their own title, description, and OG image from frontmatter, falling back to the site-wide defaults from the config.
