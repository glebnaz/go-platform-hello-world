package main

import (
	"github.com/glebnaz/go-platform/http"
	"github.com/glebnaz/go-platform/metrics"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type app struct {
	debug *echo.Echo

	http *echo.Echo
	cfg  cfgApp
}

func (a app) Start() {
	go a.debug.Start(a.cfg.DebugPort)

	a.http.HideBanner = true
	a.debug.HideBanner = true

	a.http.Start(a.cfg.Port)
}

type cfgApp struct {
	DebugPort string `json:"debug_port" envconfig:"DEBUG_PORT" default:"8084"`
	Port      string `json:"port" envconfig:"PORT" default:":8080"`
}

func newApp() app {
	var cfg cfgApp

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("cfg app is bad: %s", err)
	}

	//debug port
	serverDBG := echo.New()
	serverDBG.GET("/metrics", echo.WrapHandler(metrics.Handler()))

	//server
	server := http.InitRouters(routers)

	return app{
		cfg:   cfg,
		debug: serverDBG,
		http:  server,
	}
}
