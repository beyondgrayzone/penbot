<script lang="ts">
	import { page } from "$app/state";
	import { DocsLayout } from "@penbot/core";

	import { navigationWithVersion, navigation } from "$lib/navigation";
	let version = $derived(page.url.pathname.split("/")[2]);
	let n = $derived.by(() => {
		return navigationWithVersion(version);
	});

	let { children } = $props();
</script>

<DocsLayout navigation={n} {version}>
	{@render children?.()}
</DocsLayout>
