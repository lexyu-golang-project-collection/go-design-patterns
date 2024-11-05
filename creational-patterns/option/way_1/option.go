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

func New(host string, port int) *Server {
	return &Server{host, port, time.Minute, 100}
}

func NewWithTimeout(host string, port int, timeout time.Duration) *Server {
	return &Server{host, port, timeout, 100}
}

func NewWithTimeoutAndMaxConnection(host string, port int, timeout time.Duration, maxConnection int) *Server {
	return &Server{host, port, timeout, maxConnection}
}

func (s *Server) Start() error {
	panic("")
}

/*
Declare a new constructor for each configuration option
*/
func main() {
	svr := NewWithTimeoutAndMaxConnection("localhost", 1234, 30*time.Second, 10)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
