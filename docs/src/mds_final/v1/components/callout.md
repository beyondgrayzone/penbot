# Callout
> A callout component to highlight important information.



Callouts (also known as _admonitions_) are used to highlight a block of text. There are five types of callouts available: `'note'`, `'warning'`, `'danger'`, `'tip'`, and `'success'`.

You can override the default icon for the callout by passing a component via the `icon` prop.

## Usage

```svelte title="document.md"
<script>
	import { Callout } from "$lib/components";
</script>

<Callout type="note" title="Note">
	<!-- Space here-->
	This is a note, used to highlight important information or provide additional context. You can use
	markdown in here as well! Just ensure you include a space between the component and the content in
	your Markdown file.
	<!-- Space here-->
</Callout>
```

## Examples

### Warning


## 

> This is an example of a warning callout.

### Note


## 

> This is an example of a note callout.

### Danger


## 

> This is an example of a danger callout.

### Tip


## 

> This is an example of a tip callout.

### Success


## 

> This is an example of a success callout.

### Custom Icon


## 

> This is an example of a note callout with a custom icon.

### Custom Title


## Tread carefully

> This is an example of a warning callout with a custom title.

## Props

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `type` | `'warning' | 'note' | 'danger' | 'tip' | 'success'` | `'note'` | The type of callout to display. |

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `title` | `string` | `` | Override the default title for the callout. |

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `icon` | `Component` | `` | Override the default icon for the callout. |

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `children` | `Snippet` | `` | The content to display inside of the callout's body. |
