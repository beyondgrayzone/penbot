<!-- https://github.com/huntabyte/shadcn-svelte/blob/shadcn-svelte%401.0.0-next.7/sites/docs/src/lib/registry/new-york/block/sidebar-01/components/version-switcher.svelte -->
<script lang="ts">
	import { page } from "$app/state";
	import { onMount, tick } from "svelte";
	import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
	import * as Sidebar from "$lib/components/ui/sidebar/index.js";
	import Check from "lucide-svelte/icons/check";
	import ChevronsUpDown from "lucide-svelte/icons/chevrons-up-down";
	import GalleryVerticalEnd from "lucide-svelte/icons/gallery-vertical-end";
	import { useSiteConfig } from "$lib/hooks/use-site-config.svelte.js";
	import { getVersionManager } from "$lib/version.svelte.js";
	import { mode } from "mode-watcher";
	const siteConfig = useSiteConfig();
	const vm = getVersionManager();

	let { versions }: { versions: string[]; defaultVersion: string } = $props();
	let segments = $derived(page.url.pathname.split("/"));
	let newDefault = $derived(segments[2]);
	let selectedVersion = $derived(newDefault);

	let hasError = $state(false);

	let logoSrc = $state("/logo-small.png");

	$effect(() => {
		// const t = async () => {
		// 	await tick();
		// };
		// t();
		hasError = false;
		logoSrc = mode.current == "light" ? "/logo-small.png" : "/logo-small-dark.png";
	});

	// onMount(() => {
	// 	logoSrc = mode.current == "light" ? "/logo-small.png" : "/logo-small-dark.png";
	// });

	function handleVersionChange(newVersion: string) {
		const v = `v${newVersion.split(".")[0]}`;
		vm.change(v);
		selectedVersion = v;
	}
</script>

<Sidebar.Menu>
	<Sidebar.MenuItem>
		<DropdownMenu.Root>
			<DropdownMenu.Trigger>
				{#snippet child({ props })}
					<Sidebar.MenuButton
						size="lg"
						class="data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground"
						{...props}
					>
						<div
							class="bg-sidebar-primary text-sidebar-primary-foreground flex aspect-square size-8 items-center justify-center rounded-lg"
						>
							{#if !hasError}
								<img
									src={logoSrc}
									alt={""}
									onload={() => (hasError = false)}
									onerror={() => (hasError = true)}
								/>
							{:else}
								<GalleryVerticalEnd class="size-4" />
							{/if}
						</div>
						<div class="flex flex-col gap-0.5 leading-none">
							<span class="font-semibold">{siteConfig.current.name}</span>
							<span class="">{selectedVersion}</span>
						</div>
						<ChevronsUpDown class="ml-auto" />
					</Sidebar.MenuButton>
				{/snippet}
			</DropdownMenu.Trigger>
			<DropdownMenu.Content class="w-(--bits-dropdown-menu-anchor-width)" align="start">
				{#each versions as version (version)}
					<!-- 3. Update the onSelect handler -->
					<DropdownMenu.Item onSelect={() => handleVersionChange(version)}>
						{version}
						{#if version === selectedVersion}
							<Check class="ml-auto" />
						{/if}
					</DropdownMenu.Item>
				{/each}
			</DropdownMenu.Content>
		</DropdownMenu.Root>
	</Sidebar.MenuItem>
</Sidebar.Menu>
