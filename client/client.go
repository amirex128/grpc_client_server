package main

import (
	"context"
	"fmt"
	"github.com/amirex/grpc/github.com/amirex/microservice"
	"google.golang.org/grpc"
	"log"
)

func showErr(e ...interface{}) {
	if e[0] != nil {
		log.Fatalln(e)
	}
}
func main() {

	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	showErr(err, "grpc dail fail")

	defer dial.Close()
	client := microservice.NewRouteGuideClient(dial)

	details, err := client.GetPersonDetails(context.Background(), &microservice.RequestPerson{Id: 1})
	showErr(err, "response fail")

	fmt.Println(details)
}
