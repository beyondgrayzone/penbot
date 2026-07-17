---
title: Checkbox
description: A checkbox component to use in examples and documentation.
section: Components
availableSinceVersion: 1.0.3
---

<script>
	import { Checkbox, Label, CardContainer, NakedContainer } from "@penbot/core";
</script>

## Usage

```svelte title="document.md"
<script>
	import { Checkbox } from "@penbot/core";
</script>

<Checkbox />
```

## Example

<NakedContainer class="flex flex-col gap-2.5 flex-wrap">
	<div class="flex flex-row gap-4">
			<Checkbox disabled checked />
			<Label for="bio">Bio</Label>
	</div>
	<div class="flex flex-row gap-4">
			<Checkbox disabled  />
			<Label for="bio">Profile</Label>
	</div>
</NakedContainer>

## Props

See [Bits UI Checkbox](https://bits-ui.com/docs/v1/components/checkbox) for available props.
