package internal

import (
	"net"
	"os"

	"github.com/miekg/dns"
	"gopkg.in/yaml.v3"
)

type Filter map[string]net.IP

func (f Filter) contains(domain string) bool {
	_, ok := f[domain]
	return ok
}

func (f Filter) get(domain string) net.IP {
	address := f[domain]
	return address
}

func (f Filter) Readconfig() {
	if len(f) != 0 {
		return
	}
	file, err := os.ReadFile("filter.yml")
	HandleFatal(err)
	temp := &Filter{}
	err = yaml.Unmarshal(file, temp)
	HandleFatal(err)
	for key, val := range *temp {
		f[dns.Fqdn(key)] = val
	}
}
