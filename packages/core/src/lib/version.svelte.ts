import { page } from "$app/state";
import { goto } from "$app/navigation";
import { getContext, setContext } from "svelte";

const VERSION_KEY = Symbol("version_manager");

class VersionManager {
	#versions = $state<string[]>([]);

	constructor(versions: string[]) {
		this.#versions = versions;
	}

	get all() {
		return this.#versions;
	}

	get current() {
		const segments = page.url.pathname.split("/");
		return segments[1] === "docs" && segments[2] ? segments[2] : "";
	}

	change(newVersion: string) {
		const segments = page.url.pathname.split("/");
		if (segments[1] === "docs" && segments[2]) {
			segments[2] = newVersion;
			goto(segments.join("/"));
		} else {
			goto(`/docs/${newVersion}`);
		}
	}
}

// Initialize and provide to the tree
export function setVersionManager(versions: string[]) {
	const manager = new VersionManager(versions);
	return setContext(VERSION_KEY, manager);
}

// Retrieve from anywhere in the tree
export function getVersionManager() {
	const manager = getContext<VersionManager>(VERSION_KEY);
	if (!manager) throw new Error("VersionManager not initialized");
	return manager;
}
