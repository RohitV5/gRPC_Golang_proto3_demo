package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {

}

//function 
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error){

	log.Printf("Recieved message body from client %v",message)
	return &Message{Body: "Hello from the server"}, nil



}