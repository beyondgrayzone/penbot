# PropField
> Display a prop field with a name, type, and description.



Use the `PropField` component to annotate props/params in your documentation.

## Usage

```svelte title="document.md"
<script>
	import { PropField } from "@bladocs/core";
</script>

<PropField name="checked" type="boolean" required defaultValue="false">
	<!-- Space here-->
	The checked state of the checkbox.
	<!-- Space here-->
</PropField>
```

## Examples

### Basic

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `checked` | `boolean` | `false` | The checked state of the checkbox. |

### Object

You can use `PropField` in combination with the [`Collapsible`](/docs/v2/components/collapsible) component to represent more complex types.

```svelte title="document.md"
<script>
	import { PropField, Collapsible } from "@bladocs/core";
</script>

<PropField name="options" type="CheckboxOptions" required>
	<!-- Space here -->
	Configuration options to customize the behavior of the `Checkbox` component.
	<!-- Space here -->
	<Collapsible title="properties">
		<PropField name="width" type="number" required>
			The width to apply to the checkbox.
		</PropField>
		<PropField name="height" type="number" required defaultValue="20">
			The height to apply to the checkbox.
		</PropField>
	</Collapsible>
</PropField>
```

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `options` | `CheckboxOptions` | `` | Configuration options to customize the behavior of the `Checkbox` component. |

## properties
| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `width` | `number` | `` | The width to apply to the checkbox. |
| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `height` | `number` | `20` | The height to apply to the checkbox. |
## Some Really Long Title That will Wrap

## Props

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `name` | `string` | `` | The name of the prop. |

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `type` | `string` | `` | The type of the prop. |

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `defaultValue` | `string` | `` | The default value of the prop. |

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `required` | `boolean` | `false` | Whether the prop is required. |

| Prop | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `children` | `Snippet` | `` | The description/content to display within the prop field. |
