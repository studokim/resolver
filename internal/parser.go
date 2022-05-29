package internal

import (
	"log"
	"net"
	"time"

	"github.com/miekg/dns"
)

func parseAnswer(reply *dns.Msg) (net.IP, time.Duration) {
	for _, record := range reply.Answer {
		if record.Header().Rrtype == dns.TypeA {
			log.Println("Parsed answer:    ", record)
			return record.(*dns.A).A, time.Duration(record.Header().Ttl)
		}
	}
	return nil, 0
}

func parseAuthority(reply *dns.Msg) string {
	for _, record := range reply.Ns {
		if record.Header().Rrtype == dns.TypeNS {
			log.Println("Parsed authority:", record)
			return record.(*dns.NS).Ns
		}
	}
	return ""
}

func parseAdditional(reply *dns.Msg) net.IP {
	for _, record := range reply.Extra {
		if record.Header().Rrtype == dns.TypeA {
			log.Println("Parsed additional:", record)
			return record.(*dns.A).A
		}
	}
	return nil
}

func parseQuestion(request *dns.Msg) string {
	for _, question := range request.Question {
		if question.Qtype == dns.TypeA {
			return question.Name
		}
	}
	return ""
}
