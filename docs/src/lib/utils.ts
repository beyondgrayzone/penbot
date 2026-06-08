import { docs, type Doc } from "$content/index.js";
import { error } from "@sveltejs/kit";
import type { Component } from "svelte";
import type { ImportGlobFunction } from "vite/types/importGlob.js";

import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs));
}

export function getDocMetadataC(slug: string = "index", slugPrefix: string) {
	const splt = slugPrefix.split("/");
	// docs.find((doc) => {
	// 	if (doc.slug == "v2") {

	// 		console.log(slug ,slugPrefix, JSON.stringify(doc, null, 2))

	// 		// throw Error(splt[2])
	// 	}
	//  });
	return docs.find(
		(doc) => doc.slugFull === `${slugPrefix}${slug}` || (slug === "index" && doc.path === splt[2]),
	);
	// return docs.find((doc) => doc.slugFull === `${slugPrefix}${slug}`);
}

export function getDocMetadata(slug: string = "index") {
	return docs.find((doc) => doc.slugFull === `${slug}`);
}

export function getAllDocs() {
	return docs;
}

function slugFromPath(path: string, pp: string) {
	return path.replace(`/${pp}/`, "").replace(".md", "");
}

export type DocResolver = () => Promise<{ default: Component; metadata: Doc }>;

export async function getDoc(slug: string = "index", p: () => ImportGlobFunction, pp: string) {
	const modules = p();
	// import.meta.glob("/src/mds/**/*.md");

	let match: { path?: string; resolver?: DocResolver } = {};

	for (const [path, resolver] of Object.entries(modules)) {
		if (slugFromPath(path, pp) === slug) {
			match = { path, resolver: resolver as unknown as DocResolver };
			break;
		}
	}
	if (!match) {
		throw Error(slug);
	}
	const doc = await match?.resolver?.();
	const metadata = getDocMetadata(slug);
	if (!doc || !metadata) {
		error(404, "Could not find the document.");
	}

	return {
		component: doc.default,
		metadata,
	};
}

export async function getDocC(
	slug: string = "index",
	p: () => ImportGlobFunction,
	pp: string,
	slugPrefix: string,
) {
	const modules = p();

	let match: { path?: string; resolver?: DocResolver } = {};

	for (const [path, resolver] of Object.entries(modules)) {
		if (slugFromPath(path, pp) === slug) {
			match = { path, resolver: resolver as unknown as DocResolver };
			break;
		}
	}
	if (!match) {
		throw Error(slug);
	}
	const doc = await match?.resolver?.();

	const metadata = getDocMetadataC(slug, slugPrefix);
	if (!doc || !metadata) {
		error(404, "Could not find the document.");
	}

	return {
		component: doc.default,
		metadata,
	};
}
