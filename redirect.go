package terse

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"google.golang.org/appengine/datastore"
)

func NewRedirectHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		slug := mux.Vars(r)["short"]
		slug = strings.ToLower(slug)

		if !ValidSlug(slug) {
			http.Error(w, "Invalid short link.", 400)
			return
		}

		key := datastore.NewKey(ctx, "Link", slug, 0, nil)
		var link Link
		err := datastore.Get(ctx, key, &link)
		if err != nil {
			http.Error(w, "Couldn't get link from database.", 500)
			return
		}

		http.Redirect(w, r, link.Long, 301)
	})
}
