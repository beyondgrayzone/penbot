<script lang="ts">
	import SocialIcons from "../social-icons.svelte";
	import { useSiteConfig } from "@bladocs/core";

	const siteConfig = useSiteConfig();
	let siteName = siteConfig.current.name;
	let siteURL = siteConfig.current.url;

	const year = `2024 - ${new Date().getFullYear()}`;
</script>

<footer class="bg-background-secondary flex h-16 items-center gap-2 rounded-lg border-t px-4">
	<span class="text-muted-foreground text-xs">
		© {year}
		<a href={siteURL} target="_blank" class="underline-offset-2 hover:underline">
			{siteName}
		</a>
	</span>
	<div class="ml-auto flex items-center">
		<SocialIcons />
	</div>
</footer>
