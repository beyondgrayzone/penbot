---
title: Theme
description: Comprehensive theming system supporting multiple themes, dark mode, and runtime switching.
---

# Theme

**Files:** `docs/src/routes/+layout.svelte`, `docs/src/app.css`, `docs/src/app.html`, `docs/src/lib/site-config.json`

## Usage

Penbot uses a `data-theme` attribute on `<html>` — not single-file imports. All themes are imported together. The active theme is controlled by `ModeWatcher`.

### Set the default theme

The default theme is defined in `docs/src/lib/site-config.json`:

```json
{
  "name": "Penbot",
  "defaultTheme": "forest",
  "themeTimestamp": 1784356785671
}
```

Use the CLI to change it (recommended — also updates the timestamp):

```bash
bun run docs:theme:list              # list available themes
bun run docs:theme:apply oceanic     # set default + update timestamp
```

Or edit the file directly and bump `themeTimestamp` manually.

Optionally match the static fallback in `docs/src/app.html`:

```html
<html lang="en" data-theme="forest">
```

### CSS import order (critical)

In `docs/src/app.css`, `globals.css` must come **first**, themes after:

```css
@import "@penbot/core/globals.css";          /* fallback :root rules — FIRST */
@import "@penbot/core/theme-oceanic.css";    /* [data-theme="oceanic"] rules */
@import "@penbot/core/theme-forest.css";
/* ... all other themes ... */
```

Themes come after globals so their `[data-theme]` selectors (same specificity as `:root`) win by source order.

### First-time switch

If you previously visited the site, clear localStorage (`mode-watcher-theme` key) or use an incognito window. Stored user preference overrides the default.

## How it works

1. `ModeWatcher` injects an inline script that sets `data-theme` on `<html>` before first paint
2. Reads `localStorage` → falls back to `defaultTheme` prop → sets `data-theme` attribute
3. `.dark` class is toggled independently for dark mode
4. CSS uses `[data-theme="oceanic"]` and `[data-theme="oceanic"].dark` selectors

## Available themes

### Full semantic themes (define all UI tokens)

| Theme | Key | Import path |
|-------|-----|-------------|
| Oceanic Blue | `oceanic` | `@penbot/core/theme-oceanic.css` |
| Forest Green | `forest` | `@penbot/core/theme-forest.css` |
| Sober (grayscale) | `sober-1` | `@penbot/core/theme-sober-1.css` |

### Brand-only themes (define only accent/brand tokens)

| Theme | Key | Import path |
|-------|-----|-------------|
| Amber | `amber` | `@penbot/core/theme-amber.css` |
| Blue | `blue` | `@penbot/core/theme-blue.css` |
| Cyan | `cyan` | `@penbot/core/theme-cyan.css` |
| Emerald | `emerald` | `@penbot/core/theme-emerald.css` |
| Fuchsia | `fuchsia` | `@penbot/core/theme-fuchsia.css` |
| Green | `green` | `@penbot/core/theme-green.css` |
| Indigo | `indigo` | `@penbot/core/theme-indigo.css` |
| Lime | `lime` | `@penbot/core/theme-lime.css` |
| Orange | `orange` | `@penbot/core/theme-orange.css` |
| Pink | `pink` | `@penbot/core/theme-pink.css` |
| Purple | `purple` | `@penbot/core/theme-purple.css` |
| Red | `red` | `@penbot/core/theme-red.css` |
| Rose | `rose` | `@penbot/core/theme-rose.css` |
| Sky | `sky` | `@penbot/core/theme-sky.css` |
| Teal | `teal` | `@penbot/core/theme-teal.css` |
| Violet | `violet` | `@penbot/core/theme-violet.css` |
| Yellow | `yellow` | `@penbot/core/theme-yellow.css` |

## Dark mode

Dark mode works alongside any theme. `ModeWatcher` toggles `.dark` on `<html>`. CSS composes both:

```
[data-theme="oceanic"]          → light Oceanic
[data-theme="oceanic"].dark     → dark Oceanic
[data-theme="forest"].dark      → dark Forest
```

## CSS tokens

### Semantic UI tokens (full themes only)

`--theme-color-background`, `--theme-color-background-secondary`, `--theme-color-foreground`, `--theme-color-muted`, `--theme-color-muted-foreground`, `--theme-color-border`, `--theme-color-primary`, `--theme-color-primary-foreground`, `--theme-color-primary-hover`, `--theme-color-primary-active`, `--theme-color-secondary`, `--theme-color-secondary-foreground`, `--theme-color-accent`, `--theme-color-accent-foreground`, `--theme-color-destructive`, `--theme-color-destructive-foreground`, `--theme-color-destructive-border`

### Scale palette (all themes)

`--theme-color-current-50` through `950` (11-stop scale, lightest → darkest)

### Brand tokens (all themes)

`--theme-color-brand-50` through `950`, `--theme-color-brand`, `--theme-color-brand-border`, `--theme-color-brand-hover`, `--theme-color-brand-foreground`, `--theme-color-brand-link`, `--theme-color-brand-link-hover`, `--theme-color-brand-code-link`, `--theme-color-brand-code-link-hover`

## Using tokens

### Tailwind (arbitrary values)

```html
<div class="bg-[var(--theme-color-background-secondary)] text-[var(--theme-color-foreground)]">
```

### Svelte components

```svelte
<button style="background: var(--theme-color-primary)">
```

### Tailwind utility classes

- `bg-background`, `text-foreground`, `bg-background-secondary`
- `bg-primary`, `text-primary-foreground`, `bg-muted`
- `text-brand`, `bg-brand`, `border-brand`
- Gray scale: `bg-gray-50` through `bg-gray-950` (maps to current theme's scale)

## Dev-only theme switcher

A theme picker dropdown is built into the header during development (`bun run dev`). It's automatically removed from production builds via `{#if dev}`.

## Server-side theme override via timestamp

Penbot supports forcing a specific theme on all returning visitors (e.g. after a redesign or a new default theme).

### How it works

1. `siteConfig.themeTimestamp` is a Unix timestamp (ms) baked into the HTML at build time via a `<meta>` tag
2. On page load, the client compares this server timestamp against a timestamp saved in `localStorage` (`penbot-theme-timestamp`)
3. **If the server timestamp is greater**, the client theme is overridden to the default from `siteConfig.defaultTheme`
4. **When the user manually picks a theme** from the dropdown, `Date.now()` is saved as the client timestamp, preventing the server from overriding again until the next redeploy

### Usage

Use the CLI to set a new default theme and timestamp together:

```bash
bun run docs:theme:apply oceanic
```

This updates both `defaultTheme` and `themeTimestamp` in `docs/src/lib/site-config.json` automatically using `Date.now()`.

Or edit the file directly:

```json
{
  "defaultTheme": "oceanic",
  "themeTimestamp": 1721300000000
}
```

Use a timestamp after your deployment — e.g. `Date.now()` at build time, or a deliberate number for the release date.

### Lifecycle

| Scenario | localStorage value | Server timestamp | Result |
|---|---|---|---|
| First visit | (none) → `1721300000000` | `1721300000000` | Theme reset to default |
| User picks a different theme | `Date.now()` (e.g. `1721301000000`) | `1721300000000` | Client > Server → preference preserved on reload |
| Server redeploy with `1721400000000` | `1721301000000` | `1721400000000` | Server > Client → theme overridden again |
| User picks a theme again | New `Date.now()` (e.g. `1721400500000`) | `1721400000000` | Preference preserved again |

### Implementation

See `docs/src/routes/+layout.svelte` for the full logic:

- **`onMount`**: reads client timestamp from localStorage, compares with `serverThemeTimestamp`, calls `setTheme()` if override needed, stamps server timestamp so override doesn't repeat
- **`$effect`** on `theme.current`: after mount, saves `Date.now()` only on user-initiated changes (skips the server-forced `setTheme` via the `pendingServerOverride` flag)

### mds theme CLI

A dedicated Go CLI manages the default theme and timestamp together:

```bash
bun run docs:theme:list              # list all themes, current default marked with *
bun run docs:theme:apply <name>      # set default + timestamp to now
bun run docs:theme                    # show help
```

Source: `packages/mds/mds/theme.go` — discovers themes by scanning `packages/core/src/lib/styles/theme-*.css`.

Package.json scripts:

| Script | Runs |
|---|---|
| `docs:theme` | `bun run --filter @penbot/mds theme` |
| `docs:theme:list` | `bun run --filter @penbot/mds theme list` |
| `docs:theme:apply` | `bun run --filter @penbot/mds theme apply <name>` |

## Adding a custom theme

### Brand-only (accent colors only)

Create `packages/core/src/lib/styles/theme-coral.css`:

```css
[data-theme="coral"] {
  --theme-color-current-50: var(--color-red-50);
  /* ... through 950 ... */

  --theme-color-brand-50: var(--color-orange-50);
  /* ... through 950 ... */

  --theme-color-brand: var(--theme-color-brand-600);
  --theme-color-brand-border: var(--theme-color-brand-700);
  --theme-color-brand-hover: var(--theme-color-brand-500);
  --theme-color-brand-foreground: var(--theme-color-background);
  --theme-color-brand-link: var(--theme-color-brand-600);
  --theme-color-brand-link-hover: var(--theme-color-brand-700);
  --theme-color-brand-code-link: var(--theme-color-brand-600);
  --theme-color-brand-code-link-hover: var(--theme-color-brand-500);
}

[data-theme="coral"].dark,
.dark [data-theme="coral"] {
  --theme-color-brand-foreground: var(--theme-color-foreground);
  --theme-color-brand-link: var(--theme-color-brand-500);
  --theme-color-brand-link-hover: var(--theme-color-brand-400);
  --theme-color-brand-code-link: var(--theme-color-brand-400);
  --theme-color-brand-code-link-hover: var(--theme-color-brand-300);
}
```

### Full semantic (all UI tokens)

Define every token from scratch. Reference `theme-oceanic.css` in the source.

### Wiring it up

1. Add export to `packages/core/package.json`
2. Add `@import "@penbot/core/theme-coral.css"` to `docs/src/app.css` (after all other theme imports)
3. Set `defaultTheme` in `docs/src/lib/site-config.json`:
   ```json
   "defaultTheme": "coral",
   "themeTimestamp": 1721300000000
   ```
   Or use the CLI: `bun run docs:theme:apply coral`
4. Rebuild: `bun run build`

## Runtime theme switching

```svelte
<script>
  import { setTheme } from "mode-watcher";
</script>
<select onchange={(e) => setTheme(e.target.value)}>
  <option value="oceanic">Oceanic</option>
  <option value="forest">Forest</option>
</select>
```

`mode-watcher` handles localStorage persistence automatically. The `storage` event syncs across open tabs.