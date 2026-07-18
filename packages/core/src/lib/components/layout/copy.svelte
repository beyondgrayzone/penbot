<script>
	import { page } from "$app/state";
	import Check from "lucide-svelte/icons/check";
	import Copy from "lucide-svelte/icons/copy";
	import Button from "../ui/button/button.svelte";

	let copied = $state(false);
	let isFetching = $state(false);
	let fetchSegment = $derived.by(() => {
		const pathPart = page.url.href.split("docs/")[1];
		const slug = pathPart.split("/").join("_");
		if (!slug.includes("_")) {
			return slug + "_index";
		}
		return slug;
	});

	async function copyToClipboard() {
		if (isFetching) return;
		try {
			isFetching = true;

			const response = await fetch(`/api/docs/${fetchSegment}`);

			if (!response.ok) throw new Error("Markdown file not found");

			const content = await response.text();
			await navigator.clipboard.writeText(content);
			copied = true;

			setTimeout(() => {
				copied = false;
			}, 2000);
		} catch (err) {
			console.error("Failed to copy text: ", err);
		} finally {
			isFetching = false;
		}
	}
</script>

<Button
	variant="outline"
	size="sm"
	class="group relative cursor-pointer gap-1.5 px-3 text-xs transition-all duration-200"
	onclick={copyToClipboard}
	aria-label="Copy doc"
>
	{#if copied}
		<Check class="size-3.5 text-emerald-500 transition-all duration-200" />
		<span class="text-emerald-600 dark:text-emerald-400">Copied!</span>
	{:else if isFetching}
		<svg class="size-3.5 animate-spin" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
			<path d="M21 12a9 9 0 1 1-6.219-8.56" />
		</svg>
		<span>Fetching</span>
	{:else}
		<Copy class="size-3.5 transition-all duration-200 group-hover:scale-110" />
		<span>Copy doc</span>
	{/if}
</Button>
