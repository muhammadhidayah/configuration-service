package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/muhammadhidayah/configuration-service/proto/configuration"
)

func main() {
	srv := micro.NewService(
		micro.Name("inact.srv.configuration.cli"),
	)

	srv.Init()

	client := pb.NewConfigurationService("inact.srv.configuration", srv.Client())
	r, err := client.GetConfigurationClient(context.TODO(), &pb.RequestConfigCient{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)

}
