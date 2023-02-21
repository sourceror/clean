package rpc

import (
	"errors"
	"net"

	"google.golang.org/grpc"
)

var (
	ErrEmptyAddress = errors.New("empty address")
	ErrNilBooksApp  = errors.New("nil BooksApp")
)

type ServerParameters struct {
	Address  string
	BooksApp BooksApp
}

type BooksApp interface {
	NewBook
}

func (sp *ServerParameters) validate() error {
	if sp.Address == "" {
		return ErrEmptyAddress
	}

	if sp.BooksApp == nil {
		return ErrNilBooksApp
	}

	return nil
}

type Server struct {
	address        string
	newBookHandler *newBookHandler
}

func NewServer(params ServerParameters) (*Server, error) {
	err := params.validate()
	if err != nil {
		return nil, err
	}

	newBookHandler, err := createNewBookHandler(params.BooksApp)
	if err != nil {
		return nil, err
	}
	return &Server{
		address:        params.Address,
		newBookHandler: newBookHandler,
	}, nil
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	// add services

	return srv.Serve(listener)
}
