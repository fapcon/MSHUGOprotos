package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net/rpc"
	prgen "proxy/internal/grpcclient/protos/gen"
	"proxy/internal/models"
)

type ClientFactoryRpc interface {
	CreateClientAndCallSearch(input string) ([]byte, error)
	CreateClientAndCallGeocode(inp *models.GeocodeRequest) ([]byte, error)
}

type ClientGrpcFactory struct{}

func NewGrpcClientFactory() *ClientGrpcFactory {
	return &ClientGrpcFactory{}
}

func (f *ClientGrpcFactory) CreateClientAndCallSearch(input string) ([]byte, error) {
	conn, err := grpc.Dial("jgrpc:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении к серверу: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := prgen.NewGeoServiceClient(conn)

	req := &prgen.SearchRequest{Input: input}
	res, err := client.Search(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове RPC: %v", err)
		return nil, err
	}
	var result []byte

	result = res.Data

	return result, nil
}

func (f *ClientGrpcFactory) CreateClientAndCallGeocode(inp *models.GeocodeRequest) ([]byte, error) {
	conn, err := grpc.Dial("jgrpc:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении к серверу: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := prgen.NewGeoServiceClient(conn)

	req := &prgen.GeocodeRequest{
		Lat: inp.Lat,
		Lon: inp.Lng,
	}
	res, err := client.Geocode(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове RPC: %v", err)
		return nil, err
	}
	var result []byte

	result = res.Data

	return result, nil
}

type ClientJsonRpcFactory struct{}

func NewJsonRpcClientFactory() *ClientJsonRpcFactory {
	return &ClientJsonRpcFactory{}
}

func (f *ClientJsonRpcFactory) CreateClientAndCallSearch(input string) ([]byte, error) {
	client, err := rpc.DialHTTP("tcp", "jgrpc:50051")
	if err != nil {
		log.Fatal(err)
	}

	var address []byte
	err = client.Call("ServerGeo.Search", input, &address)
	if err != nil {
		log.Fatal(err)
	}

	return address, err
}

func (f *ClientJsonRpcFactory) CreateClientAndCallGeocode(inp *models.GeocodeRequest) ([]byte, error) {
	client, err := rpc.DialHTTP("tcp", "jgrpc:50051")
	if err != nil {
		log.Fatal(err)
	}

	var address []byte
	err = client.Call("ServerGeo.Geocode", inp, &address)
	if err != nil {
		log.Fatal(err)
	}

	return address, err
}

type ClientRpcFactory struct {
}

func NewClientRpcFactory() *ClientRpcFactory {
	return &ClientRpcFactory{}
}

func (f *ClientRpcFactory) CreateClientAndCallSearch(input string) ([]byte, error) {
	client, err := rpc.Dial("tcp", "jgrpc:50051")
	if err != nil {
		log.Fatal("err dial RPC:", err)
		return nil, err
	}
	var address []byte
	err = client.Call("ServerGeo.Search", input, &address)
	if err != nil {
		log.Fatal("err call RPC:", err)
		return nil, err
	}
	return address, nil
}

func (f *ClientRpcFactory) CreateClientAndCallGeocode(inp *models.GeocodeRequest) ([]byte, error) {
	client, err := rpc.Dial("tcp", "jgrpc:50051")
	if err != nil {
		log.Fatal("err dial RPC:", err)
		return nil, err
	}
	var address []byte
	err = client.Call("ServerGeo.Geocode", inp, &address)
	if err != nil {
		log.Fatal("err call RPC:", err)
		return nil, err
	}
	return address, nil
}
