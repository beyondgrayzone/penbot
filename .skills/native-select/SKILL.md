# NativeSelect

**Import:** `import { NativeSelect, Label } from "@bladocs/core";`

## Usage

A styled native `<select>` element.

```svelte
<NativeSelect>
  <option value="1">Option 1</option>
  <option value="2">Option 2</option>
  <option value="3">Option 3</option>
</NativeSelect>
```

Combine with `Label` for accessibility:

```svelte
<Label for="options">Select an option</Label>
<NativeSelect id="options">
  <option value="1">Option 1</option>
</NativeSelect>
```
