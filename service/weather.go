package service

import (
	"context"
	"errors"
	"github.com/aki-yogiri/weather-store/dao"
	pb "github.com/aki-yogiri/weather-store/pb/weather"
	"github.com/golang/protobuf/ptypes"
	"log"
)

type WeatherService struct {
	Database dao.WeatherRepository
}

func (s *WeatherService) GetWeather(ctx context.Context, message *pb.QueryMessage) (*pb.WeatherReply, error) {
	dtstart, err := ptypes.Timestamp(message.DatetimeStart)
	if err != nil {
		log.Fatalln(err)
		return nil, errors.New("Invalid timestamp: datetime_start")
	}
	dtend, err := ptypes.Timestamp(message.DatetimeEnd)
	if err != nil {
		log.Fatalln(err)
		return nil, errors.New("Invalid timestamp: datetime_end")
	}
	query := &dao.Query{message.Location, &dtstart, &dtend}

	result, err := s.Database.Find(query)
	if err != nil {
		log.Fatalln(err)
		return nil, errors.New("Could not found record")
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
		log.Fatalln(err)
		return nil, errors.New("Invalid timestamp")
	}

	err = s.Database.Add(w)
	if err != nil {
		log.Fatalln(err)
		return nil, errors.New("Could not add record")
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
