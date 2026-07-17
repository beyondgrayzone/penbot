# Theme
> Learn how to customize the theme in your Penbot project.


The theme determines the branded color scheme for your site. A theme for each of the TailwindCSS colors is provided by the `@penbot/core` package. Each theme has been designed to present well in both light and dark mode.

## Using a theme

To use a theme, import the theme file into your `src/app.css` file _before_ importing the `@penbot/core/globals.css` file.

```css
/* @import "@penbot/core/theme-orange.css"; */
@import "@penbot/core/theme-emerald.css";
@import "@penbot/core/globals.css";
```

It's not recommended to customize the theme to maintain consistency across the UI components that are provided by Penbot and align with the provided themes.

## Available themes

| Theme name | Import path                        |
| ---------- | ---------------------------------- |
| orange     | `@penbot/core/theme-orange.css`  |
| green      | `@penbot/core/theme-green.css`   |
| blue       | `@penbot/core/theme-blue.css`    |
| purple     | `@penbot/core/theme-purple.css`  |
| pink       | `@penbot/core/theme-pink.css`    |
| lime       | `@penbot/core/theme-lime.css`    |
| yellow     | `@penbot/core/theme-yellow.css`  |
| cyan       | `@penbot/core/theme-cyan.css`    |
| teal       | `@penbot/core/theme-teal.css`    |
| violet     | `@penbot/core/theme-violet.css`  |
| amber      | `@penbot/core/theme-amber.css`   |
| red        | `@penbot/core/theme-red.css`     |
| sky        | `@penbot/core/theme-sky.css`     |
| emerald    | `@penbot/core/theme-emerald.css` |
| fuchsia    | `@penbot/core/theme-fuchsia.css` |
| rose       | `@penbot/core/theme-rose.css`    |

## Tailwind Variables

Penbot uses TailwindCSS to style the UI components and provides a set of Tailwind variables that can be used to style your examples/custom components.

### Gray

We override the TailwindCSS `gray` color scale to provide our own grays.

### Brand

You can use the `brand` color to use the brand color of your project.
