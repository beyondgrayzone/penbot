---
title: Select
description: A select component to use in examples and documentation.
section: Components
---

<script>
	import { Select, CardContainer } from "@bladocs/core";
	import SelectDemo from "$lib/components/demos/select-demo.svelte";
</script>

## Usage

```svelte title="document.md"
<script>
	import { Select } from "@bladocs/core";
</script>

<Select>
	<!-- ... -->
</Select>
```

## Example

<CardContainer class="flex items-center gap-2.5 flex-wrap">
	<SelectDemo />
</CardContainer>

## Props

See [Bits UI Select](https://bits-ui.com/docs/v1/components/select) for available props.
