package main

import (
	"fmt"
	"log"

	"github.com/aristio/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'ma client")
	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer con.Close()

	c := greetpb.NewGreetServiceClient(con)
	fmt.Printf("Create client: %f", c)
}
