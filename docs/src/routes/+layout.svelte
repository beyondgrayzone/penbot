<script lang="ts">
	import { siteConfig } from "$lib/site-config";
	import "../app.css";
	import { useSiteConfig } from "@penbot/core";
	import { ModeWatcher } from "mode-watcher";

	let { children } = $props();

	useSiteConfig(() => siteConfig);
</script>

<ModeWatcher defaultTheme="oceanic" />
<div class="bg-[var(--theme-color-current-300)] dark:bg-[var(--theme-color-current-700)]">
	{@render children?.()}
</div>
