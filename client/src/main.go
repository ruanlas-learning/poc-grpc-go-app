package main

import(
	"fmt"
	"os"
	"context"
	"log"
	"time"
	"strconv"

	"google.golang.org/grpc"
	pb "github.com/ruanlas/poc-grpc-go-app-pbuser/user"
)

func main(){

	argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
	for i, v := range argsWithoutProg {
		fmt.Println(i, v)
	}
/////////////////////////////////////////////////////////
	// Set up a connection to the server.
	conn, err := grpc.Dial("app_go_grpc_server:8089", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserSenderClient(conn)

	// Contact the server and print out its response.
	firstName := "defaultFirstName"
	lastName := "defaultLastName"
	var age uint32 = 0
	if len(argsWithoutProg) > 0 {
		firstName = argsWithoutProg[0]
	}
	if len(argsWithoutProg) > 1 {
		lastName = argsWithoutProg[1]
	}
	if len(argsWithoutProg) > 2 {
		ageu64, err := strconv.ParseUint(argsWithoutProg[2], 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		age = uint32(ageu64)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Send(ctx, &pb.UserRequest{FirstName: firstName, LastName: lastName, Age: age})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("MessageResponse: %s", r.GetMessage())
/////////////////////////////////////////////////////////
	fmt.Println()

}