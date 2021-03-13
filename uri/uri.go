//Package uri provides resource access through uniform resource identifiers.
package uri

import (
	"errors"
	"io"
	"strings"
)

//String is a URI string.
//where URI = scheme:[//authority]path[?query][#fragment]
//and authority = [userinfo@]host[:port]
type String string

func (s String) String() string { return string(s) }

//Scheme returns the scheme of the URI string.
func (s String) Scheme() Scheme {
	i := strings.Index(string(s), ":")
	if i == -1 {
		return Scheme(s)
	}
	return Scheme(s[:i])
}

//Scheme of a URI, determines what kind of resource is
//being identified.
type Scheme string

//Opener can open a URI.
type Opener interface {
	Open(uri String) (io.ReadCloser, error)
}

//Open opens the given URI and returns a readable
//result and/or an error.
func Open(uri String) (io.ReadCloser, error) {
	if opener, ok := openers[uri.Scheme()]; ok {
		return opener.Open(uri)
	}
	return nil, errors.New("unsupported scheme")
}

var openers = make(map[Scheme]Opener)

//RegisterOpenerFor registers an Poster for the
//given scheme so that it can be posted with Post.
func RegisterOpenerFor(scheme Scheme, opener Opener) { openers[scheme] = opener }

//Poster can post to a URI.
type Poster interface {
	Post(uri String) (io.WriteCloser, error)
}

//Post initialises a post to the given URI and
//returns a writable stream. The post only
//succeeds if Close returns nil.
func Post(uri String) (io.WriteCloser, error) {
	if poster, ok := posters[uri.Scheme()]; ok {
		return poster.Post(uri)
	}
	return nil, errors.New("unsupported scheme")
}

var posters = make(map[Scheme]Poster)

//RegisterPosterFor registers an Opener for the
//given scheme so that it can be opened with Open.
func RegisterPosterFor(scheme Scheme, poster Poster) { posters[scheme] = poster }
