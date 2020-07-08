module github.com/aki-yogiri/weather-store

go 1.14

replace (
	github.com/aki-yogiri/weather-store/dao => ./dao
	github.com/aki-yogiri/weather-store/pb/weather => ./pb/weather
	github.com/aki-yogiri/weather-store/service => ./service
)

require (
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.14
	github.com/kelseyhightower/envconfig v1.4.0
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
)
