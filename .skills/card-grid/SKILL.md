# CardGrid

**Import:** `import { CardGrid, Card } from "@bladocs/core";`

## Usage

Displays a grid of `Card` components.

```svelte
<CardGrid cols={2}>
  <Card title="Card title" icon={SomeIcon}>
    Card content here — leave space before/after in MD files.
  </Card>
  <Card title="Another card" href="/docs">
    More content.
  </Card>
</CardGrid>
```

## Props

| Prop   | Type     | Default | Description                      |
|--------|----------|---------|----------------------------------|
| `cols` | `number` | `2`     | Number of columns (responsive)   |
