---
title: Button
description: A button component to use in examples and documentation.
section: Components
---

<script>
	import { Button, CardContainer } from "@bladocs/core";
</script>

## Usage

```svelte title="document.md"
<script>
	import { Button } from "@bladocs/core";
</script>

<Button>Default</Button>
<Button variant="brand">Brand</Button>
<Button variant="ghost">Ghost</Button>
<Button variant="outline">Outline</Button>
<Button variant="subtle">Subtle</Button>
<Button variant="link">Link</Button>
```

## Example

### Default Size

<CardContainer class="flex items-center gap-2.5 flex-wrap">
	<Button>Default</Button>
	<Button variant="brand">Brand</Button>
	<Button variant="destructive">Destructive</Button>
	<Button variant="ghost">Ghost</Button>
	<Button variant="outline">Outline</Button>
	<Button variant="subtle">Subtle</Button>
	<Button variant="link">Link</Button>
</CardContainer>

### Small Size

<CardContainer class="flex items-center gap-4 flex-wrap">
	<Button size="sm">Default</Button>
	<Button variant="brand" size="sm">Brand</Button>
	<Button variant="destructive" size="sm">Destructive</Button>
	<Button variant="ghost" size="sm">Ghost</Button>
	<Button variant="outline" size="sm">Outline</Button>
	<Button variant="subtle" size="sm">Subtle</Button>
	<Button variant="link" size="sm">Link</Button>
</CardContainer>
