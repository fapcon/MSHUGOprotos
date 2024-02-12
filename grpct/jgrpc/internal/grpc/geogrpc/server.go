package geogrpc

import (
	"context"
	"fmt"
	"json-g-rpc/internal/service"
	grpcpr "json-g-rpc/protos/gen"
)

type Geo interface {
	Search(input string) ([]byte, error)
	Geocode(lat, lng string) ([]byte, error)
}

type ServiceGeo struct {
	grpcpr.UnimplementedGeoServiceServer
	geo service.GeoService
}

func (g *ServiceGeo) Search(ctx context.Context, request *grpcpr.SearchRequest) (*grpcpr.SearchResponse, error) {

	address, err := g.geo.Search(request.Input)
	if err != nil {
		return nil, fmt.Errorf("err get address:%v", err)
	}

	return &grpcpr.SearchResponse{Data: address}, nil
}

func (g *ServiceGeo) Geocode(ctx context.Context, req *grpcpr.GeocodeRequest) (*grpcpr.GeocodeResponse, error) {
	address, err := g.geo.Geocode(req.Lat, req.Lon)
	if err != nil {
		return nil, fmt.Errorf("err get address:%v", err)
	}

	return &grpcpr.GeocodeResponse{Data: address}, nil
}
