package getho

import "errors"

type Getho struct {
	domain string
}

func New(domain string) (*Getho, error) {
	if len(domain) == 0 {
		return nil, errors.New("domain length is zero.")
	}

	getho := &Getho{
		domain: domain,
	}

	return getho, nil
}
