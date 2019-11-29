package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/muhammadhidayah/configuration-service/api/delivery/microgrpc"
	"github.com/muhammadhidayah/configuration-service/api/repository"
	"github.com/muhammadhidayah/configuration-service/api/usecase"
	pb "github.com/muhammadhidayah/configuration-service/proto/configuration"

	_ "github.com/lib/pq"
	"github.com/micro/go-micro"
)

func createConnection() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, DBName)

	return sql.Open("postgres", dbinfo)
}

func main() {
	db, err := createConnection()
	if err != nil {
		log.Fatalf(fmt.Sprintf("Could not connect to DB: %v", err))
	}

	defer db.Close()

	srv := micro.NewService(
		micro.Name("inact.srv.configuration"),
	)

	srv.Init()

	repo := repository.NewPgConfiguration(db)
	ucase := usecase.NewConfigurationUsecase(repo, time.Second*5)
	handler := microgrpc.NewMicroGrpc(ucase)
	pb.RegisterConfigurationServiceHandler(srv.Server(), handler)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
