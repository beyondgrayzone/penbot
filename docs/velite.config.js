import { defineConfig, s, z } from "velite";
const sections = ["Overview", "Components", "Configuration", "Utilities"];

const baseSchema = s.object({
	title: s.string().optional(),
	description: s.string().optional(),
	path: s.path(),
	availableSinceVersion: s.string().optional(),
	currentVersion: s.string().optional(),
	content: s.markdown(),
	navLabel: s.string().optional(),
	raw: s.raw(),
	toc: s.toc(),
	order: z
		.number()
		.nullable()
		.transform((v) => v ?? 0)
		.default(0),
	section: s.enum(sections),
});

const docSchema = baseSchema.transform((data) => {
	return {
		...data,
		slug: data.path,
		slugFull: `/docs/${data.path}`,
	};
});

export default defineConfig({
	root: "./src/mds",
	collections: {
		docs: {
			name: "Doc",
			pattern: "**/*.md",
			schema: docSchema,
		},
	},
});
