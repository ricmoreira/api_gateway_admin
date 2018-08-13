package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/gin-contrib/cors.v1"

	"api_gateway_admin/middleware"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
)

func main() {
	port := flag.Int("p", 0, "Port of the service")
	logLevel := flag.String("l", "ERROR", "Logging level")
	debug := flag.Bool("d", false, "Enable the debug")
	configFile := flag.String("c", "{{path to file}}/configuration.json", "Path to the configuration filename")
	flag.Parse()

	parser := config.NewParser()
	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}

	logger, err := logging.NewLogger(*logLevel, os.Stdout, "[KRAKEND]")
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}

	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine:         gin.Default(),
		ProxyFactory:   customProxyFactory{logger, proxy.DefaultFactory(logger)},
		Logger:         logger,
		HandlerFactory: krakendgin.EndpointHandler,
		Middlewares: []gin.HandlerFunc{
			cors.New(cors.Config{
				AllowOrigins: []string{"http://localhost:4200", "http://127.0.0.1:4200", "http://localhost:8089", "http://localhost:8069", "http://localhost:8080"},
				AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
				AllowHeaders: []string{"Accept",
					"Accept-Encoding",
					"Accept-Language",
					"access-control-allow-origin",
					"Access-Control-Request-Headers",
					"Access-Control-Request-Method",
					"authorization",
					"Cache-Control",
					"Connection",
					"Content-Type",
					"Host",
					"If-Modified-Since",
					"Keep-Alive",
					"Key",
					"Origin",
					"Pragma",
					"User-Agent",
					"X-Custom-Header"},
				ExposeHeaders:    []string{"Content-Length", "Content-Type"},
				AllowCredentials: true,
				MaxAge:           48 * time.Hour,
			}),
			middleware.JwtCheck(),
		},
	})

	routerFactory.New().Run(serviceConfig)
}

// customProxyFactory adds a logging middleware wrapping the internal factory
type customProxyFactory struct {
	logger  logging.Logger
	factory proxy.Factory
}

// New implements the Factory interface
func (cf customProxyFactory) New(cfg *config.EndpointConfig) (p proxy.Proxy, err error) {
	p, err = cf.factory.New(cfg)
	if err == nil {
		p = proxy.NewLoggingMiddleware(cf.logger, cfg.Endpoint)(p)
	}
	return
}
