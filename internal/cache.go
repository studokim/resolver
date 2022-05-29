package internal

import (
	"net"
	"time"
)

type host struct {
	Domain  string
	Address net.IP
	Expires time.Time
}

type cache map[string]host

func (c *cache) contains(domain string) bool {
	host, ok := (*c)[domain]
	return ok && host.Expires.After(time.Now())
}

func (c *cache) get(domain string) net.IP {
	host := (*c)[domain]
	return host.Address
}

func (c *cache) save(domain string, address net.IP, ttl time.Duration) {
	(*c)[domain] = host{Domain: domain, Address: address, Expires: time.Now().Add(ttl)}
}
