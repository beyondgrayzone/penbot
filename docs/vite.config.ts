import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";
import tailwindcss from "@tailwindcss/vite";
import { resolve } from "node:path";
import devToolsJson from "vite-plugin-devtools-json";

const __dirname = new URL(".", import.meta.url).pathname;

export default defineConfig({
	plugins: [tailwindcss(), sveltekit(), devToolsJson()],
	optimizeDeps: {
		exclude: ["@bladocs/core"],
	},
	ssr: {
		noExternal: ["@bladocs/core"],
	},
	server: {
		host: process.env.VITE_HOST || undefined,
		port: process.env.VITE_PORT ? parseInt(process.env.VITE_PORT) : undefined,
		fs: {
			allow: [resolve(__dirname, "./.velite"), resolve(__dirname, "../packages")],
		},
	},
});
