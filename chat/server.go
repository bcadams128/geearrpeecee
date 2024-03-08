package chat

import (
	"fmt"
	"geearrpeecee/pb"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
}

func NewServer() *ChatServer {
	return &ChatServer{}
}

func (c *ChatServer) SendMessage(msgStream pb.ChatService_SendMessageServer) error {
	msg, err := msgStream.Recv()
	ack := pb.MessageAck{Status: "SENT"}
	msgStream.SendAndClose(&ack)
	fmt.Println(msg, err)

	return nil
}
