package terse

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Link struct {
	Short string
	Long  string
}

func init() {
	http.HandleFunc("/_/create/", dev_addHandler)
	http.HandleFunc("/", dev_indexHandler)
}

func dev_indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("index handler")
	ctx := appengine.NewContext(r)

	short := strings.Split(r.URL.EscapedPath(), "/")[1]
	log.Printf("Getting %s...\n", short)
	key := datastore.NewKey(ctx, "Link", short, 0, nil)
	link := &Link{}
	err := datastore.Get(ctx, key, link)
	if err != nil {
		log.Printf("Error getting URL: %v\n", err)
		return
	}

	log.Printf("%v\n", link)

	http.Redirect(w, r, link.Long, 301)
}

func dev_addHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Creating...\n")
	ctx := appengine.NewContext(r)

	parts := strings.Split(r.URL.EscapedPath(), "/")
	short := parts[3]
	long := parts[4]
	fmt.Fprintf(w, "%s -> %s\n", short, long)

	k := datastore.NewKey(ctx, "Link", short, 0, nil)
	link := Link{
		Short: short,
		Long:  long,
	}

	if _, err := datastore.Put(ctx, k, &link); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}

	fmt.Fprintf(w, "Done. %v", link)
}
