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
