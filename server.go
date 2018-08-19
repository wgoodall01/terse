package terse // import "github.com/wgoodall01/terse"

import (
	"fmt"
	"net/http"
	"text/tabwriter"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func init() {
	schema := graphql.MustParseSchema(GRAPHQL_SCHEMA, &QueryResolver{})

	r := mux.NewRouter()
	r.Use(WithAppengineContext)

	api := r.PathPrefix("/_/").Subrouter()
	api.Use(WithAuthentication)
	api.Handle("/graphql", &relay.Handler{Schema: schema})
	api.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tw := tabwriter.NewWriter(w, 0, 0, 1, ' ', 0)

		fmt.Fprintf(tw, "App Engine")
		fmt.Fprintf(tw, "\tAppID\t%s\n", appengine.AppID(ctx))
		fmt.Fprintf(tw, "\tDatacenter\t%s\n", appengine.Datacenter(ctx))
		fmt.Fprintf(tw, "\tDefaultVersionHostname\t%s\n", appengine.DefaultVersionHostname(ctx))
		fmt.Fprintf(tw, "\tInstanceID\t%s\n", appengine.InstanceID())
		fmt.Fprintf(tw, "\tIsDevAppServer\t%t\n", appengine.IsDevAppServer())
		fmt.Fprintf(tw, "\tModuleName\t%s\n", appengine.ModuleName(ctx))
		fmt.Fprintf(tw, "\tRequestID\t%s\n", appengine.RequestID(ctx))
		fmt.Fprintf(tw, "\tServerSoftware\t%s\n", appengine.ServerSoftware())
		fmt.Fprintf(tw, "\tVersionID\t%s\n", appengine.VersionID(ctx))

		claims := GetClaims(ctx)
		fmt.Fprintf(tw, "Authentication")
		fmt.Fprintf(tw, "\tAud\t%s\n", claims.Aud)
		fmt.Fprintf(tw, "\tEmail\t%s\n", claims.Email)
		fmt.Fprintf(tw, "\tEmailVerified\t%t\n", claims.EmailVerified)
		fmt.Fprintf(tw, "\tExp\t%d\n", claims.Exp)
		fmt.Fprintf(tw, "\tFamilyName\t%s\n", claims.FamilyName)
		fmt.Fprintf(tw, "\tGivenName\t%s\n", claims.GivenName)
		fmt.Fprintf(tw, "\tIat\t%d\n", claims.Iat)
		fmt.Fprintf(tw, "\tIss\t%s\n", claims.Iss)
		fmt.Fprintf(tw, "\tLocale\t%s\n", claims.Locale)
		fmt.Fprintf(tw, "\tPicture\t%s\n", claims.Picture)
		fmt.Fprintf(tw, "\tSub\t%s\n", claims.Sub)
		tw.Flush()
	})

	r.Handle("/", NewRedirectHandler())

	http.Handle("/", r) // register handler for App Engine
}
