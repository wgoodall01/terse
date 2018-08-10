package terse

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	graphql "github.com/graph-gophers/graphql-go"
)

type QueryResolver struct{}

type LinkQueryArgs struct {
	Limit int
	Short *string
	Long  *string
}

func (q *QueryResolver) Link(ctx context.Context, args LinkQueryArgs) ([]LinkResolver, error) {
	query := datastore.NewQuery("Link")

	if args.Limit != 0 {
		query = query.Limit(args.Limit)
	}

	if args.Short != nil {
		query = query.Filter("Short =", *args.Short)
	}

	if args.Long != nil {
		query = query.Filter("Long =", *args.Long)
	}

	links := make([]LinkResolver, 0)
	for iter := query.Run(ctx); ; {
		var linkResolver LinkResolver
		linkResolver.link = &Link{}
		key, err := iter.Next(linkResolver.link)

		if err == datastore.Done {
			break
		}
		if err != nil {
			log.Errorf(ctx, "Couldn't query links (key=%v):%v", key, err)
			return nil, errors.Wrap(err, "couldn't query Links")
		}

		links = append(links, linkResolver)
	}

	return links, nil
}

type CreateLinkQueryArgs struct {
	Link struct {
		Long  string
		Short string
	}
}

func (q *QueryResolver) CreateLink(ctx context.Context, args CreateLinkQueryArgs) (*LinkResolver, error) {
	link, validErr := NewLink(args.Link.Short, args.Link.Long)
	if validErr != nil {
		return nil, errors.Wrap(validErr, "invalid Link")
	}

	key := datastore.NewKey(ctx, "Link", link.Short, 0, nil)
	_, err := datastore.Put(ctx, key, link)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't save entity")
	}

	log.Infof(ctx, "Created link: %v", link)
	return &LinkResolver{link: link}, nil
}

type LinkResolver struct {
	link *Link
}

func (l LinkResolver) ID() graphql.ID {
	return graphql.ID(l.link.Key.Encode())
}

func (l LinkResolver) Long() string {
	return l.link.Long
}

func (l LinkResolver) Short() string {
	return l.link.Short
}
