package main

import (
	"democracy"
)

func main() {
	nc := democracy.NewConfig(
		democracy.WithMaxPacketSize(50),
		democracy.WithPeers([]string{"127.0.0.1:12345"}),
	)
	_, ch := democracy.NewDemocracy(nc)
	<-ch
}

// do it the other way
