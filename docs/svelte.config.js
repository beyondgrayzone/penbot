import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";
import { mdsx } from "@penbot/mdsx";
import mdsxConfig from "./mdsx.config.js";
import adapter from "@sveltejs/adapter-static";

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: [mdsx(mdsxConfig), vitePreprocess()],
	kit: {
		prerender: {
			handleHttpError: "warn",
		},
		output: {
			bundleStrategy: "split",
		},
		// paths: {
		//   base: ''
		// },
		alias: {
			"$content/*": ".velite/*",
			"@/*": "src/*",
			// "@penbot/core/*": "../packages/core/src/lib/*",
		},
		adapter: adapter({
			pages: "build",
			assets: "build",
			fallback: "index.html",
			precompress: false,
			strict: true,
		}),
	},
	extensions: [".svelte", ".md"],
};

export default config;
