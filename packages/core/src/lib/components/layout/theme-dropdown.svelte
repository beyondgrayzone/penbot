<script lang="ts">
	import { onMount } from "svelte";
	import { setMode, mode } from "mode-watcher";
	import { Switch } from "$lib/components/ui/switch/index.js";
	import MoonStars from "phosphor-svelte/lib/MoonStars";
	import Sun from "phosphor-svelte/lib/CloudSun";

	let checkedV = $state(false);

	const modes = ["light", "dark", "system"] as const;
	const onCheckedChange = (checked: boolean) => {
		// alert(checkedV);
		if (checked) {
			setMode(modes[1]);
		} else {
			setMode(modes[0]);
		}
		checkedV = checked;
	};

	onMount(() => {
		checkedV = mode.current == "dark"
	})
</script>

<div class="flex gap-2">
	{#if checkedV}
		<Sun class="size-6" />
	{/if}
	<Switch checked={checkedV}  id="airplane-mode" {onCheckedChange} />

	{#if !checkedV}
		<MoonStars class="size-6" />
	{/if}
</div>
