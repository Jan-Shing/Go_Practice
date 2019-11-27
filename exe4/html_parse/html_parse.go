package main

import(
	"os"
	"golang.org/x/net/html"
	"flag"
	"log"
	"fmt"
	"strings"
)

type Link struct{
	href_c string
	text_c string
}

func main(){
	var links []Link
	file_name := flag.String("file", "ex1.html", "the html file to used to parsed text inside href label")
	flag.Parse()

	f, err := os.Open(*file_name); 
	if err != nil{
		panic(err)
	}

	doc, err2 := html.Parse(f); 
	if err2 != nil{
		log.Fatal(err2)
	}

	var fn func(*html.Node)
	fn = func(n *html.Node){
		var link Link
	/*									  <tag name>	*/
		if n.Type == html.ElementNode && n.Data == "a"{
			for _, a := range n.Attr{
				if a.Key == "href"{
					link.href_c = a.Val
					/* find the child contain text*/
					for id := n.FirstChild; id != nil; id=id.NextSibling{
						if n.FirstChild.Type == html.TextNode{
							link.text_c += n.FirstChild.Data
						}
						
					}
					link.text_c = strings.Join(strings.Fields(link.text_c), " ")
					links = append(links, link)
					break
				}
			}
		}
		/* DFS search */
		for c := n.FirstChild; c != nil; c = c.NextSibling{
			fn(c)
		}
	}

	fn(doc)
	fmt.Println(links)

}
