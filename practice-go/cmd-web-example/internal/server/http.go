package server

import (
	"crypto/tls"
	"net"
	"net/http"
	"strconv"
	"time"

	"my-web/internal"
	"my-web/internal/conf"
)

func (s *server) setupAPI() {

	// routes
	r := internal.SetUpRouter()

	// HTTP 挂载
	addr := net.JoinHostPort(conf.ServerHost, strconv.Itoa(conf.ServerPort))
	s.server = &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  time.Duration(1000) * time.Microsecond,
		WriteTimeout: time.Duration(5000) * time.Microsecond,
	}

	// HTTPS 挂载
	if conf.SSLCert != "" && conf.SSLKey != "" {
		addrSSL := net.JoinHostPort(conf.SSLHost, strconv.Itoa(conf.SSLPort))
		s.serverSSL = &http.Server{
			Addr:         addrSSL,
			Handler:      r,
			ReadTimeout:  time.Duration(1000) * time.Microsecond,
			WriteTimeout: time.Duration(5000) * time.Microsecond,
			TLSConfig: &tls.Config{
				// Causes servers to use Go's default ciphersuite preferences,
				// which are tuned to avoid attacks. Does nothing on clients.
				PreferServerCipherSuites: true,
			},
		}
	}
}
