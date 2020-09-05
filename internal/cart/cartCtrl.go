package cart

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) handlePOSTCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _, ok := r.BasicAuth()

		// We know that is a valid user if he reached this far, but just in
		// case....
		if !ok {
			ErrorUnauthorised(w, "")
			return
		}

		ID, err := GetUserID(user)

		if err != nil {
			ErrorUnprocessable(w, "user ID not found")
		}

		newCart := Cart{
			UserID: ID,
		}

		ret := s.DB.Create(&newCart)
		if ret.Error != nil {
			ErrorBadRequest(w, ret.Error.Error())
			return
		}

		res := &CartResponse{
			ID:        newCart.ID,
			CreatedAt: newCart.CreatedAt,
			UpdatedAt: newCart.UpdatedAt,
			UserID:    newCart.UserID,
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(res)
	}
}

func (s *Service) handleGETCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, _, ok := r.BasicAuth()

		// We know that is a valid user if he reached this far, but just in
		// case....
		if !ok {
			ErrorUnauthorised(w, "")
			return
		}

		ID, err := GetUserID(user)

		if err != nil {
			ErrorUnprocessable(w, "user ID not found")
		}

		vars := mux.Vars(r)

		cartID, ok := vars["cartID"]

		if !ok {
			ErrorNotFound(w, "")
			return
		}

		var cart Cart
		ret := s.DB.Where("user_id = ?", ID).First(&cart, cartID)

		if ret.Error != nil {
			ErrorNotFound(w, "")
			return
		}

		res := &CartResponse{
			ID:        cart.ID,
			CreatedAt: cart.CreatedAt,
			UpdatedAt: cart.UpdatedAt,
			UserID:    cart.UserID,
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	}
}
