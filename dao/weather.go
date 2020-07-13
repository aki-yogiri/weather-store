package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Weather struct {
	gorm.Model
	Id          int    `gorm:"primary_key;unique;not null"`
	Location    string `gorm:"size:255"`
	Weather     string `gorm:"type:varchar(100)"`
	Temperature float64
	Clouds      uint32
	Wind        float64
	WindDeg     uint32
	Timestamp   time.Time
}

type Query struct {
	Location      string
	DatetimeStart *time.Time
	DatetimeEnd   *time.Time
}

type WeatherRepository interface {
	Find(q *Query) ([]Weather, error)
	Add(w *Weather) error
	Connect() error
	Close()
}

type WeatherImplPostgres struct {
	db       *gorm.DB
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (wip *WeatherImplPostgres) Connect() error {
	connect := "host=" + wip.Host + " port=" + wip.Port + " user=" + wip.User + " dbname=" + wip.DBName + " password=" + wip.Password + " sslmode=disable"
	db, err := gorm.Open("postgres", connect)

	if err != nil {
		return err
	}
	wip.db = db

	return nil
}

func (wip *WeatherImplPostgres) Close() {
	wip.db.Close()
}

func (wip *WeatherImplPostgres) Find(q *Query) ([]Weather, error) {
	record := []Weather{}

	var err error

	if q.DatetimeStart != nil && q.DatetimeEnd != nil {
		err = wip.db.Where("location == ? AND timestamp BETWEEN ? AND ?", q.Location, q.DatetimeStart, q.DatetimeEnd).Find(&record).Error

	} else if q.DatetimeStart != nil {
		err = wip.db.Where("location == ? AND timestamp >=", q.Location, q.DatetimeStart).Find(&record).Error

	} else if q.DatetimeEnd != nil {
		err = wip.db.Where("location == ? AND timestamp <=", q.Location, q.DatetimeEnd).Find(&record).Error

	} else {
		err = wip.db.Where("location == ?", q.Location).Find(&record).Error
	}

	if err != nil {
		return nil, err
	}

	return record, nil
}

func (wip *WeatherImplPostgres) Add(w *Weather) error {
	wip.db.Begin()
	err := wip.db.Create(&w).Error
	if err != nil {
		wip.db.Rollback()
		return err
	}

	wip.db.Commit()

	return nil
}
