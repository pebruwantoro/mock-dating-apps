package handler

import "github.com/pebruwantoro/technical_test_dealls/usecase"


type Server struct {
	Usecase usecase.UsecaseInterface
}

type NewUsecaseOptions struct {
	Usecase usecase.UsecaseInterface
}

func NewServer(opts NewUsecaseOptions) *Server {
	return &Server{
		Usecase: opts.Usecase,
	}
}
