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

type Env struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func main() {
	var goenv Env
	envconfig.Process("WEATHER_STORE", &goenv)

	db := &dao.WeatherImplPostgres{
		Host:     goenv.Host,
		Port:     goenv.Port,
		User:     goenv.User,
		Password: goenv.Password,
		DBName:   goenv.DBName,
	}

	err := db.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	listenPort, err := net.Listen("tcp", ":19003")
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
