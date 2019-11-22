package main
/* copy urlshort folder into /root/go/src/ or 
	/usr/local/go/src/
*/
import(
	"fmt"
	"net/http"
	"./urlshort"
)

func main(){
	// type ServeMux
	mux := defaultMux()

	/* Build the map database */
	pathsToUrls := map[string]string{
		"/urlshort-godoc" : "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc" : 	"https://godoc.org/gopkg.in/yaml.v2",
	}
	// type Handler interface
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

    /* Add new url mapping in yaml format*/
yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)

	if err != nil{
		panic(err)
	}
	fmt.Println("Starting the server on : 8080")
	http.ListenAndServe(":8080", yamlHandler)


}

func defaultMux() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r * http.Request){
	fmt.Fprintln(w, " Helloo go lang user!")
}