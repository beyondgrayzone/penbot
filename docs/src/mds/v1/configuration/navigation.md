---
title: Navigation
description: Learn how to customize the navigation in your Penbot project.
section: Configuration
---

Navigation is a key component of every site, documenting the structure of your site and providing a clear path for users to navigate through your content.

Penbot comes with a navigation structure that is designed to be flexible and customizable. Each page in your site should have a corresponding navigation item, and the navigation items should be nested according to their hierarchy.

## Navigation Types

The navigation configuration supports four types of items, defined in the `Navigation` type:

```ts
type Navigation = {
	anchors?: AnchorNavItem[];
	header?: HeaderNavItem[];
	sections?: SidebarNavSection[];
	items?: SidebarNavItem[];
};
```

### Anchors

Anchors are links displayed at the top of the sidebar. They typically highlight important pages or provide quick access to key content. Each anchor can include an icon component:

```ts
type AnchorNavItem = {
	title: string;
	href: string;
	icon: Component; // A Svelte component (e.g., from Phosphor Icons)
	disabled?: boolean;
};
```

### Header

Header items are links displayed in the top navigation bar. They provide quick access to external links or top-level pages:

```ts
type HeaderNavItem = {
	title: string;
	href: string;
};
```

### Sections

Sections group related navigation items under a category label in the sidebar:

```ts
type SidebarNavSection = {
	title?: string; // Section label (rendered as a group header)
	items: SidebarNavItem[];
};
```

### Items

A flat list of sidebar items (rendered without a section header):

```ts
type SidebarNavItem = {
	title: string;
	href?: string;
	disabled?: boolean;
	external?: boolean;
	label?: string;
};
```

## Configuration Files

Navigation is split across two files in your project:

### 1. Static Structure (`navigation.json`)

The `docs/src/lib/navigation.json` file defines the static structure  which sections and anchors exist. This is decoupled from the dynamic population of docs pages.

```json
{
	"anchors": [
		{
			"title": "Introduction",
			"href": "/",
			"icon": "ChalkboardTeacher"
		},
		{
			"title": "Getting Started",
			"href": "/getting-started",
			"icon": "RocketLaunch"
		}
	],
	"sections": [
		{
			"title": "Configuration"
		},
		{
			"title": "Components"
		}
	]
}
```

The `icon` field references a component name (imported via an icon map in `navigation.ts`).

### 2. Dynamic Builder (`navigation.ts`)

The `docs/src/lib/navigation.ts` file imports the static config and dynamically populates section items from your Velite-processed docs:

```ts
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
		href: version?.length
			? `${baseURL}/${version}${anchor.href}`
			: `${baseURL}${anchor.href}`,
		icon: iconMap[anchor.icon],
	}));

const dynamicSections = (version: string) =>
	navConfig.sections.map((section) => {
		const items = allDocs
			.filter((doc) => doc.section === section.title)
			.filter((doc) => doc.slug.indexOf(version) > -1)
			.map((doc) => ({
				title: doc.title,
				href: version?.length
					? `${baseURL}/${doc.slug}`
					: `${baseURL}/${doc.slug}`,
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
```

Key points:

- **`getAllDocs()`** returns all Velite-processed documentation entries with metadata (`title`, `slug`, `section`, etc.)
- **`dynamicAnchors()`** builds anchor links with the correct version prefix and resolves icon components
- **`dynamicSections()`** populates each section with docs that match that section's title
- **`defineNavigation()`** is a typed utility function (returns the object as-is for type safety)
- **`navigationWithVersion()`** allows creating versioned navigation objects dynamically

## Header Navigation

The navigation object also supports a `header` property for links in the top navigation bar:

```ts
export const navigation = defineNavigation({
	anchors: [...],
	header: [
		{ title: "GitHub", href: "https://github.com/beyondgrayzone/penbot" },
		{ title: "API", href: "/api" },
	],
	sections: [...],
});
```

## Version-Aware Navigation

The Penbot docs starter supports multiple documentation versions. The `navigationWithVersion` helper filters docs by version slug (`v1`, `v2`, etc.) so the sidebar only shows pages relevant to the selected version.

The version is derived from the URL path (e.g., `/docs/v1/getting-started`). When a user switches versions via the version dropdown, the navigation updates accordingly.

## DefineNavigation Utility

The `defineNavigation` function is exported from `@penbot/core` and provides type-safe navigation configuration:

```ts
import { defineNavigation } from "@penbot/core";

export const navigation = defineNavigation({
	anchors: [
		{
			title: "Introduction",
			href: "/",
			icon: MyIcon,
		},
	],
	header: [
		{ title: "GitHub", href: "https://github.com/beyondgrayzone/penbot" },
	],
	sections: [
		{
			title: "Guides",
			items: [
				{ title: "Getting Started", href: "/docs/v1/getting-started" },
				{ title: "Configuration", href: "/docs/v1/configuration" },
			],
		},
	],
	items: [
		{ title: "Changelog", href: "/changelog" },
	],
});
```
