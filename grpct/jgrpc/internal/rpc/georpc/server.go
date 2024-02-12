package georpc

import (
	"json-g-rpc/internal/models"
	"json-g-rpc/internal/service"
	"log"
)

type ServerGeo struct {
	geo service.GeoService
}

func (g *ServerGeo) Search(args string, reply *[]byte) error {
	address, err := g.geo.Search(args)
	if err != nil {
		log.Fatal("err call rpc:", err)
	}
	*reply = address
	return nil
}

func (g *ServerGeo) Geocode(inp *models.GeocodeRequest, reply *[]byte) error {
	address, err := g.geo.Geocode(inp.Lat, inp.Lng)
	if err != nil {
		log.Fatal("err call rpc:", err)
	}
	*reply = address
	return nil
}
