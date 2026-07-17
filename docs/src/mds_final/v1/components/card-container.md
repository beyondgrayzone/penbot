# CardContainer
> Display a container with a border and a background color for examples/demos.



Often times you'll want to display some demo/example components in a container. The `CardContainer` component is a great way to do this, as it aligns effortlessly with the rest of the docs theme.

## Usage

```svelte title="document.md"
<script>
	import { CardContainer, Button } from "@penbot/ui";
</script>

<CardContainer class="flex flex-wrap gap-4">
	<Button variant="default">Default</Button>
	<Button variant="brand">Brand</Button>
	<Button variant="outline">Outline</Button>
	<Button variant="ghost">Ghost</Button>
	<Button variant="subtle">Subtle</Button>
	<Button variant="link">Link</Button>
</CardContainer>
```

## Example

