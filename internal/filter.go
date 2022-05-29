package internal

import (
	"net"
	"os"

	"gopkg.in/yaml.v3"
)

type Filter map[string]net.IP

func (f *Filter) contains(domain string) bool {
	_, ok := (*f)[domain]
	return ok
}

func (f *Filter) get(domain string) net.IP {
	address := (*f)[domain]
	return address
}

func (f *Filter) Readconfig() {
	if len(*f) != 0 {
		return
	}
	file, err := os.ReadFile("filter.yml")
	HandleFatal(err)
	err = yaml.Unmarshal(file, f)
	HandleFatal(err)
}
