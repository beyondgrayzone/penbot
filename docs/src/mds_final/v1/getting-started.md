# Getting Started
> A quick guide to get started using Bladocs



The following guide will walk you through the process of getting a Bladocs project up and running.

## Clone the starter template

Clone the Bladocs starter template:

```bash
pnpx degit bladocs_eco/bladocs/start
```

## Navigation

The starter template comes with a basic navigation structure to get your started. To customize the navigation, adjust the `src/lib/navigation.ts` file.

```ts
import { defineNavigation } from "@bladocs/core";

export const navigation = defineNavigation({
	// Customize the navigation here
});
```

## Site config

The site config is used to configure site-wide settings, such as the title, description, keywords, ogImage, and other metadata.

The config is located in the `src/lib/site-config.ts` file.

```ts
import { defineSiteConfig } from "@bladocs/core";

export const siteConfig = defineSiteConfig({
	title: "Bladocs",
	description: "A SvelteKit docs starter template",
	keywords: "sveltekit, docs, starter, template",
	ogImage: {
		url: "https://docs.bladocs.dev/bladocs.png",
		height: 630,
		width: 1200,
	},
});
```

### Per-Route Site Config

You can override any part of the site config on a per-route basis using the `useSiteConfig` hook.


## Under Development

> This feature is still being worked on.

## Theme

The starter template comes with the default Bladocs theme (orange). To customize the theme, adjust the import in the `src/app.css` file to reflect the color scheme you want to use for your project. Each theme has been designed to work well in both light and dark mode.

```css {1-2}
/* @import "@bladocs/core/themes/orange.css"; */
@import "@bladocs/core/themes/emerald.css";
@import "@bladocs/core/globals.css";
```

## Logo

To customize the logo displayed in the sidebar header, head to the `src/routes/(docs)/+layout.svelte` file and adjust the contents of the `logo` snippet. If the logo has a light and dark version, ensure to handle those similarly to the default Bladocssystem logo.

```go title="main.go"
func main() {
    fmt.Println()
}
```


```svelte title="src/routes/(docs)/+layout.svelte"
{#snippet logo()}
	<LogoDark class="hidden h-7 dark:block" />
	<LogoLight class="block h-7 dark:hidden" />
	<span class="sr-only">The project name here</span>
{/snippet}
```
