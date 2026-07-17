<script lang="ts">
	import { dev } from "$app/environment";
	import { onMount } from "svelte";
	import { setTheme, theme } from "mode-watcher";

	const themes = [
		{ value: "oceanic", label: "Oceanic" },
		{ value: "forest", label: "Forest" },
		{ value: "sober-1", label: "Sober" },
		{ value: "amber", label: "Amber" },
		{ value: "blue", label: "Blue" },
		{ value: "cyan", label: "Cyan" },
		{ value: "emerald", label: "Emerald" },
		{ value: "fuchsia", label: "Fuchsia" },
		{ value: "green", label: "Green" },
		{ value: "indigo", label: "Indigo" },
		{ value: "lime", label: "Lime" },
		{ value: "orange", label: "Orange" },
		{ value: "pink", label: "Pink" },
		{ value: "purple", label: "Purple" },
		{ value: "red", label: "Red" },
		{ value: "rose", label: "Rose" },
		{ value: "sky", label: "Sky" },
		{ value: "teal", label: "Teal" },
		{ value: "violet", label: "Violet" },
		{ value: "yellow", label: "Yellow" },
	] as const;

	let current = $state(theme.current || "oceanic");

	// Stay in sync when theme changes from anywhere (setTheme, ModeWatcher init, etc.)
	$effect(() => {
		current = theme.current || "oceanic";
	});

	// Stay in sync when localStorage changes in another tab
	onMount(() => {
		function onStorage(e: StorageEvent) {
			if (e.key === "mode-watcher-theme" && e.newValue) {
				current = e.newValue;
			}
		}
		window.addEventListener("storage", onStorage);
		return () => window.removeEventListener("storage", onStorage);
	});

	function handleChange(e: Event) {
		const value = (e.target as HTMLSelectElement).value;
		setTheme(value);
	}
</script>

{#if dev}
	<div class="flex items-center gap-1.5">
		<select
			value={current}
			onchange={handleChange}
			class="bg-background text-foreground border-border hover:border-brand focus:border-brand cursor-pointer rounded-md border px-2 py-1 text-xs font-medium outline-none transition-colors"
			aria-label="Switch theme"
		>
			{#each themes as t}
				<option value={t.value}>{t.label}</option>
			{/each}
		</select>
	</div>
{/if}