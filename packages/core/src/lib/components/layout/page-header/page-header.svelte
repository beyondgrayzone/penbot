<script lang="ts">
	import Separator from "$lib/components/ui/separator/separator.svelte";
	import PropField from "$lib/components/prop-field.svelte";
	import type { PrimitiveAttributes } from "$lib/types.js";
	import { cn } from "$lib/utils.js";
	import PageHeaderDescription from "./page-header-description.svelte";
	import PageHeaderHeading from "./page-header-heading.svelte";

	let {
		class: className,
		children,
		availableSinceVersion,
		title,
		description = "",
		...restProps
	}: PrimitiveAttributes & {
		title: string;
		availableSinceVersion: string | undefined;
		description?: string;
	} = $props();

</script>

<div class={cn("relative", className)} {...restProps}>
	{#if children}
		{@render children?.()}
	{:else}
		<PageHeaderHeading>
			{title}
			{#if availableSinceVersion}
				<PropField name={"Available since"} type={availableSinceVersion} />
			{/if}
		</PageHeaderHeading>

		{#if description}
			<PageHeaderDescription>{description}</PageHeaderDescription>
		{/if}
		{#if title || description}
			<Separator class="mt-6" />
		{/if}
	{/if}
</div>
