import { getDocC } from "$lib/utils";

export async function load({ params }) {
	const version = "v1";
	const p = () => import.meta.glob("/src/mds/v1/**/*.md");
	const d = getDocC(params.slug, p, `src/mds/${version}`, `/docs/${version}/`);
	// return getDocC(params.slug, p, "src/mds/v1", "/docs/v1/");
	return { version: `${version}`, component: (await d).component, metadata: (await d).metadata };
}
