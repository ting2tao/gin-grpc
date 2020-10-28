package main

import (
	"fmt"
	pb "gin-grpc/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//client := grpc.NewServer()

	//建立grpc连接
	client := pb.NewSimpleClient(conn)
	r := gin.Default()
	r.GET("/rest/n/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		fmt.Println(name)
		res, err := client.Route(ctx, &pb.SimpleRequest{Data: name})
		if err != nil {
			log.Fatalf("Call Route err: %v", err)
			return
		}
		fmt.Println(res.Code, res.Value, res.GetCode())

	})

	// Run http server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
