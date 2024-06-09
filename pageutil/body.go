package pageutil

import (
	"fmt"

	"golang.org/x/net/html"
)

// 找到视频概要信息
func FindProfile(doc *html.Node) (profile []string, link []string) {
	var scan func(n *html.Node)
	scan = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == "desc-info-text" {
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						switch c.Type {
						case html.ElementNode:
							if c.Data == "a" {
								fmt.Print(c.FirstChild.Data)
								link = append(link, c.FirstChild.Data+",")
								profile = append(profile, c.FirstChild.Data)
							}
						case html.TextNode:
							profile = append(profile, c.Data)
						}
					}
				}
			}
		}
		// 递归遍历过程
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			scan(c)
		}
	}

	scan(doc)
	return profile, link
}

// 找到视频标签
func FindVideoTags(doc *html.Node) (tags []string) {
	var scan func(n *html.Node)
	scan = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == "v_tag" {
					i := n.FirstChild
					tags = scanTags(i)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			scan(c)
		}
	}
	scan(doc)
	return tags
}

func scanTags(s *html.Node) (tag []string) {
	var t *html.Node
	var n *html.Node
	//tag-panel 标签容器s.FirstChild
	for i := s.FirstChild; i != nil; i = i.NextSibling {
		//单标签容器遍历属性
		for _, attr := range i.Attr {
			if attr.Key == "class" && attr.Val == "tag not-btn-tag" {
				switchTag := true
				switch i.FirstChild.Attr[0].Val {
				case "topic-tag":
					t = i.FirstChild.FirstChild
					switchTag = false
				default:
					n = i.FirstChild.FirstChild
				}
				if switchTag {
					for n = n.FirstChild; n != nil; n = n.NextSibling {
						tag = append(tag, n.Data)
					}

				} else {
					for t = t.FirstChild; t != nil; t = t.NextSibling {
						if t.Type == html.ElementNode && t.Data == "span" {
							tag = append(tag, t.FirstChild.Data)
						}
					}
				}
			}
		}
	}
	return tag
}
