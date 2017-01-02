package main

import (
	"regexp"
	"strconv"

	"github.com/abh/geoip"
)

// GeoIPServer manages the geoip data
type GeoIPServer struct {
	country *geoip.GeoIP
	asn     *geoip.GeoIP
}

// NewGeoIPServer creates a new geoip database
func NewGeoIPServer() *GeoIPServer {
	return &GeoIPServer{}
}

// LoadData loads the geoip data files
func (g *GeoIPServer) LoadData(country string, asn string) error {
	var err error
	if g.country, err = geoip.Open(country); err != nil {
		return err
	}
	if g.asn, err = geoip.Open(asn); err != nil {
		return err
	}
	return nil
}

// GeoIP contains geoip details of an IP address
type GeoIP struct {
	Country string `json:"country"`
	ASN     int    `json:"asn"`
}

// Lookup returns what's known about the given IP
func (g *GeoIPServer) Lookup(ip string) GeoIP {
	country, _ := g.country.GetCountry(ip)
	asnRaw, _ := g.asn.GetName(ip)
	asn, _ := strconv.Atoi(string(regexp.MustCompile(`(\d+)`).Find([]byte(asnRaw))))
	return GeoIP{Country: country, ASN: asn}
}
