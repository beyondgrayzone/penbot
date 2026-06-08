// import { redirect } from "@sveltejs/kit";

// export function load() {
// 		redirect(302, "/docs/v2/getting-started");
// }

import { getDocC } from "$lib/utils";

export async function load({ params }) {
	const version = "v2";
	const p = () => import.meta.glob("/src/mds/v2/**/*.md");

	const d = getDocC("index", p, `src/mds/${version}`, `/docs/${version}/`);
	return { version: `${version}`, component: (await d).component, metadata: (await d).metadata };
}
