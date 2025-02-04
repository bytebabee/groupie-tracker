package utilities

import (
	"fmt"
	"net"
	"net/http"
)

func ListenPort() (net.Listener, int, error) {
	//create any port address
	listener, err := net.Listen("tcp", "localhost:0")

	if err != nil {
		return nil, 0, fmt.Errorf("failed to listen on dynamic port: %v", err)
	}
	//define the port address, IP + random port number
	address := listener.Addr().(*net.TCPAddr)
	port := address.Port

	return listener, port, nil
}

// after creating and definin the port, serve to that port
func Serve(listener net.Listener, mux *http.ServeMux) error {
	server := &http.Server{
		Handler: mux,
	}
	return server.Serve(listener)
}
