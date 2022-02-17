package server

import (
	"fmt"
	"net/http"
	"os"

	"my-web/cmd"
	"my-web/internal/conf"
	"my-web/internal/log"
)

// server is the manager of Manager API, which is responsible for managing the life cycle of Manager API
// including initialization, start, stop and so on
type server struct {
	server    *http.Server
	serverSSL *http.Server
	options   *Options
}

type Options struct{}

// NewServer create a server manager
func NewServer(options *Options) (*server, error) {
	return &server{options: options}, nil
}

func (s *server) Start(errSig chan error) {
	// initialize server
	err := s.init()
	if err != nil {
		errSig <- err
		return
	}
}

func (s *server) init() error {
	log.Info("Initialize Manager API server")
	s.setupAPI()

	return nil
}

func (s *server) printInfo() {
	fmt.Fprint(os.Stdout, "Themanager-api is running successfully! \n\n")
	cmd.PrintVersion()
	fmt.Fprintf(os.Stdout, "%-8s: %s:%d\n", "Listen", conf.ServerHost, conf.ServerPort)
	if conf.SSLCert != "" && conf.SSLKey != "" {
		fmt.Fprintf(os.Stdout, "%-8s: %s:%d\n", "HTTPS LIsten", conf.SSLHost, conf.SSLPort)
	}
	fmt.Fprintf(os.Stdout, "%-8s: %s\n", "Loglevel", conf.ErrorLogLevel)
	fmt.Fprintf(os.Stdout, "%-8s: %s\n", "ErrorLogFile", conf.ErrorLogPath)
	fmt.Fprintf(os.Stdout, "%-8s: %s\n\n", "AccessLogFile", conf.AccessLogPath)
}
