package cart

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) handlePOSTItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
	}
}

func (s *Service) handlePUTItemByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Service) handleDELETEItemByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Service) handleDELETEItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
