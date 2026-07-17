import { Processed, PreprocessorGroup } from 'svelte/compiler';

import { PluggableList } from 'unified';


type MDSXConfig = {

    /** Remark plugins that apply to all documents */

    remarkPlugins?: PluggableList;

    /** Rehype plugins that apply to all documents */

    rehypePlugins?: PluggableList;

    extensions?: string[];

    blueprints?: BlueprintMap;

    /**

     * Path to a svelte config file, either absolute or relative to `process.cwd()`.

     *

     * Set to `false` to ignore the svelte config file.

     */

    svelteConfigPath?: string | false;

    frontmatterParser?: (str: string) => Record<string, unknown> | void;

};

type BlueprintMap = Record<string, Blueprint> & {

    default: Blueprint;

};

type Blueprint = {

    /** Path to the blueprint */

    path: string;

    /** Remark plugins that only apply to this blueprint */

    remarkPlugins?: PluggableList;

    /** Rehype plugins that only apply to this blueprint */

    rehypePlugins?: PluggableList;

};

declare function defineConfig(config: MDSXConfig): MDSXConfig;


type CompileOptions = {

    config?: MDSXConfig;

    filename?: string;

    preprocessors: PreprocessorGroup[];

};

declare function compile(source: string, { config, filename, preprocessors }: CompileOptions): Promise<Processed>;


declare function mdsx(config?: MDSXConfig): PreprocessorGroup;


export { type Blueprint, type BlueprintMap, type MDSXConfig, compile, defineConfig, mdsx };