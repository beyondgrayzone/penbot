package mds

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Input text matches the example provided in the prompt
var content = `...`

// ---------------------------------------------------------------------------
// Structs and Constants
// ---------------------------------------------------------------------------

type ComponentContext struct {
	TagName string
	Props   map[string]string
	Buffer  strings.Builder
}

// Regex patterns
var (
	rxFrontmatterSep = regexp.MustCompile(`^---$`)
	rxScriptOpen     = regexp.MustCompile(`^<script>\s*$`)
	rxScriptClose    = regexp.MustCompile(`^</script>\s*$`)
	rxImport         = regexp.MustCompile(`import\s+\{(.+)\}\s+from\s+"@penbot/core";`)
	rxCodeBlock      = regexp.MustCompile("^```")

	// Matches <TagName ...> (Opening block)
	rxTagOpen = regexp.MustCompile(`^<(\w+)([^>]*)>$`)
	// Matches </TagName> (Closing block)
	rxTagClose = regexp.MustCompile(`^</(\w+)>$`)

	// Matches <TagName ... /> (Self-closing)
	rxSelfClosing = regexp.MustCompile(`^<(\w+)([^>]*)/>$`)

	// Matches <TagName ...>Content</TagName> (Inline)
	rxInlineContainer = regexp.MustCompile(`^<(\w+)([^>]*)>(.*)</(\w+)>$`)

	// Detect wrapper divs to strip them
	rxDivWrapper = regexp.MustCompile(`(?i)^\s*</?div.*>?$`)

	// Matches attr="val", attr={val}, or boolean attr
	rxAttributes = regexp.MustCompile(`(\w+)(?:=(?:"([^"]*)"|\{([^}]*)\}))?`)
)

// ---------------------------------------------------------------------------
// Processing Logic
// ---------------------------------------------------------------------------

func Start(in, out string) (err error) {
	b, err := os.ReadFile(in)
	if err != nil {
		err = fmt.Errorf("Input file err: %w", err.Error())
		return
	}

	final, err := processContent(string(b))
	if err != nil {
		return err
	}

	dir := filepath.Dir(out)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(out, []byte(final), 0600)
}

func processContent(input string) (final string, err error) {
	var output strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(input))

	// State flags
	inFrontmatter := false
	frontmatterDone := false
	inScript := false
	inCodeBlock := false

	// Data stores
	importedComponents := make(map[string]bool)
	var componentStack []*ComponentContext

	// Helper to determine if we are currently inside a component handling hierarchy
	inComponent := func() bool {
		return len(componentStack) > 0
	}

	// Helper to write result (either to buffer or main output)
	writeResult := func(s string) {
		if inComponent() {
			componentStack[len(componentStack)-1].Buffer.WriteString(s)
		} else {
			output.WriteString(s)
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// 0. Strip wrapper divs (Cleanup)
		if rxDivWrapper.MatchString(trimmedLine) {
			continue
		}

		// 1. Handle Frontmatter
		if !frontmatterDone {
			if rxFrontmatterSep.MatchString(trimmedLine) {
				if inFrontmatter {
					inFrontmatter = false
					frontmatterDone = true
					output.WriteString("\n")
					continue
				} else {
					inFrontmatter = true
					continue
				}
			}
			if inFrontmatter {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					val := strings.TrimSpace(parts[1])
					if key == "title" {
						output.WriteString(fmt.Sprintf("# %s\n", val))
					} else if key == "description" {
						output.WriteString(fmt.Sprintf("> %s\n", val))
					}
				}
				continue
			}
		}

		// 2. Handle Code Blocks
		if rxCodeBlock.MatchString(line) {
			inCodeBlock = !inCodeBlock
			if inComponent() {
				componentStack[len(componentStack)-1].Buffer.WriteString(line + "\n")
			} else {
				output.WriteString(line + "\n")
			}
			continue
		}
		if inCodeBlock {
			if inComponent() {
				componentStack[len(componentStack)-1].Buffer.WriteString(line + "\n")
			} else {
				output.WriteString(line + "\n")
			}
			continue
		}

		// 3. Handle Script & Imports
		if !inScript && rxScriptOpen.MatchString(trimmedLine) {
			inScript = true
			continue
		}
		if inScript {
			if rxScriptClose.MatchString(trimmedLine) {
				inScript = false
				continue
			}
			matches := rxImport.FindStringSubmatch(line)
			if len(matches) > 1 {
				imports := strings.Split(matches[1], ",")
				for _, imp := range imports {
					importedComponents[strings.TrimSpace(imp)] = true
				}
			}
			continue
		}

		// 4. Handle HTML/Component Tags

		// A. Self-Closing Tag (e.g. <Checkbox />)
		if matches := rxSelfClosing.FindStringSubmatch(trimmedLine); len(matches) > 0 {
			tagName := matches[1]
			attrStr := matches[2]
			if importedComponents[tagName] {
				props := parseAttributes(attrStr)
				ctx := &ComponentContext{TagName: tagName, Props: props}
				writeResult(transformComponent(ctx))
				continue
			}
		}

		// B. Inline Container (e.g. <Step>Title</Step>)
		if matches := rxInlineContainer.FindStringSubmatch(trimmedLine); len(matches) > 0 {
			startTag := matches[1]
			attrStr := matches[2]
			innerContent := matches[3]
			endTag := matches[4]

			if startTag == endTag && importedComponents[startTag] {
				props := parseAttributes(attrStr)
				ctx := &ComponentContext{TagName: startTag, Props: props}
				ctx.Buffer.WriteString(innerContent)
				writeResult(transformComponent(ctx))
				continue
			}
		}

		// C. Opening Tag (Start of Block)
		if matches := rxTagOpen.FindStringSubmatch(trimmedLine); len(matches) > 0 {
			tagName := matches[1]
			attrStr := matches[2]

			if importedComponents[tagName] {
				props := parseAttributes(attrStr)
				ctx := &ComponentContext{
					TagName: tagName,
					Props:   props,
				}
				componentStack = append(componentStack, ctx)
				continue
			}
		}

		// D. Closing Tag (End of Block)
		if matches := rxTagClose.FindStringSubmatch(trimmedLine); len(matches) > 0 {
			tagName := matches[1]

			if importedComponents[tagName] {
				if len(componentStack) > 0 {
					current := componentStack[len(componentStack)-1]
					if current.TagName == tagName {
						componentStack = componentStack[:len(componentStack)-1]
						writeResult(transformComponent(current))
						continue
					}
				}
			}
		}

		// 5. Handle Content
		if inComponent() {
			componentStack[len(componentStack)-1].Buffer.WriteString(line + "\n")
		} else {
			output.WriteString(line + "\n")
		}
	}

	final = output.String()
	return
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func parseAttributes(attrStr string) map[string]string {
	props := make(map[string]string)
	matches := rxAttributes.FindAllStringSubmatch(attrStr, -1)
	for _, m := range matches {
		key := m[1]
		val := ""
		if m[2] != "" {
			val = m[2]
		} else if m[3] != "" {
			val = m[3]
		} else {
			val = "true" // Boolean attribute
		}
		props[key] = val
	}
	return props
}

func transformComponent(ctx *ComponentContext) string {
	content := strings.TrimSpace(ctx.Buffer.String())

	switch ctx.TagName {
	case "CardGrid":
		return content + "\n"

	case "Card":
		title := ctx.Props["title"]
		href := ctx.Props["href"]
		var header string
		if href != "" {
			header = fmt.Sprintf("## [%s](%s)", title, href)
		} else {
			header = fmt.Sprintf("## %s", title)
		}
		return fmt.Sprintf("\n%s\n\n%s\n", header, content)

	case "Steps":
		// Iterate lines, add > to text, keep headers and code blocks as is
		lines := strings.Split(content, "\n")
		var sb strings.Builder
		inCode := false

		for _, line := range lines {
			trim := strings.TrimSpace(line)

			// Detect Code Block
			if strings.HasPrefix(trim, "```") {
				inCode = !inCode
				sb.WriteString(line + "\n")
				continue
			}

			// If inside code block, print raw
			if inCode {
				sb.WriteString(line + "\n")
				continue
			}

			// If it's a header generated by <Step> (starts with ##)
			if strings.HasPrefix(trim, "##") {
				sb.WriteString(line + "\n")
				continue
			}

			// Keep empty lines empty
			if trim == "" {
				sb.WriteString("\n")
				continue
			}

			// It's text content, quote it
			sb.WriteString("> " + line + "\n")
		}
		return sb.String()

	case "Step":
		// <Step>Title</Step> -> ## Title
		// Return with surrounding newlines to ensure proper spacing in the Steps loop
		return fmt.Sprintf("\n## %s\n", content)

	case "Callout":
		title := ctx.Props["title"]
		// Naive blockquote logic for simple callouts (per earlier instructions)
		quotedContent := strings.ReplaceAll(content, "\n", "\n> ")
		quotedContent = "> " + quotedContent
		return fmt.Sprintf("\n## %s\n\n%s\n", title, quotedContent)

	case "Collapsible":
		title := ctx.Props["title"]
		return fmt.Sprintf("\n## %s\n%s\n", title, content)

	case "Tabs":
		return content + "\n"

	case "TabItem":
		val := ctx.Props["value"]
		return fmt.Sprintf("\n### %s\n%s\n", val, content)

	case "Checkbox":
		checked := ctx.Props["checked"] == "true"
		if checked {
			return "- [x] "
		}
		return "- [ ] "

	case "Label":
		return content + "\n"

	case "PropField":
		name := ctx.Props["name"]
		typ := ctx.Props["type"]
		def := ctx.Props["defaultValue"]

		var desc string
		var nestedContent string

		splitIdx := -1
		if idx := strings.Index(content, "\n##"); idx != -1 {
			splitIdx = idx
		}

		if splitIdx != -1 {
			desc = content[:splitIdx]
			nestedContent = content[splitIdx:]
		} else {
			desc = content
		}

		desc = strings.ReplaceAll(strings.TrimSpace(desc), "\n", " ")

		sb := strings.Builder{}
		sb.WriteString("| Prop | Type | Default | Description |\n")
		sb.WriteString("| :--- | :--- | :--- | :--- |\n")
		sb.WriteString(fmt.Sprintf("| `%s` | `%s` | `%s` | %s |\n", name, typ, def, desc))

		if nestedContent != "" {
			sb.WriteString(nestedContent)
		}

		return sb.String()

	case "NakedContainer":
		return content

	default:
		return ""
	}
}
