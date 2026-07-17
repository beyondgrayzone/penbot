# Collapsible
> Show and hide content in a collapsible container.



## Usage

```svelte title="document.md"
<script>
	import { Collapsible, CardContainer } from "@penbot/core";
</script>

<Collapsible title="more info">
	<!-- space here so MD can render -->
	To learn more about SvelteKit, check out the [SvelteKit documentation](https://svelte.dev/kit).
	<!-- space here so MD can render -->
</Collapsible>
```

## Example


## more info
To learn more about SvelteKit, check out the [SvelteKit documentation](https://svelte.dev/kit).

## Props

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `title` | `string` | `` | The title to display in the trigger. "Hide" and "Show" prefix will be added automatically. |

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `open` | `boolean` | `false` | Whether the content should be open or closed. |

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `triggerContent` | `Snippet` | `` | Override the content inside of the trigger button. |

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `children` | `Snippet` | `` | The content that is collapsible. |
