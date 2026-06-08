import { getDocC } from "$lib/utils";

export async function load({ params }) {
	const version = "v2";
	const p = () => import.meta.glob("/src/mds/v2/**/*.md");
	const d = getDocC(params.slug, p, `src/mds/${version}`, `/docs/${version}/`);
	return { version: `${version}`, component: (await d).component, metadata: (await d).metadata };
}
