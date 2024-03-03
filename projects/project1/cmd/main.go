package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kirigaikabuto/hero_preparation/projects/project1"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/pkg/handler"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/pkg/repository"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error install configuration file %s", err.Error())
		return
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error during reading from .env file %s", err.Error())
		return
	}
	db, err := repository.NewPostgres(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("db_password"),
		DBName:   viper.GetString("db.db_name"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	})
	if err != nil {
		logrus.Fatalf("error during connection to postgres %s", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		logrus.Fatalf("error during ping of postgres %s", err.Error())
		return
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srv := new(project1.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error running of server %s", err.Error())
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
