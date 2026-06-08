import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		alias: {
			"@/*": "./*",
			// "@bladocs/core/*": "../packages/core/src/lib/*",
		},
	},
	preprocess: vitePreprocess(),
};

export default config;
