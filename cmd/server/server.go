package main

import (
	"irpc_tcp_example"
	"log"
	"net"

	"github.com/marben/irpc"
)

func main() {
	// backendImpl does all the 'real work', but contains no network code whatsoever
	backendImpl := BackendImplementation{}

	// backendService of type *irpc_tcp_example.BackendIRpcService is generated using the irpc tool
	// it translates irpc network calls to the 'real work' implementation and back
	backendService := irpc_tcp_example.NewBackendIRpcService(backendImpl)

	// irpc.Server is generic irpc server that operates on a socket (io.ReadWriteCloser)
	// it can register any number of different services and service versions
	irpcServer := irpc.NewServer(backendService)

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
