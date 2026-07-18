import { getContext, hasContext, setContext } from "svelte";

export type SiteConfig = {
	name: string;
	url: string;
	siteLink: string;
	description: string;
	links?: {
		logo?: string;
		x?: string;
		github?: string;
	};
	author?: string;
	keywords?: string[];
	ogImage?: {
		url: string;
		width: string;
		height: string;
	};
	license?: {
		name: string;
		url: string;
	};
	/**
	 * A Unix timestamp (in milliseconds) representing when the site's theme
	 * configuration was last updated. When the server renders a page, this
	 * timestamp is embedded in the HTML. On the client, it is compared against
	 * a timestamp saved in localStorage. If the server timestamp is greater,
	 * the client's theme is overridden with the server's default.
	 *
	 * Update this value (e.g. at build/deploy time) whenever you change the
	 * default theme so returning visitors pick up the new theme.
	 */
	themeTimestamp?: number;

	/**
	 * The default theme to use for first-time visitors.
	 * Must match one of the theme CSS files imported in app.css, e.g. "oceanic", "forest".
	 */
	defaultTheme?: string;
};

export function createSiteConfig(config: SiteConfigState) {
	return config;
}

class SiteConfigState {
	current = $derived.by(() => this.getProps());
	constructor(readonly getProps: () => SiteConfig) {}
}

const SITE_CONFIG_KEY = Symbol("penbot-site-config");

export function useSiteConfig(getProps?: () => SiteConfig): SiteConfigState {
	if (getProps) {
		return setContext(SITE_CONFIG_KEY, new SiteConfigState(getProps));
	} else if (hasContext(SITE_CONFIG_KEY)) {
		return getContext(SITE_CONFIG_KEY);
	} else {
		throw new Error(
			"useSiteConfig must be called with a function that returns a SiteConfigProps object",
		);
	}
}
