package client

import (
	"golang.org/x/net/html"
	"strings"
)

func find(tag, classPrefix string, start *html.Node) *html.Node {
	var helper func(n *html.Node) *html.Node
	helper = func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.Data == tag {
			var class string
			for _, attr := range n.Attr {
				if attr.Key != "class" {
					continue
				}
				class = attr.Val
			}
			if strings.HasPrefix(class, classPrefix) {
				return n
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if res := helper(c); res != nil {
				return res
			}
		}
		return nil
	}
	return helper(start)
}

func flatten(node *html.Node) (str string, err error) {
	b := strings.Builder{}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			_, err = b.WriteString(c.Data)
		}
	}
	str = b.String()
	return
}
