package service

import (
	"context"
	"github.com/aki-yogiri/weather-store/dao"
	pb "github.com/aki-yogiri/weather-store/pb/weather"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type WeatherService struct {
	Database dao.WeatherRepository
}

func (s *WeatherService) GetWeather(ctx context.Context, message *pb.QueryMessage) (*pb.WeatherReply, error) {
	query := &dao.Query{}

	if message.Location == "" {
		log.Println("Error: Location not found")
		return nil, status.Errorf(codes.InvalidArgument, "Location not found")
	}
	query.Location = message.Location

	if message.DatetimeStart != nil {
		dtstart, err := ptypes.Timestamp(message.DatetimeStart)
		if err != nil {
			log.Printf("Error: %v", err)
			return nil, status.Errorf(codes.InvalidArgument, "Invalid timestamp: datetime_start")
		}
		query.DatetimeStart = &dtstart
	}

	if message.DatetimeEnd != nil {
		dtend, err := ptypes.Timestamp(message.DatetimeEnd)
		if err != nil {
			log.Printf("Error: %v", err)
			return nil, status.Errorf(codes.InvalidArgument, "Invalid timestamp: datetime_end")
		}
		query.DatetimeEnd = &dtend
	}

	result, err := s.Database.Find(query)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Could not execute query")
	}

	records := make([]*pb.WeatherMessage, len(result))

	for i, v := range result {
		records[i] = makeWetherMessageProto(&v)
	}

	return &pb.WeatherReply{Weather: records}, nil

}

func (s *WeatherService) PutWeather(ctx context.Context, message *pb.WeatherMessage) (*pb.WeatherReply, error) {
	log.Println("Recieve PutWeather Request: " + message.String())
	var err error
	w := &dao.Weather{}
	w.Location = message.Location
	w.Weather = message.Weather
	w.Temperature = message.Temperature
	w.Clouds = message.Clouds
	w.Wind = message.Wind
	w.WindDeg = message.WindDeg
	w.Timestamp, err = ptypes.Timestamp(message.Timestamp)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid timestamp")
	}

	err = s.Database.Add(w)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, status.Errorf(codes.Aborted, "Could not execute add record")
	}

	return &pb.WeatherReply{Weather: []*pb.WeatherMessage{message}}, nil
}

func makeWetherMessageProto(w *dao.Weather) *pb.WeatherMessage {
	wm := pb.WeatherMessage{}
	wm.Location = w.Location
	wm.Weather = w.Weather
	wm.Temperature = w.Temperature
	wm.Clouds = w.Clouds
	wm.Wind = w.Wind
	wm.WindDeg = w.WindDeg
	wm.Timestamp, _ = ptypes.TimestampProto(w.Timestamp)

	return &wm
}
