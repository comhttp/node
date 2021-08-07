package main

import (
	"github.com/comhttp/jorm/pkg/utl"
)

type BitNode struct {
	IP   string `json:"ip"`
	Port int64  `json:"p"`
	Jrc  *utl.Endpoint
}

// BitNodes is array of bitnodes addresses
type BitNodes []BitNode

// BitNoded data
type BitNoded struct {
	Coin     string          `json:"coin"`
	BitNodes []BitNodeStatus `json:"bitnodes"`
}

// Coin stores identifying information about coins in the database
type NodeCoin struct {
	Rank   int      `json:"r"`
	Name   string   `json:"n"`
	Ticker string   `json:"t"`
	Slug   string   `json:"s"`
	Algo   string   `json:"a"`
	Nodes  BitNodes `json:"b"`
}

type NodeCoins struct {
	N int        `json:"n"`
	C []NodeCoin `json:"c"`
}

// NodeStatus stores current data for a node
type BitNodeStatus struct {
	Live           bool        `json:"live"`
	IP             string      `json:"ip"`
	GetInfo        interface{} `json:"getinfo"`
	GetPeerInfo    interface{} `json:"getpeerinfo"`
	GetRawMemPool  interface{} `json:"getrawmempool"`
	GetMiningInfo  interface{} `json:"getmininginfo"`
	GetNetworkInfo interface{} `json:"getnetworkinfo"`
	GeoIP          interface{} `json:"geoip"`
}

type Nodes []Node

// Node stores info retrieved via geoip about a node
type Node struct {
	IP            string  `json:"ip"`
	Port          int64   `json:"port"`
	Host          string  `json:"host"`
	Rdns          string  `json:"rdns"`
	ASN           string  `json:"asn"`
	ISP           string  `json:"isp"`
	CountryName   string  `json:"countryname"`
	CountryCode   string  `json:"countrycode"`
	RegionName    string  `json:"regionname"`
	RegionCode    string  `json:"regioncode"`
	City          string  `json:"city"`
	Postcode      string  `json:"postcode"`
	ContinentName string  `json:"continentname"`
	ContinentCode string  `json:"continentcode"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Zipcode       string  `json:"zipcode"`
	Timezone      string  `json:"timezone"`
	LastSeen      string  `json:"lastseen"`
	Live          bool    `json:"live"`
}
