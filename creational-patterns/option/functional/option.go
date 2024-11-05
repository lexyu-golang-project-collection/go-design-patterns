package main

import (
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
	panic("")
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
	svr := New(
		WithHost("localhost"),
		WithPort(8080),
		WithTimeout(time.Minute),
		WithMaxConn(120),
	)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
