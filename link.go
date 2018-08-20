package terse

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"google.golang.org/appengine/datastore"
)

type Link struct {
	// Datastore key for entity
	Key *datastore.Key

	// URL shortname.
	Short string

	// Long URL.
	Long string
}

func NewLink(short string, long string) (*Link, error) {
	l := &Link{
		Short: short,
		Long:  long,
	}

	if l.Valid() != nil {
		return nil, errors.New("invalid Link")
	}

	return l, nil
}

var validSlugRegex = regexp.MustCompile("(?i)^[a-z0-9-]+$")

func ValidSlug(slug string) bool {
	return validSlugRegex.MatchString(slug)
}

func (l *Link) Valid() error {
	_, parseErr := url.Parse(l.Long)
	if parseErr != nil {
		return errors.New("invalid long URL")
	}

	l.Short = strings.ToLower(l.Short)
	if !ValidSlug(l.Short) {
		return errors.New("invalid short URL")
	}

	return nil
}

func (l *Link) Load(props []datastore.Property) error {
	if err := datastore.LoadStruct(l, props); err != nil {
		return err
	}

	if l.Valid() != nil {
		return errors.New("invalid Link")
	}

	return nil
}

func (l *Link) Save() ([]datastore.Property, error) {
	if l.Valid() != nil {
		return nil, errors.New("invalid Link")
	}

	return datastore.SaveStruct(l)
}
