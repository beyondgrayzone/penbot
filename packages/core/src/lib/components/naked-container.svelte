<script lang="ts">
	import { cn } from "$lib/utils.js";
	import ScrollArea from "$lib/components/ui/scroll-area/scroll-area.svelte";
	import type { ComponentProps } from "svelte";

	let { children, class: className, ...restProps }: ComponentProps<typeof ScrollArea> = $props();
</script>

<ScrollArea
	orientation="both"
	class={cn(
		""
	)}
	viewportClasses="rounded-none"
	{...restProps}
>
	<div class={cn("p-6", className)}>
		{@render children?.()}
	</div>
</ScrollArea>
