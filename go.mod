module github.com/aki-yogiri/weather-store

go 1.14

require (
	github.com/aki-yogiri/weather-store/dao v0.0.0-00010101000000-000000000000
	github.com/aki-yogiri/weather-store/pb/weather v0.0.0-00010101000000-000000000000
	github.com/aki-yogiri/weather-store/service v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.14
	github.com/kelseyhightower/envconfig v1.4.0
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
)

replace (
	github.com/aki-yogiri/weather-store/dao v0.0.0-00010101000000-000000000000 => ./dao
	github.com/aki-yogiri/weather-store/pb/weather v0.0.0-00010101000000-000000000000 => ./pb/weather
	github.com/aki-yogiri/weather-store/service v0.0.0-00010101000000-000000000000 => ./service
)
