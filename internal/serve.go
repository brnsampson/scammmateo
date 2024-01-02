package internal

import (
	"os"
	"crypto/tls"
	"net/http"
	"github.com/brnsampson/scammmateo/pkg/server"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/charmbracelet/log"
	"github.com/spf13/pflag"
)

var static = http.FileServer(http.Dir("./static"))
//var index_templ = template.Must(template.ParseFiles("./templates/index.html"))

func routeRoot(r chi.Router) {
	r.Mount("/", static)
}

// Struct used as primary entrypoint for an RPC based interface.
func NewScamServer(flags *pflag.FlagSet) (*ScamServer, error) {
    log.SetLevel(log.DebugLevel)
    logger := log.NewWithOptions(os.Stderr, log.Options{
        ReportCaller: true,
        ReportTimestamp: true,
        Prefix: "golang_example_service",
        Level: log.DebugLevel,
    })

    conf := &ScamConfig{ }

    router := chi.NewRouter()
    router.Use(middleware.Logger)
    router.Use(middleware.Recoverer)

    router.Route("/", routeRoot)

	return &ScamServer{
		router,
		server.NewServer(logger),
		logger,
		conf,
	}, nil
}

type ScamConfig struct { }

func (c *ScamConfig) GetAddr() (string, error) {
	return "127.0.0.1:1443", nil
}

func (c *ScamConfig) GetTlsConfig() *tls.Config {
    tlsConf := tls.Config{
		Certificates:             []tls.Certificate{},
		MinVersion:               tls.VersionTLS13,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	return &tlsConf
}

func (c *ScamConfig) GetTlsEnabled() bool {
	return false
}

type ScamServer struct {
    router *chi.Mux
	server *server.Server
	logger *log.Logger
	config *ScamConfig
}

func (ss *ScamServer) Run() {
    // Should never produce an error since we are skipping the update
	ss.server.Run(ss.router, ss.config, func() error { return nil })
}

func (ss *ScamServer) BlockingRun() int {
    // Should never produce an error since we are skipping the update
    return ss.server.BlockingRun(ss.router, ss.config, func() error { return nil })

}
