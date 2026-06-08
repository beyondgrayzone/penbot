package mds

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------------
// Version Subcommand: add / remove / sync versioned route folders
// ---------------------------------------------------------------------------

// findProjectRoot walks up from cwd until it finds a directory containing
// both a "docs" folder and a "package.json" file (the monorepo root).
func findProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("getwd: %w", err)
	}

	dir := cwd
	for {
		docsDir := filepath.Join(dir, "docs")
		pkgFile := filepath.Join(dir, "package.json")

		if info, err := os.Stat(docsDir); err == nil && info.IsDir() {
			if _, err := os.Stat(pkgFile); err == nil {
				return dir, nil
			}
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find project root (no docs/ + package.json found) from %s", cwd)
		}
		dir = parent
	}
}

// validVersionName checks that the version string is valid like "v1", "v2", "v3", etc.
var rxVersion = regexp.MustCompile(`^v\d+$`)

// ---------------------------------------------------------------------------
// Shared Templates
// ---------------------------------------------------------------------------

const pageSvelteTmpl = `<script lang="ts">
	import { DocPage } from "@bladocs/core";
	let { data } = $props();
</script>

<DocPage component={data.component} {...data.metadata} />
`

const searchServerTmpl = `import type { RequestHandler } from "@sveltejs/kit";
import search from "./search.json" with { type: "json" };

export const prerender = true;

export const GET: RequestHandler = () => {
	return Response.json(search);
};
`

func pageTsContent(version string) string {
	return fmt.Sprintf(`import { getDocC } from "$lib/utils";

export async function load({ params }) {
	const version = "%[1]s";
	const p = () => import.meta.glob("/src/mds/%[1]s/**/*.md");

	const d = getDocC("index", p, "src/mds/"+version, "/docs/"+version+"/");
	return { version: version, component: (await d).component, metadata: (await d).metadata };
}
`, version)
}

func slugTsContent(version string) string {
	return fmt.Sprintf(`import { getDocC } from "$lib/utils";

export async function load({ params }) {
	const version = "%[1]s";
	const p = () => import.meta.glob("/src/mds/%[1]s/**/*.md");
	const d = getDocC(params.slug, p, "src/mds/"+version, "/docs/"+version+"/");
	return { version: version, component: (await d).component, metadata: (await d).metadata };
}
`, version)
}

func indexMdContent(version string) string {
	num := versionNum(version)
	semver := fmt.Sprintf("%d.0.0", num)
	return fmt.Sprintf(`---
title: Introduction
description: Getting started with %s
section: Overview
currentVersion: %s
---
`, version, semver)
}

// ---------------------------------------------------------------------------
// Helpers: ensure each file / directory
// ---------------------------------------------------------------------------

func ensureMdsDir(root, version string) (created bool, err error) {
	mdsDir := filepath.Join(root, "docs", "src", "mds", version)
	mdsIndex := filepath.Join(mdsDir, "index.md")

	if _, err := os.Stat(mdsDir); os.IsNotExist(err) {
		if err := os.MkdirAll(mdsDir, 0755); err != nil {
			return false, fmt.Errorf("mkdir %s: %w", mdsDir, err)
		}
		fmt.Printf("  ✔ Created %s\n", relPath(root, mdsDir))

		content := indexMdContent(version)
		if err := os.WriteFile(mdsIndex, []byte(content), 0644); err != nil {
			return false, fmt.Errorf("write %s: %w", mdsIndex, err)
		}
		fmt.Printf("  ✔ Created %s\n", relPath(root, mdsIndex))
		return true, nil
	}

	// Directory exists but index.md may be missing
	if _, err := os.Stat(mdsIndex); os.IsNotExist(err) {
		content := indexMdContent(version)
		if err := os.WriteFile(mdsIndex, []byte(content), 0644); err != nil {
			return false, fmt.Errorf("write %s: %w", mdsIndex, err)
		}
		fmt.Printf("  ✔ Created %s\n", relPath(root, mdsIndex))
		return true, nil
	}

	return false, nil
}

// copyDir recursively copies a directory tree from src to dst.
func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)

		if info.IsDir() {
			return os.MkdirAll(target, info.Mode())
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, info.Mode())
	})
}

// copyMdsFromPrevious copies the previous version's entire mds directory into the new version,
// then updates only the currentVersion in the new index.md to match the new version.
func copyMdsFromPrevious(root, version string) error {
	num := versionNum(version)
	if num <= 1 {
		return fmt.Errorf("no previous version to copy from for %q", version)
	}
	prevVersion := fmt.Sprintf("v%d", num-1)

	srcDir := filepath.Join(root, "docs", "src", "mds", prevVersion)
	dstDir := filepath.Join(root, "docs", "src", "mds", version)

	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		return fmt.Errorf("previous version %q not found at %s", prevVersion, srcDir)
	}

	if err := copyDir(srcDir, dstDir); err != nil {
		return fmt.Errorf("copy %s -> %s: %w", srcDir, dstDir, err)
	}
	fmt.Printf("  ✔ Copied %s -> %s\n", relPath(root, srcDir), relPath(root, dstDir))

	// Rewrite index.md with the new currentVersion only
	newIndexPath := filepath.Join(dstDir, "index.md")
	semver := versionToSemver(version)

	// Read the copied index.md
	data, err := os.ReadFile(newIndexPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", newIndexPath, err)
	}

	// Replace the currentVersion line with the new semver
	rx := regexp.MustCompile(`(?m)^currentVersion:\s*\S+`)
	newContent := rx.ReplaceAllString(string(data), fmt.Sprintf("currentVersion: %s", semver))

	if err := os.WriteFile(newIndexPath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("write %s: %w", newIndexPath, err)
	}
	fmt.Printf("  ✔ Updated currentVersion in %s\n", relPath(root, newIndexPath))

	return nil
}

func ensureSearchJsonDir(root, version string) (created bool, err error) {
	searchDir := filepath.Join(root, "docs", "src", "routes", "api", version+".search.json")
	serverPath := filepath.Join(searchDir, "+server.ts")
	jsonPath := filepath.Join(searchDir, "search.json")

	if _, err := os.Stat(searchDir); os.IsNotExist(err) {
		if err := os.MkdirAll(searchDir, 0755); err != nil {
			return false, fmt.Errorf("mkdir %s: %w", searchDir, err)
		}
		fmt.Printf("  ✔ Created %s\n", relPath(root, searchDir))
	}

	anyCreated := false

	if _, err := os.Stat(serverPath); os.IsNotExist(err) {
		if err := os.WriteFile(serverPath, []byte(searchServerTmpl), 0644); err != nil {
			return false, fmt.Errorf("write %s: %w", serverPath, err)
		}
		fmt.Printf("  ✔ Created %s\n", relPath(root, serverPath))
		anyCreated = true
	}

	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		if err := os.WriteFile(jsonPath, []byte("[]\n"), 0644); err != nil {
			return false, fmt.Errorf("write %s: %w", jsonPath, err)
		}
		fmt.Printf("  ✔ Created %s\n", relPath(root, jsonPath))
		anyCreated = true
	}

	return anyCreated, nil
}

func ensureRouteDir(root, version string) (created bool, err error) {
	slugDir := filepath.Join(root, "docs", "src", "routes", "(docs)", "docs", version)
	slugCatchall := filepath.Join(slugDir, "[...slug]")

	if _, err := os.Stat(slugDir); os.IsNotExist(err) {
		if err := os.MkdirAll(slugCatchall, 0755); err != nil {
			return false, fmt.Errorf("mkdir %s: %w", slugCatchall, err)
		}
		fmt.Printf("  ✔ Created %s\n", relPath(root, slugDir))
	}

	anyCreated := false

	// +page.svelte
	path := filepath.Join(slugDir, "+page.svelte")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.WriteFile(path, []byte(pageSvelteTmpl), 0644); err != nil {
			return false, fmt.Errorf("write %s: %w", path, err)
		}
		fmt.Printf("  ✔ Created %s\n", relPath(root, path))
		anyCreated = true
	}

	// +page.ts
	path = filepath.Join(slugDir, "+page.ts")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.WriteFile(path, []byte(pageTsContent(version)), 0644); err != nil {
			return false, fmt.Errorf("write %s: %w", path, err)
		}
		fmt.Printf("  ✔ Created %s\n", relPath(root, path))
		anyCreated = true
	}

	// [...slug]/+page.svelte
	path = filepath.Join(slugCatchall, "+page.svelte")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return false, fmt.Errorf("mkdir: %w", err)
		}
		if err := os.WriteFile(path, []byte(pageSvelteTmpl), 0644); err != nil {
			return false, fmt.Errorf("write %s: %w", path, err)
		}
		fmt.Printf("  ✔ Created %s\n", relPath(root, path))
		anyCreated = true
	}

	// [...slug]/+page.ts
	path = filepath.Join(slugCatchall, "+page.ts")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.WriteFile(path, []byte(slugTsContent(version)), 0644); err != nil {
			return false, fmt.Errorf("write %s: %w", path, err)
		}
		fmt.Printf("  ✔ Created %s\n", relPath(root, path))
		anyCreated = true
	}

	return anyCreated, nil
}

// discoverVersions returns all version folder names found in the routes dir.
func discoverVersions(root string) ([]string, error) {
	routesDir := filepath.Join(root, "docs", "src", "routes", "(docs)", "docs")
	entries, err := os.ReadDir(routesDir)
	if err != nil {
		return nil, fmt.Errorf("read routes dir %s: %w", routesDir, err)
	}
	var versions []string
	for _, e := range entries {
		if e.IsDir() && rxVersion.MatchString(e.Name()) {
			versions = append(versions, e.Name())
		}
	}
	sort.Strings(versions)
	return versions, nil
}

// discoverMdsVersions returns all version folder names found in the mds dir.
func discoverMdsVersions(root string) ([]string, error) {
	mdsDir := filepath.Join(root, "docs", "src", "mds")
	entries, err := os.ReadDir(mdsDir)
	if err != nil {
		return nil, fmt.Errorf("read mds dir %s: %w", mdsDir, err)
	}
	var versions []string
	for _, e := range entries {
		if e.IsDir() && rxVersion.MatchString(e.Name()) {
			versions = append(versions, e.Name())
		}
	}
	sort.Strings(versions)
	return versions, nil
}

// discoverSearchVersions returns all version names found in the api search json dir.
func discoverSearchVersions(root string) ([]string, error) {
	apiDir := filepath.Join(root, "docs", "src", "routes", "api")
	entries, err := os.ReadDir(apiDir)
	if err != nil {
		return nil, fmt.Errorf("read api dir %s: %w", apiDir, err)
	}
	var versions []string
	rxSearch := regexp.MustCompile(`^(v\d+)\.search\.json$`)
	for _, e := range entries {
		if m := rxSearch.FindStringSubmatch(e.Name()); len(m) > 1 {
			versions = append(versions, m[1])
		}
	}
	sort.Strings(versions)
	return versions, nil
}

// versionSet returns a set-like map for quick lookup.
func versionSet(versions []string) map[string]bool {
	s := make(map[string]bool, len(versions))
	for _, v := range versions {
		s[v] = true
	}
	return s
}

// ---------------------------------------------------------------------------
// Add Version
// ---------------------------------------------------------------------------

func CmdVersionAdd(version string, copyPrev bool) error {
	if !rxVersion.MatchString(version) {
		return fmt.Errorf("invalid version name %q – must match v<number>, e.g. v3", version)
	}

	root, err := findProjectRoot()
	if err != nil {
		return err
	}

	// Check if version already exists
	docDir := filepath.Join(root, "docs", "src", "routes", "(docs)", "docs", version)
	if _, err := os.Stat(docDir); err == nil {
		return fmt.Errorf("version %q already exists at %s", version, docDir)
	}

	if _, err := ensureRouteDir(root, version); err != nil {
		return err
	}
	if _, err := ensureSearchJsonDir(root, version); err != nil {
		return err
	}
	if copyPrev {
		if err := copyMdsFromPrevious(root, version); err != nil {
			return err
		}
	} else {
		if _, err := ensureMdsDir(root, version); err != nil {
			return err
		}
	}

	// Update core search.json with the new version
	routeVersions, err := discoverVersions(root)
	if err == nil {
		if err := updateCoreSearchJson(root, routeVersions); err != nil {
			return err
		}
	} else {
		fmt.Printf("  ~ Warning: could not update core search.json: %s\n", err)
	}

	if copyPrev {
		fmt.Printf("\n✔ Version %q added (copied from previous version).\n", version)
		fmt.Printf("  Next steps:\n")
		fmt.Printf("    - Review and update the copied markdown files in docs/src/mds/%s/\n", version)
		fmt.Printf("    - Update docs/src/routes/api/%s.search.json/search.json with appropriate search data\n", version)
	} else {
		fmt.Printf("\n✔ Version %q added successfully.\n", version)
		fmt.Printf("  Next steps:\n")
		fmt.Printf("    - Add more markdown files under docs/src/mds/%s/ as needed\n", version)
		fmt.Printf("    - Update currentVersion in docs/src/mds/%s/index.md\n", version)
		fmt.Printf("    - Update docs/src/routes/api/%s.search.json/search.json with appropriate search data\n", version)
	}

	return nil
}

// ---------------------------------------------------------------------------
// Remove Version
// ---------------------------------------------------------------------------

func CmdVersionRemove(version string) error {
	if !rxVersion.MatchString(version) {
		return fmt.Errorf("invalid version name %q – must match v<number>, e.g. v3", version)
	}

	root, err := findProjectRoot()
	if err != nil {
		return err
	}

	// Prevent removal of v1 (the minimum required version)
	// if version == "v1" {
	// 	return fmt.Errorf("cannot remove %q – v1 is the minimum required version", version)
	// }

	dirsToRemove := []string{
		filepath.Join(root, "docs", "src", "routes", "(docs)", "docs", version),
		filepath.Join(root, "docs", "src", "routes", "api", version+".search.json"),
		filepath.Join(root, "docs", "src", "mds", version),
	}

	for _, dir := range dirsToRemove {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("  ~ Skipped (not found) %s\n", relPath(root, dir))
			continue
		}
		if err := os.RemoveAll(dir); err != nil {
			return fmt.Errorf("remove %s: %w", dir, err)
		}
		fmt.Printf("  ✔ Removed %s\n", relPath(root, dir))
	}

	// Update core search.json after removal
	routeVersions, err := discoverVersions(root)
	if err == nil {
		if err := updateCoreSearchJson(root, routeVersions); err != nil {
			return err
		}
	} else {
		fmt.Printf("  ~ Warning: could not update core search.json: %s\n", err)
	}

	fmt.Printf("\n✔ Version %q removed successfully.\n", version)
	return nil
}

// ---------------------------------------------------------------------------
// Sync Version
// ---------------------------------------------------------------------------

func CmdVersionSync() error {
	root, err := findProjectRoot()
	if err != nil {
		return err
	}

	fmt.Println("Syncing documentation versions...\n")

	// 1. Discover versions from route directories (source of truth)
	routeVersions, err := discoverVersions(root)
	if err != nil {
		return err
	}

	if len(routeVersions) == 0 {
		fmt.Println("  No versioned routes found.")
		return nil
	}

	fmt.Printf("  Found %d versioned route(s): %v\n\n", len(routeVersions), routeVersions)

	// 2. Build lookup sets for cross-checking
	mdsVersions, err := discoverMdsVersions(root)
	if err != nil {
		fmt.Printf("  Warning: could not scan mds dir: %s\n", err)
		mdsVersions = nil
	}
	searchVersions, err := discoverSearchVersions(root)
	if err != nil {
		fmt.Printf("  Warning: could not scan api dir: %s\n", err)
		searchVersions = nil
	}

	mdsSet := versionSet(mdsVersions)
	searchSet := versionSet(searchVersions)

	// 3. Ensure each route version has its mds dir and search json
	for _, v := range routeVersions {
		fmt.Printf("  [%s]\n", v)

		if mdsSet[v] {
			fmt.Printf("    ~ mds folder exists\n")
		} else {
			created, err := ensureMdsDir(root, v)
			if err != nil {
				fmt.Printf("    ✗ Error creating mds dir: %s\n", err)
			} else if created {
				fmt.Printf("    ✔ mds folder created\n")
			}
		}

		if searchSet[v] {
			fmt.Printf("    ~ search.json exists\n")
		} else {
			created, err := ensureSearchJsonDir(root, v)
			if err != nil {
				fmt.Printf("    ✗ Error creating search.json: %s\n", err)
			} else if created {
				fmt.Printf("    ✔ search.json created\n")
			}
		}

		if _, err := ensureRouteDir(root, v); err != nil {
			fmt.Printf("    ✗ Error ensuring route files: %s\n", err)
		}

		fmt.Println()
	}

	// 4. Cross-check: mds versions without routes
	for _, v := range mdsVersions {
		if !versionSet(routeVersions)[v] {
			fmt.Printf("  ⚠  Version %q has mds files but no route directory – run `version add %s` to create routes\n", v, v)
		}
	}

	// 5. Cross-check: search json versions without routes
	for _, v := range searchVersions {
		if !versionSet(routeVersions)[v] {
			fmt.Printf("  ⚠  Version %q has search.json but no route directory – run `version add %s` to create routes\n", v, v)
		}
	}

	// 6. Validate: check core search.json for duplicates
	{
		duplicates, err := validateCoreSearchJson(root)
		if err != nil {
			fmt.Printf("  ! Warning: could not validate core search.json: %s\n", err)
		} else if len(duplicates) > 0 {
			fmt.Printf("  ✗ Duplicate version(s) found in packages/core/search.json: %v\n", duplicates)
		} else {
			fmt.Println("  ✔ No duplicate versions in core search.json")
		}
	}

	// 7. Validate: check each version's currentVersion is proper semver
	{
		issues := validateCurrentVersions(root, routeVersions)
		if len(issues) > 0 {
			for _, issue := range issues {
				fmt.Printf("  ✗ %s\n", issue)
			}
		} else {
			fmt.Println("  ✔ All currentVersion values are valid semver")
		}
	}

	// 8. Rebuild core search.json from discovered route versions
	fmt.Println()
	if err := updateCoreSearchJson(root, routeVersions); err != nil {
		return err
	}

	fmt.Println("\n✔ Sync complete.")
	return nil
}

// ---------------------------------------------------------------------------
// Shared Helpers
// ---------------------------------------------------------------------------

// relPath returns a relative path from root for cleaner output.
func relPath(root, full string) string {
	r, err := filepath.Rel(root, full)
	if err != nil {
		return full
	}
	return r
}

// versionNum extracts the numeric part from a version name like "v3" -> 3.
func versionNum(version string) int {
	n, err := strconv.Atoi(strings.TrimPrefix(version, "v"))
	if err != nil {
		return 1
	}
	return n
}

// versionToSemver converts "v3" to "3.0.0".
func versionToSemver(version string) string {
	return fmt.Sprintf("%d.0.0", versionNum(version))
}

// coreSearchJsonPath returns the path to the core layout search.json file.
func coreSearchJsonPath(root string) string {
	return filepath.Join(root, "packages", "core", "src", "lib", "components", "layout", "search.json")
}

// readCoreSearchJson reads the current search.json array from the core package.
func readCoreSearchJson(root string) ([]string, error) {
	path := coreSearchJsonPath(root)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	var versions []string
	if err := json.Unmarshal(data, &versions); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return versions, nil
}

// writeCoreSearchJson writes the search.json array to the core package, sorted descending (latest first).
func writeCoreSearchJson(root string, versions []string) error {
	path := coreSearchJsonPath(root)
	// Sort descending by semver
	sort.Slice(versions, func(i, j int) bool {
		return compareSemver(versions[i], versions[j]) > 0
	})
	data, err := json.Marshal(versions)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	fmt.Printf("  ✔ Updated %s\n", relPath(root, path))
	return nil
}

// updateCoreSearchJson rebuilds the core search.json from the given route version names.
// It reads the actual currentVersion from each version's index.md to preserve existing values.
func updateCoreSearchJson(root string, routeVersions []string) error {
	semvers := make([]string, 0, len(routeVersions))
	for _, v := range routeVersions {
		semver := readCurrentVersion(root, v)
		semvers = append(semvers, semver)
	}
	return writeCoreSearchJson(root, semvers)
}

// readCurrentVersion reads the currentVersion frontmatter from a version's index.md.
// Falls back to "{n}.0.0" if it cannot be read or parsed.
var rxCurrentVersion = regexp.MustCompile(`(?m)^currentVersion:\s*(\S+)`)

func readCurrentVersion(root, version string) string {
	path := filepath.Join(root, "docs", "src", "mds", version, "index.md")
	data, err := os.ReadFile(path)
	if err != nil {
		return versionToSemver(version)
	}
	m := rxCurrentVersion.FindSubmatch(data)
	if len(m) < 2 {
		return versionToSemver(version)
	}
	return string(m[1])
}

// compareSemver compares two semver strings. Returns >0 if a > b, <0 if a < b, 0 if equal.
func compareSemver(a, b string) int {
	parse := func(s string) []int {
		parts := strings.Split(s, ".")
		nums := make([]int, 3)
		for i, p := range parts {
			if i >= 3 {
				break
			}
			n, _ := strconv.Atoi(p)
			nums[i] = n
		}
		return nums
	}
	aNums := parse(a)
	bNums := parse(b)
	for i := 0; i < 3; i++ {
		if aNums[i] != bNums[i] {
			return aNums[i] - bNums[i]
		}
	}
	return 0
}

// ---------------------------------------------------------------------------
// Validation Helpers
// ---------------------------------------------------------------------------

// isValidSemver checks if a string is a valid semver like X.Y.Z.
var rxSemver = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

func isValidSemver(s string) bool {
	return rxSemver.MatchString(s)
}

// validateCoreSearchJson checks the core search.json for duplicate values.
func validateCoreSearchJson(root string) ([]string, error) {
	entries, err := readCoreSearchJson(root)
	if err != nil {
		return nil, err
	}

	seen := make(map[string]bool)
	var duplicates []string
	for _, e := range entries {
		if seen[e] {
			duplicates = append(duplicates, e)
		} else {
			seen[e] = true
		}
	}
	return duplicates, nil
}

// validateCurrentVersions checks each version's index.md for a valid currentVersion.
func validateCurrentVersions(root string, routeVersions []string) []string {
	var issues []string
	for _, v := range routeVersions {
		semver := readCurrentVersion(root, v)
		if !isValidSemver(semver) {
			issues = append(issues, fmt.Sprintf("%s: currentVersion %q is not a valid semver (expected X.Y.Z)", v, semver))
		}
	}
	return issues
}

// ---------------------------------------------------------------------------
// Help / Subcommand Routing
// ---------------------------------------------------------------------------

func PrintVersionUsage() {
	fmt.Println(`Usage: mds version <command> [version] [flags]

Manage documentation versioned routes.

Commands:
  add <version> [--copy]  Add a new version (e.g. v3)
                   Creates all routes, API search, and mds scaffolding.
                   With --copy, copies the entire mds content from the
                   previous version instead of creating from scratch.

  remove <version> Remove an existing version (e.g. v3)
                   Deletes all files created by 'add'.

  sync             Sync all versions: ensures every versioned route has
                   its corresponding mds folder with index.md, and its
                   API search.json endpoint.

Examples:
  mds version add v3
  mds version add v3 --copy
  mds version remove v3
  mds version sync
`)
}

// HandleVersionSubcommand routes "version add" / "version remove" / "version sync".
func HandleVersionSubcommand(args []string) {
	if len(args) == 0 {
		PrintVersionUsage()
		os.Exit(1)
	}

	cmd := args[0]
	switch cmd {
	case "add":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: missing version argument.\n\n")
			PrintVersionUsage()
			os.Exit(1)
		}
		copyPrev := false
		versionArg := ""
		for _, a := range args[1:] {
			if a == "--copy" {
				copyPrev = true
			} else if !strings.HasPrefix(a, "-") {
				versionArg = a
			}
		}
		if versionArg == "" {
			fmt.Fprintf(os.Stderr, "Error: missing version argument.\n\n")
			PrintVersionUsage()
			os.Exit(1)
		}
		if err := CmdVersionAdd(versionArg, copyPrev); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

	case "remove":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: missing version argument.\n\n")
			PrintVersionUsage()
			os.Exit(1)
		}
		version := args[1]
		if err := CmdVersionRemove(version); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

	case "sync":
		if err := CmdVersionSync(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

	case "help", "--help", "-h":
		PrintVersionUsage()

	default:
		fmt.Fprintf(os.Stderr, "Error: unknown version command %q.\n\n", cmd)
		PrintVersionUsage()
		os.Exit(1)
	}
}
