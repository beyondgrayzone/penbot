---
title: Navigation
description: Learn how to customize the navigation in your Penbot project.
section: Configuration
---

Navigation is a key component of your documentation, providing a clear path for users to browse content. Penbot's navigation is configured across two files: a static structure file and a dynamic builder.

## Static Structure (`navigation.json`)

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

This file is consumed in two ways:

- **At runtime**  `navigation.ts` imports it directly (`import navConfig from "./navigation.json"`) and uses its `anchors` and `sections` arrays as the structural skeleton for building the full versioned navigation object.
- **Via the CLI**  the `mds sections` command reads and writes this file to keep it in sync with `velite.config.js` when sections are added or removed.

## Managing Sections via CLI

The `mds sections` command manages documentation sections in both `velite.config.js` and `navigation.json` simultaneously:

```bash
bun run docs:sections                    # list all sections
bun run docs:sections add <name>         # add a new section
bun run docs:sections remove <name>      # remove an existing section
```

When you add or remove a section, the CLI updates both files automatically, keeping the Velite schema and the navigation structure in sync.

## Dynamic Builder (`navigation.ts`)

The `docs/src/lib/navigation.ts` file imports the static config and populates section items automatically from your Velite-processed docs. It also resolves icon strings to actual Svelte components and prefixes links with the correct version path.

You typically don't need to touch this file  it handles everything automatically. If you do need to customize it, the file is located at `docs/src/lib/navigation.ts`.

## Version-Aware Navigation

Penbot supports multiple documentation versions out of the box. The navigation automatically filters docs by version (`v1`, `v2`, etc.), so the sidebar only shows pages relevant to the selected version. When a user switches versions via the version dropdown, the navigation updates accordingly.
