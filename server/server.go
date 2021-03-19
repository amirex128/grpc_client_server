package main

import (
	"context"
	"fmt"
	"github.com/amirex/grpc/github.com/amirex/microservice"
	"google.golang.org/grpc"
	"log"
	"net"
)

func showErr(e ...interface{}){
	if e[0] != nil {
		log.Fatalln(e)
	}
}
type Server struct{
	microservice.UnimplementedRouteGuideServer
}

func (r *Server) GetPersonDetails(ctx context.Context, req *microservice.RequestPerson) (*microservice.ResponsePerson, error) {
	fmt.Println(req.GetId())
	return &microservice.ResponsePerson{Person: &microservice.Person{Name: "amir shirdel"}},nil
}



func main() {

	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	showErr(err,"go server fail")

	grpcServer  := grpc.NewServer()

	microservice.RegisterRouteGuideServer(grpcServer , &Server{})

	fmt.Println("server is running...")

	showErr(grpcServer.Serve(listen),"failed to serve")

}
