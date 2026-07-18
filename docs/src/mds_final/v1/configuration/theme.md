# Theme
> Learn how to customize and create themes in your Penbot project.


Penbot ships with a **multi-theme system** powered by CSS custom properties and `data-theme` attributes. Pick a built-in theme, switch at runtime, or build your own.

## Quick Start: Change the Default Theme

You only need to change **one file**  `docs/src/routes/+layout.svelte`:

```svelte
<script>
	import { ModeWatcher } from "mode-watcher";
</script>

<ModeWatcher defaultTheme="oceanic" />   <!-- change this -->
```

Available values: `oceanic`, `forest`, `sober-1`, `amber`, `blue`, `cyan`, `emerald`, `fuchsia`, `green`, `indigo`, `lime`, `orange`, `pink`, `purple`, `red`, `rose`, `sky`, `teal`, `violet`, `yellow`.

Optionally, update the static fallback in `docs/src/app.html` for the initial HTML snapshot before JavaScript runs:

```html
<html lang="en" data-theme="oceanic">   <!-- match the defaultTheme -->
```

> **First-time switch?** Clear your browser's localStorage (`mode-watcher-theme` key in DevTools → Application → Local Storage) or use an incognito window. `ModeWatcher` respects stored user preferences over the default.

## How It Works

### Architecture

`ModeWatcher` from the `mode-watcher` library (`v1.1.0`) is the single source of truth. On page load:

1. `ModeWatcher` injects an inline `<script>` into `<head>` that calls `setInitialMode()` **before the page paints**
2. That script reads `localStorage.getItem("mode-watcher-theme")`  if it exists, that's the theme; if not, it falls back to the `defaultTheme` prop
3. It sets `data-theme="<value>"` on `<html>` and a `.dark` class based on light/dark mode preference
4. CSS `[data-theme="..."]` selectors activate the correct color variables
5. On subsequent visits, the stored preference is used automatically

No custom inline scripts needed  `ModeWatcher` handles everything (persistence, FOUC prevention, dark mode).

### CSS Import Order (Critical)

In `docs/src/app.css`, themes must be imported **after** globals:

```css
@import "@penbot/core/globals.css";         /* ← MUST come FIRST (fallback :root rules) */
@import "@penbot/core/theme-oceanic.css";    /* [data-theme="oceanic"] rules */
@import "@penbot/core/theme-forest.css";     /* [data-theme="forest"] rules */
@import "@penbot/core/theme-sober-1.css";    /* [data-theme="sober-1"] rules */
@import "@penbot/core/theme-amber.css";      /* brand-only themes */
/* ... rest of themes ... */
```

**Why this order matters:** `globals.css` defines `:root { --theme-color-background: … }` as the fallback. Theme files define `[data-theme="<name>"] { --theme-color-background: … }`. Both have the same specificity, so the **last declared** wins. Themes must come AFTER globals so their `[data-theme]` selectors override the `:root` fallback.

### Two Distinct Token Layers

| Layer | Prefix | Purpose |
|-------|--------|---------|
| **Scale palette** | `--theme-color-current-50` … `950` | Raw 11-stop color scale |
| **Semantic UI tokens** | `--theme-color-background`, `--theme-color-primary`, etc. | Purpose-driven tokens used by components |

## Available Themes

### Full Semantic Themes

These define every UI token  backgrounds, text, borders, primaries, accents, destructive colors, and brand colors. Use these as your main theme.

| Theme | Key | Import path |
|-------|-----|-------------|
| Oceanic Blue | `oceanic` | `@penbot/core/theme-oceanic.css` |
| Forest Green | `forest` | `@penbot/core/theme-forest.css` |
| Sober (grayscale) | `sober-1` | `@penbot/core/theme-sober-1.css` |

### Brand-Only Themes

These only define `--theme-color-brand-*` tokens (buttons, links, accents) and the scale palette. They inherit all other semantic UI tokens from `globals.css`'s `:root` fallback. Use these to change accent colors only.

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

## Dark Mode

Dark mode works automatically alongside themes. `ModeWatcher` toggles the `.dark` class on `<html>`, and CSS composes it with the active `data-theme`:

```
[data-theme="oceanic"]          → light Oceanic
[data-theme="oceanic"].dark     → dark Oceanic
[data-theme="forest"]           → light Forest
[data-theme="forest"].dark      → dark Forest
```

Each theme CSS file defines both light and dark variants:

```css
[data-theme="oceanic"] {
	/* Light mode tokens */
	--theme-color-background: var(--theme-color-current-50);
	/* ... */
}

[data-theme="oceanic"].dark,
.dark [data-theme="oceanic"] {
	/* Dark mode tokens */
	--theme-color-background: var(--theme-color-current-950);
	/* ... */
}
```

## Available CSS Tokens

### Full Semantic Tokens (oceanic, forest, sober-1)

Defined in both `[data-theme="..."]` light and `[data-theme="..."].dark` selectors.

| Token | Purpose |
|-------|---------|
| `--theme-color-background` | Page background |
| `--theme-color-background-secondary` | Card/surface background |
| `--theme-color-foreground` | Body text color |
| `--theme-color-muted` | Muted surface |
| `--theme-color-muted-foreground` | Text on muted surfaces |
| `--theme-color-border` | Default border color |
| `--theme-color-primary` | Primary button / CTA |
| `--theme-color-primary-foreground` | Text on primary |
| `--theme-color-primary-hover` | Primary hover state |
| `--theme-color-primary-active` | Primary pressed state |
| `--theme-color-secondary` | Secondary button |
| `--theme-color-secondary-foreground` | Text on secondary |
| `--theme-color-accent` | Accent highlights |
| `--theme-color-accent-foreground` | Text on accent |
| `--theme-color-destructive` | Delete / danger actions |
| `--theme-color-destructive-foreground` | Text on destructive |
| `--theme-color-destructive-border` | Destructive border |

### Scale Palette (All Themes)

| Token | Range |
|-------|-------|
| `--theme-color-current-50` … `950` | 11-stop scale (lightest → darkest) |

### Brand Tokens (All Themes)

| Token | Purpose |
|-------|---------|
| `--theme-color-brand-50` … `950` | Brand color scale |
| `--theme-color-brand` | Default brand color |
| `--theme-color-brand-border` | Brand borders |
| `--theme-color-brand-hover` | Brand hover state |
| `--theme-color-brand-foreground` | Text on brand |
| `--theme-color-brand-link` | Link color |
| `--theme-color-brand-link-hover` | Link hover |
| `--theme-color-brand-code-link` | Inline code link |
| `--theme-color-brand-code-link-hover` | Inline code link hover |

## Tailwind v4 Integration

Penbot uses Tailwind CSS v4's `@theme` directive in `globals.css` to expose all theme tokens as Tailwind utility classes. This means you can use any token directly in your HTML:

```html
<!-- Using Tailwind utility classes mapped to theme tokens -->
	<button class="bg-primary text-primary-foreground hover:bg-primary-hover">
		Click me
	</button>
	<a class="text-brand-link hover:text-brand-link-hover">Link</a>
```

The gray scale is also overridden to match the active theme:

```html
```

Available Tailwind utility classes from theme tokens:

| Tailwind class | Maps to |
|----------------|---------|
| `bg-background` | `--theme-color-background` |
| `bg-background-secondary` | `--theme-color-background-secondary` |
| `text-foreground` | `--theme-color-foreground` |
| `bg-muted` / `text-muted-foreground` | `--theme-color-muted` / `--theme-color-muted-foreground` |
| `border-border` | `--theme-color-border` |
| `bg-primary` / `text-primary-foreground` | Primary button colors |
| `hover:bg-primary-hover` | Primary hover state |
| `bg-secondary` / `text-secondary-foreground` | Secondary colors |
| `bg-accent` / `text-accent-foreground` | Accent colors |
| `bg-destructive` / `text-destructive-foreground` | Destructive colors |
| `border-destructive-border` | Destructive border |
| `text-brand-link` / `hover:text-brand-link-hover` | Brand link colors |
| `bg-brand` / `text-brand-foreground` | Brand backgrounds |
| `text-brand` | Brand text color |
| `bg-gray-50` … `bg-gray-950` | Current theme's scale palette |
| `text-brand-50` … `text-brand-950` | Brand color scale |

## Using Tokens in Your Own CSS

### In Svelte Components

```svelte
<button
	style="background: var(--theme-color-primary); color: var(--theme-color-primary-foreground)"
>
	Click me
</button>
```

### In Tailwind (Arbitrary Values)

```html
	…
```

### In Plain CSS

```css
.my-card {
	background: var(--theme-color-background-secondary);
	border: 1px solid var(--theme-color-border);
	color: var(--theme-color-foreground);
}
.my-card a {
	color: var(--theme-color-brand-link);
}
.my-card a:hover {
	color: var(--theme-color-brand-link-hover);
}
```

## Creating a Custom Theme

### Brand-Only Theme (Quick)

Create `packages/core/src/lib/styles/theme-coral.css`:

```css
[data-theme="coral"] {
	--theme-color-current-50: var(--color-red-50);
	--theme-color-current-100: var(--color-red-100);
	/* ... through 950 ... */
	--theme-color-current-950: var(--color-red-950);

	--theme-color-brand-50: var(--color-orange-50);
	/* ... through 950 ... */
	--theme-color-brand-950: var(--color-orange-950);

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
	--theme-color-brand: var(--theme-color-brand-600);
	--theme-color-brand-border: var(--theme-color-brand-700);
	--theme-color-brand-hover: var(--theme-color-brand-500);
	--theme-color-brand-foreground: var(--theme-color-foreground);
	--theme-color-brand-link: var(--theme-color-brand-500);
	--theme-color-brand-link-hover: var(--theme-color-brand-400);
	--theme-color-brand-code-link: var(--theme-color-brand-400);
	--theme-color-brand-code-link-hover: var(--theme-color-brand-300);
}
```

### Full Semantic Theme

Define every UI token from scratch. See `theme-oceanic.css` for a complete reference. Use OKLCH or any CSS color format.

### Wiring It Up

1. **Register the export** in `packages/core/package.json`:
   ```json
   "./theme-coral.css": "./dist/styles/theme-coral.css"
   ```

2. **Import it** in `docs/src/app.css` (after `globals.css`):
   ```css
   @import "@penbot/core/globals.css";
   @import "@penbot/core/theme-coral.css";
   ```

3. **Set it as default** in `docs/src/routes/+layout.svelte`:
   ```svelte
   <ModeWatcher defaultTheme="coral" />
   ```

4. **Update the static fallback** in `docs/src/app.html`:
   ```html
   <html lang="en" data-theme="coral">
   ```

5. **Rebuild**: `bun run build`

## Switching Themes at Runtime

Build a theme picker UI by changing `data-theme` and persisting with `mode-watcher`'s `setTheme()`:

```svelte
<script>
	import { setTheme } from "mode-watcher";
</script>

<select onchange={(e) => setTheme(e.target.value)}>
	<option value="oceanic">Oceanic</option>
	<option value="forest">Forest</option>
	<option value="sober-1">Sober</option>
	<option value="amber">Amber</option>
</select>
```

`mode-watcher` handles localStorage persistence automatically.

## Tips

- **No custom scripts**: `ModeWatcher` handles FOUC prevention, persistence, and dark mode. Don't add your own inline scripts for theme management.
- **Don't hardcode colors**: always use `var(--theme-color-*)` or Tailwind utility classes in your components, never raw hex/rgb/oklch values.
- **Test both modes**: every theme has light and dark variants. Toggle `.dark` on `<html>` in DevTools.
- **Keep themes self-contained**: don't `@import` other theme files inside a theme CSS file.
- **Locations that matter**:
  - `docs/src/routes/+layout.svelte` → `<ModeWatcher defaultTheme="…" />` (primary setting)
  - `docs/src/app.html` → `<html data-theme="…">` (static fallback only)
  - `docs/src/app.css` → `@import` order (globals first, themes after)
  - `packages/core/package.json` → export entry for each theme
