import { redirect } from "@sveltejs/kit";

export function load() {
	const modules = import.meta.glob("@/routes/api/**/*.json");

	let maxVersion = 1;

	for (const path in modules) {
		const match = path.match(/v(\d+)\.search\.json/);

		if (match && match[1]) {
			const version = parseInt(match[1], 10);
			if (version > maxVersion) {
				maxVersion = version;
			}
		}
	}

	redirect(302, `/docs/v${maxVersion}`);
}
