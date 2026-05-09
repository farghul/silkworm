package main

import (
	"strings"

	"golang.org/x/net/html"
)

func converter(input string) (*ADFDoc, error) {
	root, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return nil, err
	}

	adf := &ADFDoc{
		Type:    "doc",
		Version: 1,
		Content: []ADFBlock{},
	}

	var changelogNode *html.Node

	// 1. Find first <h2> containing "Changelog"
	var findChangelog func(*html.Node)
	findChangelog = func(n *html.Node) {
		if changelogNode != nil {
			return
		}

		if n.Type == html.ElementNode && n.Data == "h2" {
			if strings.Contains(strings.ToLower(extractText(n)), "changelog") {
				changelogNode = n
				return
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findChangelog(c)
		}
	}

	findChangelog(root)

	if changelogNode == nil {
		return adf, nil // no changelog found
	}

	// 2. Add the <h2> heading
	adf.Content = append(adf.Content, makeHeading(changelogNode))

	// 3. Walk siblings AFTER <h2> until first <ul>
	for node := changelogNode.NextSibling; node != nil; node = node.NextSibling {

		if node.Type != html.ElementNode {
			continue
		}

		switch node.Data {
		case "h3", "h4", "h5", "h6":
			adf.Content = append(adf.Content, makeHeading(node))

		case "ul":
			adf.Content = append(adf.Content, makeBulletList(node))
			return adf, nil // ✅ STOP after first <ul>

		case "p":
			adf.Content = append(adf.Content, makeParagraph(node))
		}
	}

	return adf, nil
}

// --- HELPERS ---

func makeHeading(n *html.Node) ADFBlock {
	level := int(n.Data[1] - '0')
	return ADFBlock{
		Type:  "heading",
		Attrs: map[string]int{"level": level},
		Content: []ADFInline{
			{Type: "text", Text: extractText(n)},
		},
	}
}

func makeParagraph(n *html.Node) ADFBlock {
	return ADFBlock{
		Type: "paragraph",
		Content: []ADFInline{
			{Type: "text", Text: extractText(n)},
		},
	}
}

func makeBulletList(n *html.Node) ADFBlock {
	list := ADFBlock{
		Type:    "bulletList",
		Content: []ADFInline{},
	}

	for li := n.FirstChild; li != nil; li = li.NextSibling {
		if li.Type == html.ElementNode && li.Data == "li" {
			list.Content = append(list.Content, ADFInline{
				Type: "listItem",
				Content: []ADFInline{
					{
						Type: "paragraph",
						Content: []ADFInline{
							{Type: "text", Text: extractText(li)},
						},
					},
				},
			})
		}
	}

	return list
}

// Extract text + decode HTML entities
func extractText(n *html.Node) string {
	var parts []string

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.TextNode {
			text := strings.TrimSpace(html.UnescapeString(n.Data))
			if text != "" {
				parts = append(parts, text)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(n)
	return strings.Join(parts, " ")
}
