package gfm_serve

import "golang.org/x/net/html"

// s is a helper function to create a text node.
//
//	s("Hello, World!")
func s(text string) *html.Node {
	return &html.Node{
		Type: html.TextNode,
		Data: text,
	}
}

// a is a helper function to create a slice of html.Attributes from a map.
//
//	a(map[string]string{"data-something": "not-something"})
func a(attrs map[string]string) []html.Attribute {
	attributes := make([]html.Attribute, len(attrs))
	i := 0
	for k, v := range attrs {
		attributes[i] = html.Attribute{Key: k, Val: v}
		i += 1
	}
	return attributes
}

// h is a helper function to create an html.Node with the given tag, attributes and children.
//
//	h("div",
//		a(map[string]string{"data-something": "not-something"}),
//		s("Hello, World!")
//		h("div",
//			nil,
//			s("Another one!"),
//		),
//	)
func h(tag string, attrs []html.Attribute, children ...*html.Node) *html.Node {
	root := &html.Node{Type: html.ElementNode, Data: tag, Attr: attrs}
	for _, child := range children {
		if child.Parent != nil {
			child.Parent.RemoveChild(child)
		}
		root.AppendChild(child)
	}
	return root
}
