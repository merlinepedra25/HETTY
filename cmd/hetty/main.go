package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/dstotijn/hetty/db/cayley"
	"github.com/dstotijn/hetty/pkg/api"
	"github.com/dstotijn/hetty/pkg/proxy"
	"github.com/dstotijn/hetty/pkg/reqlog"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

var (
	caCertFile string
	caKeyFile  string
	dbFile     string
	addr       string
	adminPath  string
)

func main() {
	flag.StringVar(&caCertFile, "cert", "", "CA certificate file path")
	flag.StringVar(&caKeyFile, "key", "", "CA private key file path")
	flag.StringVar(&dbFile, "db", "hetty.db", "Database file path")
	flag.StringVar(&addr, "addr", ":80", "TCP address to listen on, in the form \"host:port\"")
	flag.StringVar(&adminPath, "adminPath", "", "File path to admin build")
	flag.Parse()

	tlsCA, err := tls.LoadX509KeyPair(caCertFile, caKeyFile)
	if err != nil {
		log.Fatalf("[FATAL] Could not load CA key pair: %v", err)
	}

	caCert, err := x509.ParseCertificate(tlsCA.Certificate[0])
	if err != nil {
		log.Fatalf("[FATAL] Could not parse CA: %v", err)
	}

	db, err := cayley.NewDatabase(dbFile)
	if err != nil {
		log.Fatalf("[FATAL] Could not initialize database: %v", err)
	}
	defer db.Close()

	reqLogService := reqlog.NewService(db)

	p, err := proxy.NewProxy(caCert, tlsCA.PrivateKey)
	if err != nil {
		log.Fatalf("[FATAL] Could not create Proxy: %v", err)
	}

	p.UseRequestModifier(reqLogService.RequestModifier)
	p.UseResponseModifier(reqLogService.ResponseModifier)

	var adminHandler http.Handler
	if adminPath == "" {
		// Used for embedding with `rice`.
		box, err := rice.FindBox("../../admin/dist")
		if err != nil {
			log.Fatalf("[FATAL] Could not find embedded admin resources: %v", err)
		}
		adminHandler = http.FileServer(box.HTTPBox())
	} else {
		adminHandler = http.FileServer(http.Dir(adminPath))
	}

	router := mux.NewRouter().SkipClean(true)

	adminRouter := router.MatcherFunc(func(req *http.Request, match *mux.RouteMatch) bool {
		hostname, _ := os.Hostname()
		host, _, _ := net.SplitHostPort(req.Host)
		return strings.EqualFold(host, hostname) || (req.Host == "hetty.proxy" || req.Host == "localhost:8080")
	}).Subrouter()

	// GraphQL server.
	adminRouter.Path("/api/playground").Handler(playground.Handler("GraphQL Playground", "/api/graphql"))
	adminRouter.Path("/api/graphql").Handler(handler.NewDefaultServer(api.NewExecutableSchema(api.Config{Resolvers: &api.Resolver{
		RequestLogService: reqLogService,
	}})))

	// Admin interface.
	adminRouter.PathPrefix("").Handler(adminHandler)

	// Fallback (default) is the Proxy handler.
	router.PathPrefix("").Handler(p)

	s := &http.Server{
		Addr:         addr,
		Handler:      router,
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){}, // Disable HTTP/2
	}

	log.Printf("[INFO] Running server on %v ...", addr)
	err = s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("[FATAL] HTTP server closed: %v", err)
	}
}