package pageutil
import(
	"golang.org/x/net/html"
)

//从head中找到视频标题
func FindTitle(head *html.Node) (title string) {
    var scan func(n *html.Node)
    scan = func(n *html.Node){
        if n.Type == html.ElementNode && n.Data == "title" {
            if n.FirstChild != nil {
                title = n.FirstChild.Data
				return 
            }
        }
        for c := n.FirstChild; c!= nil; c = c.NextSibling {
            scan(c) 
        }
    }
    scan(head)
	return title
}