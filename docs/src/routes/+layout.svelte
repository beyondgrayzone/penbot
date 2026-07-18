<script lang="ts">
	import { siteConfig } from "$lib/site-config";
	import "../app.css";
	import { untrack } from "svelte";
	import { useSiteConfig } from "@penbot/core";
	import { ModeWatcher, setTheme, theme } from "mode-watcher";
	import { onMount } from "svelte";

	let { children } = $props();

	useSiteConfig(() => siteConfig);

	/** localStorage key for persisting the client-side theme timestamp. */
	const THEME_TIMESTAMP_KEY = "penbot-theme-timestamp";

	/** The server's baked-in theme timestamp from the site config. */
	const serverThemeTimestamp = siteConfig.themeTimestamp ?? 0;

	/** The default theme to fall back to when overriding. */
	const defaultTheme = siteConfig.defaultTheme ?? "oceanic";

	/**
	 * Guards the $effect below so it only starts saving timestamps AFTER
	 * onMount has finished the initial server-vs-client comparison.
	 */
	let comparisonDone = $state(false);

	/**
	 * Set to `true` right before the server-forced setTheme() call so the
	 * $effect can skip saving Date.now() for that automatic change. It only
	 * saves Date.now() when the user manually switches themes.
	 */
	let pendingServerOverride = $state(false);

	onMount(() => {
		const raw = localStorage.getItem(THEME_TIMESTAMP_KEY);
		const clientTs = raw ? Number.parseInt(raw, 10) : 0;

		if (serverThemeTimestamp > clientTs) {
			// The server shipped a newer theme config → reset the user's theme
			pendingServerOverride = true;
			setTheme(defaultTheme);
		}

		// Stamp the server timestamp so we don't override again until the
		// next server redeploy bumps the value.
		localStorage.setItem(THEME_TIMESTAMP_KEY, String(serverThemeTimestamp));

		comparisonDone = true;
	});

	/**
	 * After the initial mount-comparison finishes, this effect reacts to every
	 * theme change. If the change came from a server override (first load) it
	 * skips saving. Otherwise the user manually changed the theme, so it saves
	 * Date.now() to mark when the user expressed a preference.
	 *
	 * On the next page load, if `serverThemeTimestamp > clientTs`, the theme
	 * will be overridden again; otherwise the user's preference is preserved.
	 */
	$effect(() => {
		// Track theme reactively so the effect re-runs on any theme switch
		void theme.current;

		if (!comparisonDone) return;

		// Use untrack so reading pendingServerOverride doesn't register it as a
		// dependency. Otherwise setting it below (pendingServerOverride = false)
		// triggers a re-run, creating a self-triggering cascade.
		const isServerOverride = untrack(() => pendingServerOverride);

		if (isServerOverride) {
			// This change was triggered by the server override on mount  do not
			// save Date.now() because the user didn't actually interact.
			pendingServerOverride = false;
			return;
		}

		// User changed the theme  save the current time as the client timestamp
		localStorage.setItem(THEME_TIMESTAMP_KEY, String(Date.now()));
	});
</script>

<svelte:head>
	<meta name="penbot:theme-timestamp" content={serverThemeTimestamp} />
</svelte:head>

<ModeWatcher {defaultTheme} />
<div class="bg-[var(--theme-color-current-300)] dark:bg-[var(--theme-color-current-700)]">
	{@render children?.()}
</div>
