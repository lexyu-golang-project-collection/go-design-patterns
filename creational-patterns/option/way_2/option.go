package main

import (
	"log"
	"time"
)

type Server struct {
	cfg Config
}

type Config struct {
	host          string
	port          int
	timeout       time.Duration
	maxConnection int
}

func New(cfg Config) *Server {
	return &Server{cfg}
}

func (s *Server) Start() error {
	panic("")
}

/*
Define a new Config struct that holds configuration information
*/
func main() {
	server := New(Config{"localhost", 1234, 30 * time.Second, 10})
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
