# Tabs
> Break content into multiple panes to reduce cognitive load.



You can use the `Tabs` and `TabItem` components to create tabbed interfaces. A `label` prop must be provided to each `TabItem` which will be used to display the label. Whichever tab should be active by default is specified by the `value` prop on the `Tabs` component.

## Usage

```svelte title="document.md"
<script>
	import { Tabs, TabItem } from "@bladocs/core";
	const items = ["First tab", "Second tab"];
</script>

<Tabs value="First tab" {items}>
	<TabItem value="First tab">This is the first tab's content.</TabItem>
	<TabItem value="Second tab">This is the second tab's content.</TabItem>
</Tabs>
```

## Examples

### Simple Text

### First tab
This is the first tab's content.

### Second tab
This is the second tab's content.

### Markdown Syntax

### +page.svelte
```svelte
<script lang="ts">
	import { Button } from "@bladocs/core";
</script>

<Button onclick={() => alert("Hello!")}>Click me</Button>
```


### +page.server.ts
```ts
export async function load() {
	return {
		transactions: [],
	};
}
```


## 

> If you plan to use markdown-specific syntax in your tabs, ensure you include a space between the component and the content in your Markdown file.

## Props

### Tabs

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `items` | `string[]` | `` | The tab items to display. |

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `value` | `string` | `` | The label of the tab to be active by default. |

### TabItem

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `value` | `string` | `` | The value that identifies the tab. This value should map to an item within the `items` prop passed to the `Tabs` component. |
