package main

import (
	"flag"
	"github.com/ipv4sec/multibully"
	"log"
	"net"
	"os"
)


func main() {
	var iface *string
	var id *string
	iface = flag.String("iface", "eth0", "eth0")
	id = flag.String("id", "default", "id")
	flag.Parse()

	address := "224.0.0.0:9999"
	stop := make(chan struct{})
	pid := uint64(os.Getpid())

	p, err := multibully.NewParticipant(address, *iface, pid, *id, func(state int, ip *net.IP) {
		switch state {
		case multibully.Follower:
			log.Println("* Became Follower of", ip)
		case multibully.Leader:
			log.Println("* Became Leader", ip)
		}
	})

	if err != nil {
		log.Fatal(err)
	}

	go p.StartElection()
	p.RunLoop(stop)
}
