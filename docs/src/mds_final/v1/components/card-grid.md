# Card Grid
> Display a grid of cards.



Use the `CardGrid` component to display a grid of [`Card`](/docs/v1/components/card) components.

## Usage

```svelte title="document.md"
<script>
	import { CardGrid, Card } from "@penbot/ui";
</script>

<CardGrid cols={2}>
	<Card title="This is a card">
		You can use markdown in here, just ensure to include a space between the component and the
		content in your Markdown file.
	</Card>
	<Card title="This is another card">
		You can use markdown in here, just ensure to include a space between the component and the
		content in your Markdown file.
	</Card>
	<Card title="This is a third card">
		You can use markdown in here, just ensure to include a space between the component and the
		content in your Markdown file.
	</Card>
	<Card title="This is a fourth card" href="/docs">
		You can use markdown in here, just ensure to include a space between the component and the
		content in your Markdown file.
	</Card>
</CardGrid>
```

## Examples

### 2 Columns (default)

```svelte
<CardGrid>
	<!-- ... cards here-->
</CardGrid>
```

## This is a card

You can use markdown in here, just ensure to include a space between the component and the
		content in your Markdown file.

## This is another card

You can use markdown in here, just ensure to include a space between the component and the
		content in your Markdown file.

## This is a third card

You can use markdown in here, just ensure to include a space between the component and the
		content in your Markdown file.

## [This is a fourth card](/docs)

You can use markdown in here, just ensure to include a space between the component and the
		content in your Markdown file.

### 3 Columns

```svelte
<CardGrid cols={3}>
	<!-- ... cards here-->
</CardGrid>
```

## This is a card

You can use markdown in here, just ensure to include a space between the component and the
		content in your Markdown file.

## This is another card

You can use markdown in here, just ensure to include a space between the component and the
		content in your Markdown file.

## [This is a third card](/docs)

You can use markdown in here, just ensure to include a space between the component and the
		content in your Markdown file.

## Props

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `cols` | `number` | `2` | The number of columns to display the cards in. Uses flex column layout when in smaller viewports. |
