// type imports
import { type Component } from "svelte";

// external imports
import { defineNavigation } from "@penbot/core";
import ChalkboardTeacher from "phosphor-svelte/lib/ChalkboardTeacher";
import RocketLaunch from "phosphor-svelte/lib/RocketLaunch";

// relative imports
import { getAllDocs } from "./utils.js";

// asset imports
import navConfig from "./navigation.json";

const iconMap: Record<string, Component> = {
	ChalkboardTeacher,
	RocketLaunch,
};

const allDocs = getAllDocs();
const baseURL = "/docs";

const dynamicAnchors = (version: string) =>
	navConfig.anchors.map((anchor) => ({
		...anchor,
		href: version?.length ? `${baseURL}/${version}${anchor.href}` : `${baseURL}${anchor.href}`,
		icon: iconMap[anchor.icon],
	}));

const dynamicSections = (version: string) =>
	navConfig.sections.map((section) => {
		const items = allDocs
			.filter((doc) => doc.section === section.title)
			.filter((doc) => doc.slug.indexOf(version) > -1)
			.map((doc) => ({
				title: doc.title,
				href: version?.length ? `${baseURL}/${doc.slug}` : `${baseURL}/${doc.slug}`,
			}));

		return {
			title: section.title,
			items: items,
		};
	});

export const navigation = defineNavigation({
	anchors: dynamicAnchors("v1"),
	sections: dynamicSections("v1"),
});

export const navigationWithVersion = (v: string) => {
	return defineNavigation({
		anchors: dynamicAnchors(v),
		sections: dynamicSections(v),
	});
};
