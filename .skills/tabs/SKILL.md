# Tabs

**Import:** `import { Tabs, TabItem } from "@bladocs/core";`

## Usage

```svelte
<script>
  const items = ["First tab", "Second tab"];
</script>

<Tabs value="First tab" {items}>
  <TabItem value="First tab">Content for first tab.</TabItem>
  <TabItem value="Second tab">Content for second tab.</TabItem>
</Tabs>
```

### With markdown/code blocks

Leave space between components and content in `.md` files:

```svelte
<Tabs items={items}>
  <TabItem value="+page.svelte">
    
    ```svelte
    <Button>Click me</Button>
    ```
    
  </TabItem>
  <TabItem value="+page.server.ts">
    
    ```ts
    export async function load() { ... }
    ```
    
  </TabItem>
</Tabs>
```

## Props

### Tabs
| Prop    | Type       | Required | Default    | Description              |
|---------|------------|----------|------------|--------------------------|
| `items` | `string[]` | ✅       | —          | Tab label array          |
| `value` | `string`   | —        | `items[0]` | Initially active tab     |

### TabItem
| Prop    | Type     | Required | Description                                  |
|---------|----------|----------|----------------------------------------------|
| `value` | `string` | ✅       | Must match one of the `items` in parent Tabs |
