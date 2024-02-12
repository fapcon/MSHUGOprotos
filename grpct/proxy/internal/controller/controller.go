package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"proxy/internal/grpcclient"
	"proxy/internal/models"
)

type GeoHandler interface {
	SearchHandler(w http.ResponseWriter, r *http.Request)
	GeocodeHandler(w http.ResponseWriter, r *http.Request)
}

type GeoHandle struct {
	cl client.ClientFactoryRpc
}

func NewGeoHand(cl client.ClientFactoryRpc) *GeoHandle {

	return &GeoHandle{cl: cl}
}

// @Summary Search handler
// @Description Handles search requests
// @ID search-handler
// @Accept json
// @Produce json
// @Param   query     body    models.SearchRequest true        "Search query"
// @Success 200 {object} models.Address
// @Failure 400 {string} string
// @Router /api/address/search [post]
func (h *GeoHandle) SearchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SearchHandler run")
	var req models.SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err read body")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	var result []byte
	address, err := h.cl.CreateClientAndCallSearch(req.Query)
	if err != nil {
		log.Println("err:", err)
	}
	result = address
	w.Write(result)

}

// @Summary Geocode handler
// @Description Handles geocode requests
// @ID geocode-handler
// @Accept json
// @Produce json
// @Param   lat     body    float64 true        "Latitude"
// @Param   lng     body    float64 true        "Longitude"
// @Success 200 {object} models.Address
// @Failure 400 {string} string
// @Router /api/address/geocode [post]
func (h *GeoHandle) GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GeocodeHandler run")
	var req *models.GeocodeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err read body")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	var result []byte
	address, err := h.cl.CreateClientAndCallGeocode(req)
	if err != nil {
		log.Println("err:", err)
	}
	result = address
	w.Write(result)

}
