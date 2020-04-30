package common

import (
	"errors"
	"github.com/oschwald/geoip2-golang"
	"net"
)

func GeoInfo(ipStr string) (string, string, string, string, string, error) {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		errMsg := "open geo lite to city failed"
		Logger.Error(errMsg)
		return "", "", "", "", "", errors.New(errMsg)
	}
	defer db.Close()

	ip := net.ParseIP(ipStr)
	record, err := db.City(ip)
	if err != nil {
		errMsg := "get result from  geo lite to city failed"
		Logger.Error(errMsg)
		return "", "", "", "", "", errors.New(errMsg)
	}

	city := record.City.Names["en"]
	subdivision := ""
	if len(record.Subdivisions) > 0 {
		subdivision = record.Subdivisions[0].Names["en"]
	}
	country := record.Country.Names["en"]

	longitude := FloatToString(record.Location.Longitude)
	latitude := FloatToString(record.Location.Latitude)

	return city, subdivision, country, longitude, latitude, nil
}
