# PropField

**Import:** `import { PropField } from "@penbot/core";`

## Usage

Display component props/parameters with name, type, and description.

```svelte
<PropField name="checked" type="boolean" defaultValue="false" required>
  The checked state of the checkbox.
</PropField>
```

### Nested objects with Collapsible

```svelte
<PropField name="options" type="CheckboxOptions" required>
  Configuration options.
  <Collapsible title="properties">
    <PropField name="width" type="number" required>
      The width to apply.
    </PropField>
  </Collapsible>
</PropField>
```

## Props

| Prop           | Type      | Required | Description                  |
|----------------|-----------|----------|------------------------------|
| `name`         | `string`  | ✅       | Prop name                    |
| `type`         | `string`  | ✅       | Prop type                    |
| `defaultValue` | `string`  | —        | Default value                |
| `required`     | `boolean` | —        | Whether prop is required     |
| `children`     | `Snippet` | —        | Description/content          |
