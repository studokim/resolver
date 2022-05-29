package internal

import (
	"errors"
	"net"
)

func resolveLocally(domain string, c *cache, f *filter) (net.IP, error) {
	if f.contains(domain) {
		return f.get(domain), nil
	}
	if c.contains(domain) {
		return c.get(domain), nil
	}
	return nil, errors.New(domain + " not available locally")
}

func Resolve(domain string) net.IP {
	f := make(filter)
	f.readconfig()
	c := make(cache)
	address, err := resolveLocally(domain, &c, &f)
	Handle(err)

	return address
}
