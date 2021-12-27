package database

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	monitor "github.com/hypnoglow/go-pg-monitor"
	"github.com/hypnoglow/go-pg-monitor/gopgv10"
	"github.com/singhkshitij/golang-rest-service-starter/src/config"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
)

var db *pg.DB

func Setup() {
	dbConf := config.DbConfig()
	db = pg.Connect(&pg.Options{
		Addr: dbConf.Host + ":" + dbConf.Port,
		OnConnect: func(ctx context.Context, cn *pg.Conn) error {
			logger.Info("Connection to database established successfully")
			return nil
		},
		User:     dbConf.Username,
		Password: dbConf.Password,
		Database: dbConf.Name,
	})
	mon := monitor.NewMonitor(
		gopgv10.NewObserver(db),
		monitor.NewMetrics(
			monitor.MetricsWithNamespace("my_app"),
		),
	)

	mon.Open()
	defer mon.Close()
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		logger.Error("Failed to create tables", logger.KV("error", err))
	}
}

// createSchema creates database schema for User and other models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil), //add more tables here
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true, //change this to false for actually creating tables
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateUser() {
	user1 := &User{
		FirstName: "Lorem",
		LastName:  "Ipsum",
		Age:       30,
		Email:     "abc@xyz.com",
	}
	_, err := db.Model(user1).Insert()
	if err != nil {
		logger.Error("Failed to save user", logger.KV("error", err))
	}
}

func GetUser() {
	// Select user by primary key.
	user := &User{Email: "abc@xyz.com"}
	err := db.Model(user).WherePK().Select()
	if err != nil {
		logger.Error("Failed to get user", logger.KV("error", err))
	}
}
