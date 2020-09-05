package cart

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Service Service
type Service struct {
	router *mux.Router
	DB     *gorm.DB
}

func (s *Service) routes() {
	s.router.HandleFunc("/cart", s.handlePOSTCart()).Methods("POST")
	s.router.HandleFunc("/cart/{cartID:[0-9]+}", s.handleGETCart()).Methods("GET")
	s.router.HandleFunc("/cart/{cartID:[0-9]+}/item", s.handlePOSTItem()).Methods("POST")
	s.router.HandleFunc("/cart/{cartID:[0-9]+}/item", s.handleDELETEItem()).Methods("DELETE")
	s.router.HandleFunc("/cart/{cartID:[0-9]+}/item/{itemID:[0-9]+}", s.handlePUTItemByID()).Methods("PUT")
	s.router.HandleFunc("/cart/{cartID:[0-9]+}/item/{itemID:[0-9]+}", s.handleDELETEItemByID()).Methods("DELETE")
	s.router.PathPrefix("/").Handler(s.handle404()).Methods("GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH")
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Service) handle404() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ErrorNotFound(w, "")
	}
}

// NewService  Creates a new service.
func NewService() *Service {
	s := &Service{}
	s.router = mux.NewRouter()

	s.routes()

	s.router.Use(loggingMiddleware)
	s.router.Use(jsonMiddleware)
	s.router.Use(handlers.RecoveryHandler())
	s.router.Use(corsMiddleware)
	s.router.Use(mux.CORSMethodMiddleware(s.router))
	s.router.Use(authMiddleware)
	return s
}
