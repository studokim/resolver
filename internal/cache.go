package internal

import (
	"net"
	"time"
)

type record struct {
	Address net.IP
	Expires time.Time
}

type Cache map[string]record

func (c Cache) contains(domain string) bool {
	host, ok := c[domain]
	return ok && host.Expires.After(time.Now())
}

func (c Cache) get(domain string) net.IP {
	host := c[domain]
	return host.Address
}

func (c Cache) save(domain string, address net.IP, ttl time.Duration) {
	c[domain] = record{Address: address, Expires: time.Now().Add(ttl)}
}
