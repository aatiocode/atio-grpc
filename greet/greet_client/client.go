package main

import (
	"context"
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

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Kimora",
			LastName:  "Alesha",
		},
	}
	// fmt.Printf("Create client: %f", c)
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
