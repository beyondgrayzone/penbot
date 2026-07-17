<script lang="ts">
	import { page } from "$app/state";
	import SidebarTrigger from "$lib/components/ui/sidebar/sidebar-trigger.svelte";
	import Search from "$lib/components/search/search.svelte";
	import SocialIcons from "../social-icons.svelte";
	import ThemeDropdown from "./theme-dropdown.svelte";
	import ThemeSwitcher from "./theme-switcher.svelte";

	let { version }: { version?: string } = $props();

	$effect(() => {
		console.log(`${JSON.stringify(page.url, null, 2)}`);
	});
</script>

<header
	class="bg-background-secondary sticky top-0 z-10 flex h-16 items-center gap-2 rounded-lg border-b px-4 md:top-0 dark:border-t"
>
	<SidebarTrigger class="mr-2 md:hidden [&>svg]:size-5" />
	<Search {version} />

	<div class="ml-auto flex items-center gap-8">
		<div class="flex items-center gap-3">
			<SocialIcons />
			<ThemeSwitcher />
			<ThemeDropdown />
		</div>
	</div>
</header>

