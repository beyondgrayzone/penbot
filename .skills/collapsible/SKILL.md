# Collapsible

**Import:** `import { Collapsible } from "@penbot/core";`

## Usage

```svelte
<Collapsible title="more info">
  <!-- space here so MD renders -->
  Content that can be shown/hidden.
  <!-- space here so MD renders -->
</Collapsible>
```

## Props

| Prop             | Type      | Default   | Description                                    |
|------------------|-----------|-----------|------------------------------------------------|
| `title`          | `string`  | —         | Trigger label ("Show"/"Hide" prefix auto-added)|
| `open`           | `boolean` | `false`   | Whether initially open                         |
| `triggerContent` | `Snippet` | —         | Override the trigger button content            |
| `children`       | `Snippet` | —         | Collapsible body content                       |
