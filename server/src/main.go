package main

import(
	"fmt"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/ruanlas/poc-grpc-go-app-pbuser/user"
)

type server struct{
	pb.UnimplementedUserSenderServer
}

func (s *server) Send(ctx context.Context, in *pb.UserRequest) (*pb.MessageResponse, error) {
	log.Printf("Received: %v", in)
	fmt.Println("Received: ", in)
	return &pb.MessageResponse{Message: "O usuário : [" + fmt.Sprintf("%v",in) +"] foi recebido com sucesso!!"}, nil
}

func main(){
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		fmt.Println("failed to serve: ", err)
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	
	pb.RegisterUserSenderServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: ", err)
		log.Fatalf("failed to serve: %v", err)
	}
}