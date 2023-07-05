package main

import (
	"os"

	simplerestapi "github.com/2hard4me/simple-rest-api"
	"github.com/2hard4me/simple-rest-api/pkg/handler"
	"github.com/2hard4me/simple-rest-api/pkg/logging"
	"github.com/2hard4me/simple-rest-api/pkg/repository"
	"github.com/2hard4me/simple-rest-api/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	logger := logging.GetLogger()

	if err := initConfig(); err != nil {
		logger.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		logger.Fatalf("failed to initialize database: %s", err.Error() )
	}

	logger.Println("DB initializing -- Success")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(simplerestapi.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Fatalf("error occured while running http server: %s", err.Error())
	}
	
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}