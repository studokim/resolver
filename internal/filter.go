package internal

import (
	"net"
	"os"

	"gopkg.in/yaml.v3"
)

type filter map[string]net.IP

func (f *filter) contains(domain string) bool {
	_, ok := (*f)[domain]
	return ok
}

func (f *filter) get(domain string) net.IP {
	address := (*f)[domain]
	return address
}

func (f *filter) readconfig() {
	file, err := os.ReadFile("filter.yml")
	Handle(err)
	err = yaml.Unmarshal(file, f)
	Handle(err)
}
