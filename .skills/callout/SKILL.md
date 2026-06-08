# Callout

**Import:** `import { Callout } from "@bladocs/core";`

## Usage

```svelte
<Callout type="note" title="Note">
  <!-- Space here so MD renders -->
  This is a note callout.
  <!-- Space here so MD renders -->
</Callout>
```

## Types

| Type      | Description                          |
|-----------|--------------------------------------|
| `note`    | Highlight important information      |
| `warning` | Warning about something              |
| `danger`  | Dangerous/error information          |
| `tip`     | Helpful tip                          |
| `success` | Success confirmation                 |

## Props

| Prop      | Type                                  | Default  | Description                    |
|-----------|---------------------------------------|----------|--------------------------------|
| `type`    | `'warning' \| 'note' \| 'danger' \| 'tip' \| 'success'` | `'note'` | Callout style |
| `title`   | `string`                              | —        | Override default title         |
| `icon`    | `Component`                           | —        | Custom icon component          |
| `children`| `Snippet`                             | —        | Callout body content           |
