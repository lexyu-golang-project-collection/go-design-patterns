package main

import (
	"fmt"
	"log"
	"time"
)

type Server struct {
	host          string
	port          int
	timeout       time.Duration
	maxConnection int
}

func New(options ...func(*Server)) *Server {
	// Default Values
	server := &Server{
		host:          "127.0.0.1",
		port:          80,
		timeout:       time.Second * 30,
		maxConnection: 100,
	}

	// 應用傳入的 options
	for _, option := range options {
		option(server)
	}

	return server
}

func WithHost(host string) func(*Server) {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) func(*Server) {
	return func(s *Server) {
		s.maxConnection = maxConn
	}
}

func (s *Server) Start() error {
	panic("")
}

/*
Functional Option Pattern
*/
func main() {
	defaultSvr := New()
	fmt.Println(defaultSvr)

	svr := New(
		WithHost("localhost"),
		WithPort(8080),
		WithTimeout(time.Minute),
		WithMaxConn(200),
	)
	fmt.Println(svr)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
