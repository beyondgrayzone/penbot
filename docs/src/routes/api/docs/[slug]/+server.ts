import { error } from "@sveltejs/kit";
import { existsSync, readFileSync } from "fs";
import { runProcessMds } from "./process";

export const prerender = true;

const modules = import.meta.glob("@/mds/**/*.md", {
	query: "?raw",
	import: "default",
});

export function entries() {
	return Object.keys(modules).map((path) => {
		const slug = path.split("/src/mds/")[1].split("/").join("_").replace(".md", "");
		return { slug };
	});
}

export async function GET({ params }) {
	const splits = params.slug.split("_");
	const midPathArr: string[] = [];
	splits.forEach((item) => {
		midPathArr.push(item);
	});
	const midPath = midPathArr.join("/");
	const filePath = `/src/mds/${midPath}.md`;
	const finalFilePath = `/src/mds_final/${midPath}.md`;

	const WORKSPACE = "@bladocs/mds";
	const INPUT_FILE = filePath;
	const OUTPUT_FILE = finalFilePath;

	if (!modules[filePath]) {
		error(404, "Not found");
	}

	const finalPath = await runProcessMds(WORKSPACE, "processMds", INPUT_FILE, OUTPUT_FILE);

	if (!existsSync(finalPath)) {
		error(500, "File processing failed");
	}

	const content = readFileSync(finalPath, "utf-8");
	return new Response(content, {
		headers: { "Content-Type": "text/plain" },
	});
}
