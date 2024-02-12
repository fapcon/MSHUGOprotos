package main

import (
	"google.golang.org/grpc"
	"json-g-rpc/internal/grpc/geogrpc"
	"json-g-rpc/internal/json-rpc/geojson-rpc"
	"json-g-rpc/internal/rpc/georpc"
	grpcpr "json-g-rpc/protos/gen"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

// @title Proxy Service API
// @version 1.0
// @description This is the API documentation for the Proxy Service.
// @host localhost:8080
// @BasePath /
func main() {
	protocol := os.Getenv("RPC_PROTOCOL")
	switch protocol {
	case "rpc":
		geo := new(georpc.ServerGeo)
		err := rpc.Register(geo)
		if err != nil {
			panic(err)
		}
		l, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatal("Ошибка при запуске сервера:", err)
		}

		log.Println("Сервер запущен на порту 50051")
		rpc.Accept(l)

	case "json-rpc":
		geo := new(geojson_rpc.ServerGeo)
		err := rpc.Register(geo)
		if err != nil {
			panic(err)
		}
		rpc.HandleHTTP()

		l, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatal("Ошибка при запуске сервера:", err)
		}

		log.Println("Сервер запущен на порту 50051")
		http.Serve(l, nil)

	case "grpc":
		listen, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Ошибка при прослушивании порта: %v", err)
		}

		server := grpc.NewServer()
		grpcpr.RegisterGeoServiceServer(server, &geogrpc.ServiceGeo{})

		log.Println("Запуск gRPC сервера...")
		if err := server.Serve(listen); err != nil {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
	default:
		log.Println("unknown protocol JGRPC")
		return
	}
}
