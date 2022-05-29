package internal

import (
	"errors"
	"log"
	"net"

	"github.com/miekg/dns"
)

const root = "198.41.0.4"

type Resolver struct {
	cache  Cache
	filter Filter
}

func (r *Resolver) Init() {
	r.cache = make(Cache)
	r.filter = make(Filter)
	r.filter.Readconfig()
}

func (r *Resolver) resolveLocally(domain string) net.IP {
	var address net.IP
	if r.filter.contains(domain) {
		address = r.filter.get(domain)
	}
	if r.cache.contains(domain) {
		address = r.cache.get(domain)
	}
	if address != nil {
		log.Println("Resolved locally:", domain, "IN A", address)
	}
	return address
}

func (r *Resolver) resolveRemotely(domain string) (net.IP, error) {
	nameserver := net.ParseIP(root)
	request := new(dns.Msg)
	request.SetQuestion(domain, dns.TypeA)
	for {
		reply, err := dns.Exchange(request, nameserver.String()+":53")
		if err != nil {
			return nil, err
		}
		if address, ttl := parseAnswer(reply); address != nil {
			r.cache.save(domain, address, ttl)
			return address, nil
		} else if nsIp := parseAdditional(reply); nsIp != nil {
			nameserver = nsIp
		} else if nsDomain := parseAuthority(reply); nsDomain != "" {
			nameserver = r.Resolve(nsDomain)
		} else {
			return nil, errors.New(domain + " does not have any A record")
		}
	}
}

func (r *Resolver) Resolve(domain string) net.IP {
	domain = dns.Fqdn(domain)
	address := r.resolveLocally(domain)
	if address != nil {
		return address
	}
	address, err := r.resolveRemotely(domain)
	Handle(err)
	return address
}
