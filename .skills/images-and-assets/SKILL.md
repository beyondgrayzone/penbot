---
title: Images & External Assets
description: Guide for managing images, favicons, logos, fonts, and static assets.
---

# Images & External Assets

## Static Folder

All files in `docs/static/` are served at the root `/` URL path. Common assets placed here:

```
docs/static/
├── favicon.ico              # Browser tab icon
├── favicon-16x16.png        # Favicon 16px
├── favicon-32x32.png        # Favicon 32px
├── apple-touch-icon.png     # iOS home screen icon
├── android-chrome-192x192.png
├── android-chrome-512x512.png
├── site.webmanifest         # PWA manifest
├── penbot.png              # OG image (social sharing)
├── logo-small.png           # Sidebar logo (light mode)
├── logo-small-dark.png      # Sidebar logo (dark mode)
├── logo-light.svg           # Full logo light variant
└── logo-dark.svg            # Full logo dark variant
```

## Adding Images in Markdown

Use standard Markdown image syntax — images in `docs/static/` are accessible at the root `/`:

```md
![Alt text](/penbot.png)
![Alt text](/my-image.png)
```

### Sizing

Wrap in HTML for sizing:

```html
<img src="/my-image.png" alt="Alt text" width="600" />
```

## OG Image (Social Sharing)

Configured in `docs/src/lib/site-config.json`:

```json
{
  "ogImage": {
    "url": "https://yoursite.com/og-image.png",
    "height": "630",
    "width": "1200"
  }
}
```

Place the OG image in `docs/static/` and reference its URL.

## Favicon

Configured in `docs/src/app.html`:

```html
<link rel="icon" href="%sveltekit.assets%/favicon.ico" />
```

Replace favicon files in `docs/static/` with your own.

## Sidebar Logo

### Simple variant (image only)

Place `logo-small.png` (light mode) and `logo-small-dark.png` (dark mode) in `docs/static/`. These are auto-detected by the core layout.

### Custom variant (Svelte snippet)

In `src/routes/(docs)/docs/+layout.svelte` or your own layout, pass a `logo` snippet:

```svelte
{#snippet logo()}
  <LogoDark class="hidden h-7 dark:block" />
  <LogoLight class="block h-7 dark:hidden" />
  <span class="sr-only">Project Name</span>
{/snippet}

<DocsLayout navigation={n} {version}>
  {@render children?.()}
</DocsLayout>
```

Place `LogoDark` and `LogoLight` Svelte components or `<img>` tags referencing SVGs from `docs/static/`.

## External Fonts

Configured in `docs/src/app.html`:

```html
<link rel="preconnect" href="https://fonts.googleapis.com" />
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
<link
  href="https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&family=JetBrains+Mono:ital,wght@0,100..800;1,100..800&display=fallback"
  rel="stylesheet"
/>
```

Default fonts: **Inter** (UI) and **JetBrains Mono** (code). Replace links in `app.html` for custom fonts.
