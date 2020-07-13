package main

import (
	"github.com/aki-yogiri/weather-store/dao"
	pb "github.com/aki-yogiri/weather-store/pb/weather"
	"github.com/aki-yogiri/weather-store/service"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type DatabaseEnv struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type ServerEnv struct {
	Port int
}

func main() {
	var dbenv DatabaseEnv
	envconfig.Process("DB", &dbenv)

	db := &dao.WeatherImplPostgres{
		Host:     dbenv.Host,
		Port:     dbenv.Port,
		User:     dbenv.User,
		Password: dbenv.Password,
		DBName:   dbenv.Name,
	}

	err := db.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var serverenv ServerEnv
	envconfig.Process("SERVER", &serverenv)

	listenPort, err := net.Listen("tcp", "0.0.0.0:"+string(serverenv.Port))
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	weatherService := &service.WeatherService{Database: db}
	pb.RegisterWeatherServer(server, weatherService)
	server.Serve(listenPort)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	server.GracefulStop()
}
