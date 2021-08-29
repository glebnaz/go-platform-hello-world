package main

import (
	"net"

	"github.com/glebnaz/go-platform-hello-world/internal/app/services"
	pb "github.com/glebnaz/go-platform-hello-world/pkg/pb/api/v1"
	"github.com/glebnaz/go-platform/http"
	"github.com/glebnaz/go-platform/metrics"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type app struct {
	//http servers
	debug      *echo.Echo
	http       *echo.Echo
	grpcServer *grpc.Server

	//grpc server

	cfg cfgApp
}

//Start app with grpc and http
func (a app) Start() {
	//todo check error in this function
	a.http.HideBanner = true
	a.debug.HideBanner = true

	//debug
	go a.debug.Start(a.cfg.DebugPort)
	log.Infof("debug server started at port: %s", a.cfg.DebugPort)

	//http
	go a.http.Start(a.cfg.PortHTTP)
	log.Infof("http server started at port: %s", a.cfg.PortHTTP)

	//grpc
	lis, err := net.Listen("tcp", a.cfg.PortGRPC)
	if err != nil {
		log.Fatalf("bad start grpc port: %s", err)
	}

	a.grpcServer.Serve(lis)
}

type cfgApp struct {
	DebugPort string `json:"debug_port" envconfig:"DEBUG_PORT" default:":8084"`
	PortHTTP  string `json:"port_http" envconfig:"PORT_HTTP" default:":8080"`
	PortGRPC  string `json:"port_grpc" envconfig:"PORT_GRPC" default:":8082"`
}

func newApp() app {
	var cfg cfgApp

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("cfg app is bad: %s", err)
	}

	//debug port
	serverDBG := debugServer()

	//server
	server := http.InitRouters(routers)

	//todo добавить опты, скорее всего нужно будет докинуть реализацию interseptors
	var opts []grpc.ServerOption
	//grpc server
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPetStoreServer(grpcServer, services.NewService())

	return app{
		cfg:        cfg,
		debug:      serverDBG,
		http:       server,
		grpcServer: grpcServer,
	}
}

func debugServer() *echo.Echo {
	s := echo.New()

	s.GET("/metrics", echo.WrapHandler(metrics.Handler()))
	return s
}
