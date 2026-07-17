import { defineConfig } from "@penbot/mdsx";
import { baseRemarkPlugins, baseRehypePlugins } from "@penbot/core/mdsxConfig";
import { resolve } from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = fileURLToPath(new URL(".", import.meta.url));

export default defineConfig({
	remarkPlugins: [...baseRemarkPlugins],
	// @ts-expect-error shh
	rehypePlugins: [...baseRehypePlugins],
	blueprints: {
		default: {
			path: resolve(__dirname, "./src/lib/components/blueprint.svelte"),
		},
	},
	extensions: [".md"],
});
