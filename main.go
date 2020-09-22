package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/onRuntime/berrygames-strawberry/server"
	"os"
)

func main() {
	var devMode bool
	flag.BoolVar(&devMode, "devMode", false, "Non-production build")
	flag.Parse()

	if err := godotenv.Load(); err != nil  {
		panic(err)
	}

	s := server.New()
	if err := s.Start(os.Getenv("ROUTE"), devMode); err != nil {
		panic(err)
	}
}
