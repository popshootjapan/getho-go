package getho

import "errors"

type Getho struct {
	scheme string
	domain string
	host   string
}

type Options struct {
	Scheme string // default is `https://`
	Host   string // default is `getho.io`
}

func New(domain string, opts *Options) (*Getho, error) {
	if len(domain) == 0 {
		return nil, errors.New("domain length is zero.")
	}

	var (
		scheme string
		host   string
	)

	if opts != nil && opts.Scheme != "" {
		scheme = opts.Scheme
	} else {
		scheme = "https://"
	}

	if opts != nil && opts.Host != "" {
		host = opts.Host
	} else {
		host = "getho.io"
	}

	getho := &Getho{
		scheme: scheme,
		domain: domain,
		host:   host,
	}

	return getho, nil
}

func (g *Getho) GetURLString() string {
	return g.scheme + g.domain + "." + g.host
}
