# NakedContainer
> Display a container with a border and a background color for examples/demos.



Often times you'll want to render components themselves in a container. The `NakedContainer` component allows to do exactly that.

## Usage

```svelte title="document.md"
<script>
	import { NakedContainer, Button } from "@bladocs/ui";
</script>

<NakedContainer class="flex flex-wrap gap-4">
	<Button variant="default">Default</Button>
	<Button variant="brand">Brand</Button>
	<Button variant="outline">Outline</Button>
	<Button variant="ghost">Ghost</Button>
	<Button variant="subtle">Subtle</Button>
	<Button variant="link">Link</Button>
</NakedContainer>
```

## Example

