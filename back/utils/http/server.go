package httptools

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MasterYourLife/back/utils/middlewares"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	s      *http.Server
	router *mux.Router

	writeTimeout time.Duration
	readTimeout  time.Duration
	idleTimeout  time.Duration
}

type ServerOption func(*Server)

func NewServer(listenURI string, options ...ServerOption) *Server {
	router := mux.NewRouter()
	router.Use(middlewares.LoggingMiddlewares)

	var server = &Server{
		router: router,

		writeTimeout: 2 * time.Second,
		readTimeout:  2 * time.Second,
		idleTimeout:  2 * time.Second,
	}

	for _, opt := range options {
		opt(server)
	}

	server.s = &http.Server{
		Addr:         listenURI,
		WriteTimeout: server.writeTimeout,
		ReadTimeout:  server.readTimeout,
		IdleTimeout:  server.idleTimeout,
		Handler:      cors.AllowAll().Handler(server.router),
	}

	return server
}

func (this *Server) Routes() ([]string, error) {
	var routes = []string{}
	err := this.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		methods, err := route.GetMethods()
		if err != nil {
			return nil
		}
		routes = append(routes, fmt.Sprintf("[%s] %s", strings.Join(methods, ", "), path))
		return nil
	})
	return routes, err
}

func (this *Server) Start() error {
	log.WithFields(log.Fields{
		"http_bind_address": this.s.Addr,
	}).Info("starting server")
	return this.s.ListenAndServe()
}

func (this *Server) LogRoutes() {
	// print routes
	if routes, err := this.Routes(); err != nil {
		log.WithError(err).Error("cannot get routes")
	} else {
		for _, route := range routes {
			log.Info(route)
		}
	}
}

func (this *Server) Stop(ctx context.Context) {
	log.Info("shutting down server")
	this.s.Shutdown(ctx)
	log.Info("server is down")
}

func (this *Server) GetRouter(pathPrefix string) *mux.Router {
	if pathPrefix != "" {
		return this.router.PathPrefix(pathPrefix).Subrouter()
	} else {
		return this.router
	}
}

func WithPrometheusMetrics() ServerOption {
	return func(s *Server) {
		log.WithFields(log.Fields{
			"path": "/metrics",
		}).Info("with prometheus metrics profiling")
		s.router.Use(middlewares.PrometheusMetricsMiddleware)
		s.router.Handle("/metrics", promhttp.Handler())
	}
}

func WithHealthCheck() ServerOption {
	return func(s *Server) {
		log.WithFields(log.Fields{
			"path": "/health",
		}).Info("with health check")
		s.router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			if err := json.NewEncoder(w).Encode(struct {
				Status bool `json:"status"`
			}{
				Status: true,
			}); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		})
	}
}
