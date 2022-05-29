package internal

import (
	"fmt"
	"log"
	"strconv"

	"github.com/miekg/dns"
)

type serverHandler struct {
	resolver Resolver
}

func (h *serverHandler) ServeDNS(writer dns.ResponseWriter, request *dns.Msg) {
	reply := new(dns.Msg)
	reply.SetReply(request)
	domain := parseQuestion(request)
	address := h.resolver.Resolve(domain)
	if address != nil {
		rr, err := dns.NewRR(fmt.Sprintf("%s A %s", domain, address))
		if err == nil {
			reply.Answer = append(request.Answer, rr)
		}
	} else {
		reply.SetRcode(request, dns.RcodeNameError)
	}
	writer.WriteMsg(reply)
}

func Listen(port int) {
	var handler serverHandler
	handler.resolver.Init()

	server := dns.Server{Addr: ":" + strconv.Itoa(port), Net: "udp"}
	server.Handler = &handler
	log.Println("listening at", port)
	err := server.ListenAndServe()
	if err != nil {
		HandleFatal(err)
	}
}
