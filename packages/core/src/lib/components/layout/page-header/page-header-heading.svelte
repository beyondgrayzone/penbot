<script lang="ts">
	import { cn } from "$lib/utils.js";
	import type { PrimitiveHeadingAttributes } from "$lib/types.js";
	import Copy from "../copy.svelte";

	let { class: className, children, ...restProps }: PrimitiveHeadingAttributes = $props();
</script>

<div class="flex items-top gap-10 justify-between">
	<h1 class={cn("scroll-m-20 text-[30px] font-semibold", className)} {...restProps}>
		{@render children?.()}
	</h1>

	<Copy />
</div>
