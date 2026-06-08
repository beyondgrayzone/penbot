package mds

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ---------------------------------------------------------------------------
// Sections Subcommand: list / add / remove sections in velite.config.js
// ---------------------------------------------------------------------------

func sectionsConfigPath(root string) string {
	return filepath.Join(root, "docs", "velite.config.js")
}

// findSectionsArray finds the sections array in the config.
// Returns (everythingBeforeContent, arrayContent, everythingAfterContent, error).
// where arrayContent is what goes between [ and ].
func findSectionsArray(data []byte) (before, content, after string, err error) {
	rxLine := regexp.MustCompile(`(?m)^(\s*)const\s+sections\s*=\s*`)
	loc := rxLine.FindIndex(data)
	if loc == nil {
		return "", "", "", fmt.Errorf("no 'sections' assignment found")
	}

	rest := string(data[loc[1]:])
	// Full regex: (\[)([\s\S]*?)(\])
	// Groups:      1     2          3
	rxArray := regexp.MustCompile(`(\[)([\s\S]*?)(\])`)
	m := rxArray.FindStringSubmatchIndex(rest)
	if m == nil {
		return "", "", "", fmt.Errorf("no array found after 'sections ='")
	}

	// m[0]:m[1] = full match  [...]
	// m[2]:m[3] = group 1    "["
	// m[4]:m[5] = group 2    content between [ and ]
	// m[6]:m[7] = group 3    "]"

	before = string(data[:loc[1]]) + rest[:m[3]] // everything up to and including [
	content = rest[m[4]:m[5]]                     // content between [ and ]
	after = rest[m[6]:]                           // from ] onward (includes ])
	return
}

// readSections parses the current section names from the array content.
func readSections(content string) []string {
	rxString := regexp.MustCompile(`"([^"]*)"`)
	matches := rxString.FindAllStringSubmatch(content, -1)
	var sections []string
	for _, m := range matches {
		sections = append(sections, m[1])
	}
	return sections
}

// writeSections replaces the sections array content in the config file.
func writeSections(root, newContent string) error {
	configPath := sectionsConfigPath(root)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", configPath, err)
	}

	before, _, after, err := findSectionsArray(data)
	if err != nil {
		return err
	}

	newData := before + newContent + after
	return os.WriteFile(configPath, []byte(newData), 0644)
}

// rebuildArrayContent formats the sections array content given a list of names.
// If isInline is true, produces single-line: "a", "b", "c"
// Otherwise produces multi-line with each item indented on its own line.
func rebuildArrayContent(indent string, sections []string, isInline bool) string {
	if len(sections) == 0 {
		return ""
	}
	var sb strings.Builder
	if isInline {
		for i, s := range sections {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("%q", s))
		}
	} else {
		sb.WriteString("\n")
		for _, s := range sections {
			sb.WriteString(fmt.Sprintf("%s%q,\n", indent, s))
		}
	}
	return sb.String()
}

// detectIndent guesses the indent from the current array content, defaults to "\t".
func detectIndent(content string) string {
	if strings.TrimSpace(content) == "" {
		return "\t"
	}
	// Single-line array (inline) - no indent detected, use default
	if !strings.Contains(content, "\n") {
		return "\t"
	}
	for _, line := range strings.Split(content, "\n") {
		tr := strings.TrimSpace(line)
		if tr != "" {
			return line[:len(line)-len(tr)]
		}
	}
	return "\t"
}

// ---------------------------------------------------------------------------
// List
// ---------------------------------------------------------------------------

func CmdSectionsList() error {
	root, err := findProjectRoot()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(sectionsConfigPath(root))
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	_, content, _, err := findSectionsArray(data)
	if err != nil {
		return err
	}

	sections := readSections(content)
	if len(sections) == 0 {
		fmt.Println("Sections: (none)")
		return nil
	}

	fmt.Println("Sections:")
	for _, s := range sections {
		fmt.Printf("  - %s\n", s)
	}
	return nil
}

// ---------------------------------------------------------------------------
// Add
// ---------------------------------------------------------------------------

func CmdSectionAdd(name string) error {
	root, err := findProjectRoot()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(sectionsConfigPath(root))
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	_, content, _, err := findSectionsArray(data)
	if err != nil {
		return err
	}

	sections := readSections(content)
	for _, s := range sections {
		if s == name {
			return fmt.Errorf("section %q already exists", name)
		}
	}

	indent := detectIndent(content)
	isInline := !strings.Contains(content, "\n")
	sections = append(sections, name)
	newContent := rebuildArrayContent(indent, sections, isInline)

	if err := writeSections(root, newContent); err != nil {
		return err
	}
	fmt.Printf("  ✔ Updated docs/velite.config.js\n")

	// Also update navigation.json
	navSections, err := readNavigationSections(root)
	if err != nil {
		return fmt.Errorf("read navigation.json: %w", err)
	}

	// Check if already exists in navigation.json too
	for _, s := range navSections {
		if s == name {
			fmt.Printf("  ✔ Section %q already present in navigation.json\n", name)
			return nil
		}
	}

	navSections = append(navSections, name)
	if err := writeNavigationSections(root, navSections); err != nil {
		return fmt.Errorf("update navigation.json: %w", err)
	}
	fmt.Printf("  ✔ Updated docs/src/lib/navigation.json\n")
	return nil
}

// ---------------------------------------------------------------------------
// Remove
// ---------------------------------------------------------------------------

func CmdSectionRemove(name string) error {
	root, err := findProjectRoot()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(sectionsConfigPath(root))
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	_, content, _, err := findSectionsArray(data)
	if err != nil {
		return err
	}

	sections := readSections(content)
	found := false
	var filtered []string
	for _, s := range sections {
		if s == name {
			found = true
		} else {
			filtered = append(filtered, s)
		}
	}
	if !found {
		return fmt.Errorf("section %q not found", name)
	}

	indent := detectIndent(content)
	isInline := !strings.Contains(content, "\n")
	newContent := rebuildArrayContent(indent, filtered, isInline)

	if err := writeSections(root, newContent); err != nil {
		return err
	}
	fmt.Printf("  ✔ Updated docs/velite.config.js\n")

	// Also update navigation.json
	navSections, err := readNavigationSections(root)
	if err != nil {
		return fmt.Errorf("read navigation.json: %w", err)
	}

	var filteredNav []string
	foundInNav := false
	for _, s := range navSections {
		if s == name {
			foundInNav = true
		} else {
			filteredNav = append(filteredNav, s)
		}
	}

	if !foundInNav {
		fmt.Printf("  ✔ Section %q not found in navigation.json (nothing to remove)\n", name)
		return nil
	}

	if err := writeNavigationSections(root, filteredNav); err != nil {
		return fmt.Errorf("update navigation.json: %w", err)
	}
	fmt.Printf("  ✔ Updated docs/src/lib/navigation.json\n")
	return nil
}

// ---------------------------------------------------------------------------
// Navigation JSON helpers
// ---------------------------------------------------------------------------

func navigationJSONPath(root string) string {
	return filepath.Join(root, "docs", "src", "lib", "navigation.json")
}

type navFile struct {
	Anchors  []json.RawMessage `json:"anchors"`
	Sections []navSection      `json:"sections"`
}

type navSection struct {
	Title string `json:"title"`
}

// readNavigationSections reads the section titles from navigation.json.
func readNavigationSections(root string) ([]string, error) {
	path := navigationJSONPath(root)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	var nf navFile
	if err := json.Unmarshal(data, &nf); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	var titles []string
	for _, s := range nf.Sections {
		titles = append(titles, s.Title)
	}
	return titles, nil
}

// writeNavigationSections writes the section list back to navigation.json,
// preserving the anchors structure exactly as-is.
func writeNavigationSections(root string, titles []string) error {
	path := navigationJSONPath(root)
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	var nf navFile
	if err := json.Unmarshal(data, &nf); err != nil {
		return fmt.Errorf("parse %s: %w", path, err)
	}

	nf.Sections = make([]navSection, len(titles))
	for i, t := range titles {
		nf.Sections[i] = navSection{Title: t}
	}

	out, err := json.MarshalIndent(nf, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	// Add trailing newline to match existing file style
	out = append(out, '\n')

	return os.WriteFile(path, out, 0644)
}

// ---------------------------------------------------------------------------
// Routing
// ---------------------------------------------------------------------------

func PrintSectionsUsage() {
	fmt.Println(`Usage: mds sections <command> [name]

Manage documentation sections in docs/velite.config.js and docs/src/lib/navigation.json.

Commands:
  list              List all defined sections.

  add <name>        Add a new section to the end of the array.

  remove <name>     Remove an existing section from the array.

Examples:
  mds sections list
  mds sections add "API Reference"
  mds sections remove "Utilities"
`)
}

func HandleSectionsSubcommand(args []string) {
	if len(args) == 0 {
		PrintSectionsUsage()
		os.Exit(1)
	}

	cmd := args[0]
	switch cmd {
	case "list":
		if err := CmdSectionsList(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	case "add":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: missing section name argument.\n\n")
			PrintSectionsUsage()
			os.Exit(1)
		}
		name := args[1]
		if err := CmdSectionAdd(name); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	case "remove":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: missing section name argument.\n\n")
			PrintSectionsUsage()
			os.Exit(1)
		}
		name := args[1]
		if err := CmdSectionRemove(name); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	case "help", "--help", "-h":
		PrintSectionsUsage()
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown sections command %q.\n\n", cmd)
		PrintSectionsUsage()
		os.Exit(1)
	}
}
