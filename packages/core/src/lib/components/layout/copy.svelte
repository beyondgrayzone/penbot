<script>
	import { page } from "$app/state";
	import Check from "lucide-svelte/icons/check";
	import CopyIcon from "lucide-svelte/icons/copy-check";
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

<Button class="cursor-pointer hover:bg-background bg-secondary text-primary active:bg-background dark:text-white dark:hover:text-black px-3 py-0" onclick={copyToClipboard} aria-label="Copy markdown">
	{#if copied}
		<Check />  Copied!
	{:else if isFetching}
		Fetching ..
	{:else}
		<CopyIcon /> Copy doc
	{/if}
</Button>
