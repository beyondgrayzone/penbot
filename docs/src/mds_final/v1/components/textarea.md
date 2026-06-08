# Textarea
> A textarea component to use in examples and documentation.



When building documentation, it's often necessary to provide users with a textarea to showcase a specific feature. The `Textarea` component is a great way to do this, as it aligns effortlessly with the rest of the docs theme. The `Label` component is also provided to help with accessibility.

## Usage

```svelte
<script>
	import { Textarea, Label } from "@bladocs/core";
</script>

<Label for="bio">Your bio</Label>
<Textarea id="bio" name="bio" />
```

# Plausible
```Textarea [label="", button=""]
	placeholder="Tell us"
```

## Example


