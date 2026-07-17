# Input
> A form input component to use in examples and documentation.



When building documentation, it's often necessary to provide users with a form input to showcase a specific feature. The `Input` component is a great way to do this, as it aligns effortlessly with the rest of the docs theme. The `Label` component is also provided to help with accessibility.

## Usage

```svelte
<script>
	import { Input, Label } from "@penbot/core";
</script>

<Label for="name">Your name</Label>
<Input id="name" name="name" placeholder="John Doe" />
```

## Example

