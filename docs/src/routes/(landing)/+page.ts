import { redirect } from "@sveltejs/kit";

const baseURL = "/docs";

export function load() {
	redirect(302, baseURL);
}
