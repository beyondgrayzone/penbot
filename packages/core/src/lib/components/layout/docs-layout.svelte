<script lang="ts">
	// import { ModeWatcher } from "mode-watcher";
	import * as Sidebar from "$lib/components/ui/sidebar/index.js";
	import DocsSidebar from "$lib/components/layout/docs-sidebar.svelte";
	import type { Snippet } from "svelte";
	import type { Navigation } from "$lib/types.js";
	import Header from "./header.svelte";
	import Footer from "./footer.svelte";

	import versions from "./search.json" with { type: "json" };
	import { setVersionManager } from "$lib/version.svelte.js";

	setVersionManager(versions);

	let {
		children,
		navigation,
		version,
		theme,
	}: {
		children?: Snippet;
		navigation: Navigation;
		logo?: Snippet;
		theme?: string;
		version: string;
	} = $props();
</script>

<!-- <ModeWatcher defaultTheme={theme} /> -->

<Sidebar.Provider class="relative mx-auto lg:max-w-[1700px]">
	<DocsSidebar {navigation} {version} {versions} />
	<Sidebar.Inset
		class="h-svh overflow-hidden bg-[var(--theme-color-current-300)] pl-0 md:ml-[250px] dark:bg-[var(--theme-color-current-700)]"
	>
		<div
			class="bg-background flex flex-1 flex-col overflow-hidden border shadow-sm sm:m-2 sm:rounded-xl md:my-4"
		>
			<Header {version} />

			<div class="flex-1 overflow-y-auto">
				<div
					class="flex w-full flex-1 flex-row-reverse px-4 py-8 lg:pl-0 lg:pr-8 xl:gap-4"
					id="content"
				>
					{#if children}
						{@render children?.()}
					{/if}
				</div>
				<Footer />
			</div>
		</div>
	</Sidebar.Inset>
</Sidebar.Provider>
