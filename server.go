package main

import (
	"log"
	"net"

	"github.com/rohitv5/grpc_golang_basic/chat"
	"google.golang.org/grpc"
)

//what is the difference between log and fmt?

func main(){
	lis,err := net.Listen("tcp",":9000")

	if(err != nil){
	 log.Fatalf("Failed to listen on port %v",err)	
	}

	s := chat.Server{}
	

	grpcServer := grpc.NewServer();

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Failed to start grpc on port 9000 %v",err)
	}


}