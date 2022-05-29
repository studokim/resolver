package internal

import (
	"errors"
	"net"
	"time"
)

const root = "198.41.0.4"
const port = 53

func resolveLocally(domain string, c *Cache, f *Filter) (net.IP, error) {
	if f.contains(domain) {
		return f.get(domain), nil
	}
	if c.contains(domain) {
		return c.get(domain), nil
	}
	return nil, errors.New(domain + " not available locally")
}

func resolveRemotely(domain string, nameserver string, c *Cache) (net.IP, error) {
	address := net.ParseIP("0.0.0.0")
	c.save(domain, address, time.Hour)
	return address, nil
}

func Resolve(domain string, c *Cache, f *Filter) net.IP {
	address, err := resolveLocally(domain, c, f)
	if err == nil {
		return address
	}
	address, err = resolveRemotely(domain, root, c)
	Handle(err)
	return address
}
