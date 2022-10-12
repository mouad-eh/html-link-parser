package link

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

type Link struct {
	Href string
	Text string
}

// links, err := link.Parse(r io.Reader)
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, BuildLink(node))
	}
	return links, nil
}

func BuildLink(n *html.Node) Link {
	var link Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}
	link.Text = text(n)
	return link
}
func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	// if n.Type == html.CommentNode {
	// 	return n.Data
	// }
	// if you want comment to be part of text.
	if n.Type != html.ElementNode {
		return ""
	}
	var nodeText string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodeText += text(c)
		// this is not a clean way of building strings
		// you can look at "byte buffer golang"
	}
	nodeText = strings.Join(strings.Fields(nodeText), " ")
	return nodeText
}
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
