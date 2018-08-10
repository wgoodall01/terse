package terse // import "github.com/wgoodall01/terse"

import (
	"context"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func init() {
	ctx := context.Background()

	schema := graphql.MustParseSchema(GRAPHQL_SCHEMA, &QueryResolver{})
	http.Handle("/_/graphql", WithContext(&relay.Handler{Schema: schema}, ctx))
	http.Handle("/", WithContext(NewRedirectHandler(), ctx))
}
