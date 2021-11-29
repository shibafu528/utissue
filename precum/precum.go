package precum

import (
	"errors"
	"regexp"
)

var registry = []resolverPattern{
	{regexp.MustCompile(".*"), func() Resolver { return NewOGPResolver() }},
}

var (
	// 処理に全く対応していないURLが渡された時のエラー
	ErrUnsupportedUrl = errors.New("unsupported url")
	// resolverがサイトのコンテンツに対応しておらず、他のresolverに委譲する時のエラー
	ErrUnsupportedContent = errors.New("unsupported content")
)

type Material struct {
	Url         string
	Title       string
	Description string
	Image       string
	Tags        []string
}

type Resolver interface {
	Resolve(url string) (*Material, error)
}

type resolverPattern struct {
	pattern *regexp.Regexp
	factory func() Resolver
}

func Resolve(url string) (*Material, error) {
	for _, e := range registry {
		if e.pattern.MatchString(url) {
			// TODO: with timeout
			m, err := e.factory().Resolve(url)
			if errors.Is(err, ErrUnsupportedContent) {
				continue
			}
			return m, err
		}
	}
	return nil, ErrUnsupportedUrl
}
