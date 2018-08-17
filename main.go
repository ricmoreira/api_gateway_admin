package main

import (
	"flag"
	"io"
	"log"
	"net/http"
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

	// render that does not change response
	noTransformRender := func(c *gin.Context, response *proxy.Response) {
		if response == nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(response.Metadata.StatusCode)
		for k, v := range response.Metadata.Headers {
			c.Header(k, v[0])
		}
		io.Copy(c.Writer, response.Io)
	}

	// register the render at the router level
	krakendgin.RegisterRender("NoTransformRender", noTransformRender)

	// assign NoTransformRender to all endpoints loaded from config file
	for _, v := range serviceConfig.Endpoints {
		v.OutputEncoding = "NoTransformRender"
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

	backendFactory := func(backendCfg *config.Backend) proxy.Proxy {

		// status handler that does change status
		ns := proxy.NoOpHTTPStatusHandler

		// the default request executor
		re := proxy.DefaultHTTPRequestExecutor(proxy.NewHTTPClient)

		// response parser that copies Backend response body to proxy Response IO reader
		rp := proxy.NoOpHTTPResponseParser

		// build and return the new backend proxy
		return proxy.NewHTTPProxyDetailed(backendCfg, re, ns, rp)
	}

	// build the pipes on top of the custom backend factory
	proxyFactory := proxy.NewDefaultFactory(backendFactory, logger)

	engine := gin.Default()

	routerConfig := krakendgin.Config{
		Engine:         engine,
		ProxyFactory:   proxyFactory,
		Logger:         logger,
		HandlerFactory: krakendgin.EndpointHandler,
		Middlewares: []gin.HandlerFunc{
			cors.New(cors.Config{
				AllowOrigins: []string{"http://localhost:4200", "http://127.0.0.1:4200", "http://localhost:8089", "http://localhost:8069", "http://localhost:8080", "http://localhost:8099"},
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
	}

	routerFactory := krakendgin.NewFactory(routerConfig)

	routerFactory.New().Run(serviceConfig)
}
