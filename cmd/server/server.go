package main

import (
	"irpc_tcp_example"
	"log"
	"net"
	"strings"
	"time"

	"github.com/marben/irpc"
)

type StringToolImplementation struct{}

func (st StringToolImplementation) Reverse(s string) (string, error) {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	log.Printf("processing strings.Reverse(%q) => %q", s, runes)

	return string(runes), nil
}

func (st StringToolImplementation) Repeat(s string, count int) (string, error) {
	// using std lib function
	res := strings.Repeat(s, count)
	log.Printf("processing strings.Repeat(%q, %d) => %q", s, count, res)

	return res, nil
}

func (st StringToolImplementation) TimeToStr(t time.Time) (string, error) {
	kitchen := t.Format(time.Kitchen) // Kitchen format is defined as "3:04PM"
	log.Printf("formatting %s to %s", t, kitchen)
	return kitchen, nil
}

func main() {
	// stringToolImpl does all the 'real work', but contains no network code whatsoever
	stringToolImpl := StringToolImplementation{}

	// stringToolService of type *irpc_tcp_example.StringToolIRpcService is generated using the irpc tool
	// it translates irpc network calls to the 'real work' implementation and back
	stringToolService := irpc_tcp_example.NewStringToolIRpcService(stringToolImpl)

	// irpc.Server is generic irpc server that operates on a socket (io.ReadWriteCloser)
	// it can register any number of different services and service versions
	irpcServer := irpc.NewServer(stringToolService)

	// open standard tcp listener
	tcpListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error opening tcp listener: %v", err)
	}

	log.Printf("listening on port 8080")

	// Serve() creates irpc endpoint, registers our services and serves remote clients
	if err := irpcServer.Serve(tcpListener); err != nil {
		log.Println("irpcServer.Serve():", err)
	}
}
