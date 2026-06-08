# Steps
> Display a series of series of steps.



The `Steps` and `Step` components are used to display a series of steps, breaking down a process into more manageable chunks.

## Usage

```svelte title="document.md"
<script>
	import { Steps, Step } from "$lib/components";
</script>

<Steps>
	<Step>Install the package</Step>

	You can install the project via `npm` or `pnpm`.

	<!-- Code block here -->

	<Step>Start your engines</Step>

	You can start the project by running `npm run dev` or `pnpm run dev`.

	<!-- Code block here -->
</Steps>
```

## Example

## Install the package

> You can install the project via `npm` or `pnpm`.

```bash
npm install @bladocs/ui
```


## Start your engines

> You can start the project by running `npm run dev` or `pnpm dev`.

```bash
npm run dev
```


## 

> If you plan to use markdown-specific syntax in your steps, ensure you include a space between the component and the content in your Markdown file.
