package main

import (
	"math/rand"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// GetIPHandler returns all data on a given IP
func (s *Server) GetIPHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := s.log.WithFields(logrus.Fields{"func": "GetIPHandler", "id": rand.Int63()})
		log.WithFields(logrus.Fields{"remote": r.RemoteAddr, "method": r.Method, "path": r.URL.Path, "query": r.URL.Query()}).Infoln()

		vars := mux.Vars(r)
		ip := vars["ip"]

		geoip := s.data.Lookup(ip)

		httpResponse(w, &geoip, http.StatusOK)
	})
}
