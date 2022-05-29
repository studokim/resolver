package internal

import (
	"errors"
	"net"

	"github.com/miekg/dns"
)

const root = "198.41.0.4"

func resolveLocally(domain string, c *Cache, f *Filter) net.IP {
	if f.contains(domain) {
		return f.get(domain)
	}
	if c.contains(domain) {
		return c.get(domain)
	}
	return nil
}

func resolveRemotely(domain string, c *Cache, f *Filter) (net.IP, error) {
	nameserver := net.ParseIP(root)
	request := new(dns.Msg)
	request.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	for {
		reply, err := dns.Exchange(request, nameserver.String()+":53")
		if err != nil {
			return nil, err
		}
		if address, ttl := parseAnswer(reply); address != nil {
			c.save(domain, address, ttl)
			return address, nil
		} else if nsIp := parseAdditional(reply); nsIp != nil {
			nameserver = nsIp
		} else if nsDomain := parseAuthority(reply); nsDomain != "" {
			nameserver = Resolve(nsDomain, c, f)
		} else {
			return nil, errors.New(domain + " does not have any A record")
		}
	}
}

func Resolve(domain string, c *Cache, f *Filter) net.IP {
	address := resolveLocally(domain, c, f)
	if address != nil {
		return address
	}
	address, err := resolveRemotely(domain, c, f)
	Handle(err)
	return address
}
