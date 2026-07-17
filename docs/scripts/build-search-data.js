import { fileURLToPath } from "node:url";
import { writeFileSync, mkdirSync } from "node:fs";
import { resolve, join } from "node:path";
import { docs } from "../.velite/index.js";
import { defineSearchContent, cleanMarkdown } from "@penbot/core/search";

const __dirname = fileURLToPath(new URL(".", import.meta.url));
const API_BASE_DIR = resolve(__dirname, "../src/routes/api");
const KIT_DIR = resolve(__dirname, "../../packages/core/src/lib/components/layout");

const versions = [...new Set(docs.map((doc) => doc.slug.split("/")[0]))];
const currentVersions = [];
const versionsActual = () => {
	docs.map((doc) => {
		if (
			doc.section == "Overview" &&
			doc.currentVersion &&
			!currentVersions.includes(doc.currentVersion)
		) {
			currentVersions.push(doc.currentVersion);
		}
	});

	if (currentVersions.length !== versions.length) {
		console.log(versions, currentVersions);
		throw Error("You seem to have more version than you need!");
	}
};

// const addVersion = (v) => {
// 	if (!versions.includes(v)) {
// 		versions.push(v)
// 	}
// }

// export function buildDocsSearchIndex(baseURL = "/") {
// 	docs.map((doc) => {
// 		const spl = doc.slug.split("/")

// 		addVersion(spl[0])
// 	})
// 	return defineSearchContent(
// 		docs.map((doc) => ({
// 			title: doc.title,
// 			href: `${baseURL}/${doc.slug}`,
// 			description: doc.description,
// 			content: cleanMarkdown(doc.raw),
// 			category: doc.section,
// 		})),
// 	);
// }

const baseURL = "/docs";

// writeFileSync(
// 	resolve(__dirname, "../src/routes/api/search.json/search.json"),
// 	JSON.stringify(buildDocsSearchIndex(baseURL)),
// 	{ flag: "w" },
// );

function generateSearchData(filteredDocs) {
	return defineSearchContent(
		filteredDocs.map((doc) => ({
			title: doc.title,
			href: `${baseURL}/${doc.slug}`,
			description: doc.description,
			content: cleanMarkdown(doc.raw),
			category: doc.section,
		})),
	);
}

versions.forEach((version) => {
	versionsActual();
	// Create the directory: e.g., src/routes/api/v1
	const versionDir = join(API_BASE_DIR, `${version}.search.json`);
	mkdirSync(versionDir, { recursive: true });

	// Filter docs belonging to this version
	const versionDocs = docs.filter((doc) => doc.slug.startsWith(`${version}/`));
	const searchIndex = generateSearchData(versionDocs);

	// Write the file: e.g., src/routes/api/v1/search.json
	writeFileSync(join(versionDir, "search.json"), JSON.stringify(searchIndex), { flag: "w" });
	writeFileSync(join(KIT_DIR, `search.json`), JSON.stringify(currentVersions), { flag: "w" });

	console.log(`✅ Generated: api/${version}.search.json`);
});
