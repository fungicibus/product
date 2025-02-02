package v1

import "net/http"

// Create a new order
// (POST /orders)
func (api *API) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateOrder"))
}
