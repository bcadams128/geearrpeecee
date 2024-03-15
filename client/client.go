// package main

// import (
// 	"bufio"
// 	"context"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"os"

// 	"geearrpeecee/pb"

// 	"google.golang.org/grpc"
// )

// var senderName = flag.String("sender", "default", "Senders name")
// var tcpServer = flag.String("server", ":8080", "Tcp server")

// func sendMessage(ctx context.Context, client pb.ChatServiceClient, message string) {
// 	stream, err := client.SendMessage(ctx)
// 	if err != nil {
// 		log.Printf("Cannot send message: error: %v", err)
// 	}
// 	msg := pb.Message{
// 		Message: message,
// 		Sender:  *senderName,
// 	}
// 	stream.Send(&msg)

// 	ack, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Printf("Message sent: %v \n", ack)
// }

// func main() {

// 	flag.Parse()

// 	fmt.Println("--- CLIENT APP ---")
// 	var opts []grpc.DialOption
// 	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

// 	conn, err := grpc.Dial(*tcpServer, opts...)
// 	if err != nil {
// 		log.Fatalf("Fail to dail: %v", err)
// 	}

// 	defer conn.Close()

// 	ctx := context.Background()
// 	client := pb.NewChatServiceClient(conn)

// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {
// 		go sendMessage(ctx, client, scanner.Text())
// 	}

// }
