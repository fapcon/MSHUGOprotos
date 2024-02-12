package service

import (
	"encoding/json"
	"fmt"
	"io"
	"json-g-rpc/internal/models"
	"log"
	"net/http"
	"strings"
)

type GeoService struct {
}

func (g *GeoService) Search(input string) ([]byte, error) {
	var data = strings.NewReader(fmt.Sprintf("[ \"%s\" ]", input))

	req, err := http.NewRequest("POST", "https://cleaner.dadata.ru/api/v1/clean/address", data)
	if err != nil {
		log.Fatal("dadata err request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+"c4774295a9b0dc55c9b8e73cb299d5c5faeb2da8")
	req.Header.Set("X-Secret", "efefac572c92aaddf6fe54d04a31f00d72702ed1")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("dadata err request:", err)
	}
	defer resp.Body.Close()
	log.Println("dadata statuscode - ", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}

	var adrs models.AddressSearch
	err = json.Unmarshal(body, &adrs)
	if err != nil {
		log.Println("err unmarshal ")
	}

	var pr models.ResponseAddress

	for _, s := range adrs {

		prs := []models.AddressSearchEl{{
			Result: s.Result,
			GeoLat: s.GeoLat,
			GeoLon: s.GeoLon,
		},
		}

		pr.Addresses = prs
	}

	addresses, err := json.Marshal(pr)
	if err != nil {
		log.Println("err marshal :", err)
	}

	return addresses, nil
}

func (g *GeoService) Geocode(lat, lng string) ([]byte, error) {

	var data = strings.NewReader(fmt.Sprintf("{ \"lat\":%s, \"lon\":%s }", lat, lng))
	req, err := http.NewRequest("POST", "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token c4774295a9b0dc55c9b8e73cb299d5c5faeb2da8")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("err read response body", err)
	}
	var adrs models.GeocodeResponse
	err = json.Unmarshal(body, &adrs)
	if err != nil {
		log.Println("err unmarshal ", err)
	}

	var pr models.ResponseGeocode

	for _, v := range adrs.Suggestions {
		as := []models.AddressSearchEl{
			{
				Result: v.Value,
				GeoLat: v.Data.GeoLat,
				GeoLon: v.Data.GeoLon,
			},
		}
		ps := models.ResponseGeocode{
			Value:             v.Value,
			UnrestrictedValue: v.UnrestrictedValue,
			Data:              as,
		}

		pr = ps
	}
	address, err := json.Marshal(pr)
	if err != nil {
		log.Println("err marshal :", err)
	}

	return address, nil
}
