package server

import (
	"github.com/onRuntime/berrygames-strawberry/data"
	"github.com/onRuntime/berrygames-strawberry/router"
	"os"
)

type Server struct {
	Data *data.Data
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start(addr string, devMode bool) error {

	s.Data = data.New()
	err := s.Data.DBConnect(
		os.Getenv("DB_ADDR"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_NAME"),
		"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	err = s.Data.RedisConnect(
		os.Getenv("REDIS_ADDR"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PWD"),
		0)
	if err != nil {
		return err
	}

	if err := s.Data.Init(); err != nil {
		return err
	}

	r := router.New()
	if err := r.Init(addr, s.Data, !devMode); err != nil {
		return err
	}
	return nil
}
