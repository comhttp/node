package main

import (
	"encoding/json"
	"github.com/comhttp/jorm/pkg/utl"
	"io/ioutil"
	"net/http"
)

type GeoResponse struct {
	Status      string
	Description string
	Data        struct {
		Geo struct {
			Host          string  `json:"host"`
			IP            string  `json:"ip"`
			RDSN          string  `json:"rdns"`
			ASN           float64 `json:"asn"`
			ISP           string  `json:"isp"`
			CountryName   string  `json:"country_name"`
			CountryCode   string  `json:"country_code"`
			RegionName    string  `json:"region_name"`
			RegionCode    string  `json:"region_code"`
			City          string  `json:"city"`
			PostalCode    string  `json:"postal_code"`
			ContinentName string  `json:"continent_name"`
			ContinentCode string  `json:"continent_code"`
			Latitude      float64 `json:"latitude"`
			Longitude     float64 `json:"longitude"`
			MetroCode     string  `json:"metro_code"`
			Timezone      string  `json:"timezone"`
			Datetime      string  `json:"datetime"`
		}
	}
}

func GetGeoIP(ip string) (n Node) {
	//if jdb.JDB.Read(cfg.C.Out+"/geo", ip, &n) != nil {
	if ip[:3] == "10." {
		ip = "212.62.35.158"
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://tools.keycdn.com/geo.json?host="+ip, nil)
	if err != nil {
		utl.ErrorLog(err)
	}
	if req != nil {
		req.Header.Set("User-Agent", "keycdn-tools:https://com-http.us")
		resp, err := client.Do(req)
		if err != nil {
			utl.ErrorLog(err)
		}
		defer resp.Body.Close()
		mapBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			utl.ErrorLog(err)
		}
		var g GeoResponse
		err = json.Unmarshal(mapBody, &g)
		if err != nil {
		}
		geo := g.Data.Geo
		n.IP = ip
		n.Rdns = geo.RDSN
		n.ISP = geo.ISP
		n.CountryName = geo.CountryName
		n.CountryCode = geo.CountryCode
		n.RegionName = geo.RegionName
		n.RegionCode = geo.RegionCode
		n.City = geo.City
		n.Zipcode = geo.PostalCode
		n.ContinentName = geo.ContinentName
		n.ContinentCode = geo.ContinentCode
		n.Latitude = geo.Latitude
		n.Longitude = geo.Longitude
		n.Postcode = geo.PostalCode
		n.Timezone = geo.Timezone
		//jdb.JDB.Write(cfg.C.Out+"/geo", ip, n)
	}
	return n
}
