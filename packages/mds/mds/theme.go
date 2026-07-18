package mds

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// ---------------------------------------------------------------------------
// Theme Subcommand: list / apply
// ---------------------------------------------------------------------------

// themeStylesDir returns the path to the core theme CSS files.
func themeStylesDir(root string) string {
	return filepath.Join(root, "packages", "core", "src", "lib", "styles")
}

// discoverThemes scans the core styles directory for theme-*.css files
// and returns the extracted theme names sorted alphabetically.
func discoverThemes(root string) ([]string, error) {
	dir := themeStylesDir(root)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read styles dir %s: %w", dir, err)
	}

	var themes []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasPrefix(name, "theme-") || !strings.HasSuffix(name, ".css") {
			continue
		}
		theme := strings.TrimPrefix(name, "theme-")
		theme = strings.TrimSuffix(theme, ".css")
		themes = append(themes, theme)
	}

	sort.Strings(themes)
	return themes, nil
}

// siteConfigPath returns the path to the docs site-config.json.
func siteConfigPath(root string) string {
	return filepath.Join(root, "docs", "src", "lib", "site-config.json")
}

// siteConfigData represents the JSON structure of site-config.json.
type siteConfigData struct {
	Name           string `json:"name"`
	DefaultTheme   string `json:"defaultTheme"`
	ThemeTimestamp int64  `json:"themeTimestamp"`
	// Other fields are preserved via raw JSON  see readSiteConfig / writeSiteConfig.
}

// readSiteConfig reads the raw site-config.json, returning the raw bytes
// and a partial struct for known fields.
func readSiteConfig(root string) (raw []byte, cfg *siteConfigData, err error) {
	path := siteConfigPath(root)
	raw, err = os.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("read %s: %w", path, err)
	}

	var c siteConfigData
	if err := json.Unmarshal(raw, &c); err != nil {
		return nil, nil, fmt.Errorf("parse %s: %w", path, err)
	}

	return raw, &c, nil
}

// writeSiteConfig updates defaultTheme and themeTimestamp fields in the
// site-config.json file. It replaces the matching lines entirely so the
// original key ordering in the file is preserved.
func writeSiteConfig(root string, raw []byte, defaultTheme string, timestamp int64) error {
	path := siteConfigPath(root)
	rawStr := string(raw)

	// Ensure trailing newline
	if !strings.HasSuffix(rawStr, "\n") {
		rawStr += "\n"
	}

	// Update "defaultTheme" line (always with trailing comma since it's
	// expected to be followed by "themeTimestamp").
	rxDefaultTheme := regexp.MustCompile(`(?m)^\t"defaultTheme":\s*".*",?$`)
	newDefaultLine := fmt.Sprintf("\t\"defaultTheme\": %q,", defaultTheme)
	if rxDefaultTheme.MatchString(rawStr) {
		rawStr = rxDefaultTheme.ReplaceAllString(rawStr, newDefaultLine)
	} else {
		// Insert before "themeTimestamp" if present, otherwise at end before "}"
		rxTimestampKey := regexp.MustCompile(`(?m)^\t"themeTimestamp":`)
		if rxTimestampKey.MatchString(rawStr) {
			rawStr = rxTimestampKey.ReplaceAllString(rawStr, newDefaultLine+"\n\t"+`"themeTimestamp":`)
		} else {
			rawStr = strings.Replace(rawStr, "\n}", "\n"+newDefaultLine+"\n}", 1)
		}
	}

	// Update "themeTimestamp" line (never with trailing comma since it's
	// expected to be the last key before "}").
	rxTimestamp := regexp.MustCompile(`(?m)^\t"themeTimestamp":\s*\d+,?$`)
	newTimestampLine := fmt.Sprintf("\t\"themeTimestamp\": %d", timestamp)
	if rxTimestamp.MatchString(rawStr) {
		rawStr = rxTimestamp.ReplaceAllString(rawStr, newTimestampLine)
	} else {
		rawStr = strings.Replace(rawStr, "\n}", "\n"+newTimestampLine+"\n}", 1)
	}

	return os.WriteFile(path, []byte(rawStr), 0644)
}

// ---------------------------------------------------------------------------
// List
// ---------------------------------------------------------------------------

func CmdThemeList() error {
	root, err := findProjectRoot()
	if err != nil {
		return err
	}

	themes, err := discoverThemes(root)
	if err != nil {
		return err
	}

	if len(themes) == 0 {
		fmt.Println("Themes: (none found)")
		return nil
	}

	// Read current default from site-config.json
	_, cfg, err := readSiteConfig(root)
	currentDefault := ""
	if err == nil {
		currentDefault = cfg.DefaultTheme
	}

	fmt.Println("Available themes:")
	for _, t := range themes {
		mark := "  "
		if t == currentDefault {
			mark = " *"
		}
		fmt.Printf("  %s %s\n", mark, t)
	}
	fmt.Println()
	fmt.Println("Current default:", currentDefault)
	return nil
}

// ---------------------------------------------------------------------------
// Apply
// ---------------------------------------------------------------------------

func CmdThemeApply(name string) error {
	root, err := findProjectRoot()
	if err != nil {
		return err
	}

	// Validate the theme exists
	themes, err := discoverThemes(root)
	if err != nil {
		return err
	}

	found := false
	for _, t := range themes {
		if t == name {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("unknown theme %q. Available themes: %s", name, strings.Join(themes, ", "))
	}

	// Read existing site-config.json
	raw, cfg, err := readSiteConfig(root)
	if err != nil {
		return err
	}

	oldTheme := cfg.DefaultTheme
	timestamp := time.Now().UnixMilli()

	if err := writeSiteConfig(root, raw, name, timestamp); err != nil {
		return err
	}

	fmt.Printf("  ✔ Updated docs/src/lib/site-config.json\n")
	if oldTheme != "" && oldTheme != name {
		fmt.Printf("  Theme changed: %s → %s\n", oldTheme, name)
	} else {
		fmt.Printf("  Theme set to:  %s\n", name)
	}
	fmt.Printf("  Timestamp:     %d\n", timestamp)

	return nil
}

// ---------------------------------------------------------------------------
// Help / Subcommand Routing
// ---------------------------------------------------------------------------

func PrintThemeUsage() {
	fmt.Println(`Usage: mds theme <command> [name]

Manage the site's default theme and timestamp.

Commands:
  list              List all available themes (from packages/core styles).
                    The current default is marked with *.

  apply <name>      Set the default theme and update the theme timestamp
                    to the current time. Overrides returning visitors'
                    theme preferences on next deploy.

Examples:
  mds theme list
  mds theme apply oceanic
  mds theme apply forest
`)
}

func HandleThemeSubcommand(args []string) {
	if len(args) == 0 {
		PrintThemeUsage()
		os.Exit(1)
	}

	cmd := args[0]
	switch cmd {
	case "list":
		if err := CmdThemeList(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	case "apply":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: missing theme name argument.\n\n")
			PrintThemeUsage()
			os.Exit(1)
		}
		name := args[1]
		if err := CmdThemeApply(name); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	case "help", "--help", "-h":
		PrintThemeUsage()
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown theme command %q.\n\n", cmd)
		PrintThemeUsage()
		os.Exit(1)
	}
}
