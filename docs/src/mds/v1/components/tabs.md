---
title: Tabs
description: Break content into multiple panes to reduce cognitive load.
section: Components
---

<script>
	import { Tabs, TabItem, Callout, PropField } from "@penbot/core";
	const itemsA = ["First tab", "Second tab"];
	const itemsB = ["+page.svelte", "+page.server.ts"];
</script>

You can use the `Tabs` and `TabItem` components to create tabbed interfaces. A `label` prop must be provided to each `TabItem` which will be used to display the label. Whichever tab should be active by default is specified by the `value` prop on the `Tabs` component.

## Usage

```svelte title="document.md"
<script>
	import { Tabs, TabItem } from "@penbot/core";
	const items = ["First tab", "Second tab"];
</script>

<Tabs value="First tab" {items}>
	<TabItem value="First tab">This is the first tab's content.</TabItem>
	<TabItem value="Second tab">This is the second tab's content.</TabItem>
</Tabs>
```
# Plausible
```Tabs [2]
	This is the first tab's content.
	This is the second tab's content.
```

## Examples

### Simple Text

<Tabs value={"First tab"} items={itemsA}>
	<TabItem value="First tab">This is the first tab's content.</TabItem>
	<TabItem value="Second tab">This is the second tab's content.</TabItem>
</Tabs>

### Markdown Syntax

<Tabs items={itemsB}>
<TabItem value="+page.svelte">

```svelte
<script lang="ts">
	import { Button } from "@penbot/core";
</script>

<Button onclick={() => alert("Hello!")}>Click me</Button>
```

</TabItem>

<TabItem value="+page.server.ts">

```ts
export async function load() {
	return {
		transactions: [],
	};
}
```

</TabItem>
</Tabs>

<Callout type="note" class="mt-8">
	If you plan to use markdown-specific syntax in your tabs, ensure you include a space between the component and the content in your Markdown file.
</Callout>

## Props

### Tabs

<PropField name="items" type="string[]" required>
The tab items to display.
</PropField>

<PropField name="value" type="string" default="items[0]">
The label of the tab to be active by default.
</PropField>

### TabItem

<PropField name="value" type="string" required>

The value that identifies the tab. This value should map to an item within the `items` prop passed to the `Tabs` component.

</PropField>
