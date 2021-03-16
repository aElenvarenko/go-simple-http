package go_simple_http

import (
	"context"
	"log"
	"net/http"
)

// Config ...
// type Config struct {
// 	Host string
// 	Port string
// }

// Server ...
type Server struct {
	Config     *Config
	HTTPServer *http.Server
	Router     *Router
}

// NewConfig ...
// func NewConfig(host string, port string) *Config {
// 	return &Config{
// 		Host: host,
// 		Port: port,
// 	}
// }

// NewServer ...
func NewServer(config *Config) *Server {
	router := NewRouter()

	return &Server{
		Config: config,
		HTTPServer: &http.Server{
			Addr:    config.Host + ":" + config.Port,
			Handler: router,
		},
		Router: router,
	}
}

// Start ...
func (server *Server) Start() error {
	log.Println("Start listen: " + server.Config.Host + ":" + server.Config.Port)
	return server.HTTPServer.ListenAndServe()
}

// Stop ...
func (server *Server) Stop(ctx context.Context) error {
	log.Println("Stop")
	server.HTTPServer.Shutdown(ctx)
	return nil
}
