---
title: Card
description: A flexible card component with support for icons, links, and horizontal layout.
---

# Card

**Import:** `import { Card } from "@penbot/core";`

## Usage

### With icon
```svelte
<Card title="Title" icon={RocketLaunch}>
  Content here.
</Card>
```

### As a link
```svelte
<Card title="Link card" href="/docs" icon={RocketLaunch}>
  Content here.
</Card>
```

### Without icon
```svelte
<Card title="Title">
  Content here.
</Card>
```

### Horizontal layout
```svelte
<Card title="Horizontal card" horizontal icon={RocketLaunch}>
  Content here.
</Card>
```

## Props

| Prop       | Type        | Required | Description                                              |
|------------|-------------|----------|----------------------------------------------------------|
| `title`    | `string`    | ✅       | Card title                                               |
| `icon`     | `Component` | —        | Optional icon                                            |
| `href`     | `string`    | —        | Makes card a link (auto-handles `target`)                |
| `horizontal`| `boolean`  | —        | Horizontal layout                                        |
| `children` | `Snippet`   | —        | Card body content                                        |
