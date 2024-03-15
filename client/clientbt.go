package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"geearrpeecee/pb"

	"google.golang.org/grpc"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

var senderName = flag.String("sender", "default", "Sender's name")
var tcpServer = flag.String("server", ":8080", "TCP server")

type chatModel struct {
	textarea textarea.Model
	client   pb.ChatServiceClient
	ctx      context.Context
}

func initialModel(client pb.ChatServiceClient, ctx context.Context) chatModel {
	ta := textarea.New()
	ta.Placeholder = "Type your message..."
	ta.Focus()
	ta.Prompt = "> "
	ta.CharLimit = 280
	ta.ShowLineNumbers = false

	return chatModel{
		textarea: ta,
		client:   client,
		ctx:      ctx,
	}
}

func (m chatModel) Init() tea.Cmd {
	return nil
}

func (m chatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			sendMessage(m.ctx, m.client, string(m.textarea.Value()))
			m.textarea.Reset()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m chatModel) View() string {
	return m.textarea.View()
}

func sendMessage(ctx context.Context, client pb.ChatServiceClient, message string) {
	stream, err := client.SendMessage(ctx)
	if err != nil {
		log.Printf("Cannot send message: error: %v", err)
	}
	msg := pb.Message{
		Message: message,
		Sender:  *senderName,
	}
	stream.Send(&msg)

	ack, err := stream.CloseAndRecv()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Message sent: %v \n", ack)
}

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

	conn, err := grpc.Dial(*tcpServer, opts...)
	if err != nil {
		log.Fatalf("Fail to dial: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	client := pb.NewChatServiceClient(conn)

	p := tea.NewProgram(initialModel(client, ctx))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
