---
title: Theme
description: Learn how to customize the theme in your Bladocs project.
section: Configuration
---

The theme determines the branded color scheme for your site. A theme for each of the TailwindCSS colors is provided by the `@bladocs/core` package. Each theme has been designed to present well in both light and dark mode.

## Using a theme

To use a theme, import the theme file into your `src/app.css` file _before_ importing the `@bladocs/core/globals.css` file.

```css
/* @import "@bladocs/core/theme-orange.css"; */
@import "@bladocs/core/theme-emerald.css";
@import "@bladocs/core/globals.css";
```

It's not recommended to customize the theme to maintain consistency across the UI components that are provided by Bladocs and align with the provided themes.

## Available themes

| Theme name | Import path                        |
| ---------- | ---------------------------------- |
| orange     | `@bladocs/core/theme-orange.css`  |
| green      | `@bladocs/core/theme-green.css`   |
| blue       | `@bladocs/core/theme-blue.css`    |
| purple     | `@bladocs/core/theme-purple.css`  |
| pink       | `@bladocs/core/theme-pink.css`    |
| lime       | `@bladocs/core/theme-lime.css`    |
| yellow     | `@bladocs/core/theme-yellow.css`  |
| cyan       | `@bladocs/core/theme-cyan.css`    |
| teal       | `@bladocs/core/theme-teal.css`    |
| violet     | `@bladocs/core/theme-violet.css`  |
| amber      | `@bladocs/core/theme-amber.css`   |
| red        | `@bladocs/core/theme-red.css`     |
| sky        | `@bladocs/core/theme-sky.css`     |
| emerald    | `@bladocs/core/theme-emerald.css` |
| fuchsia    | `@bladocs/core/theme-fuchsia.css` |
| rose       | `@bladocs/core/theme-rose.css`    |

## Tailwind Variables

Bladocs uses TailwindCSS to style the UI components and provides a set of Tailwind variables that can be used to style your examples/custom components.

### Gray

We override the TailwindCSS `gray` color scale to provide our own grays.

### Brand

You can use the `brand` color to use the brand color of your project.
