package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "proxy_server/rpc/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

const (
	address = "localhost:50051"
)

type MessageRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (a *MessageRequest) Bind(r *http.Request) error {
	a.Title = "Hello?"
	a.Content = "Hello, Server!"
	return nil
}

func NewClient(address string) (pb.EchoMessageClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("did not connect: %v", err)
	}

	client := pb.NewEchoMessageClient(conn)

	return client, conn, nil
}

func echoMessage(ctx context.Context, c pb.EchoMessageClient, title string, content string) string {
	r, err := c.EchoMessage(ctx, &pb.MessageRequest{Title: title, Content: content})
	if err != nil {
		log.Fatalf("Could not send a message: %v", err)
	}
	log.Printf("Echo: %s", r.Title)
	return r.Title
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func main() {
	client, conn, err := NewClient(address)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer conn.Close()

  fmt.Println("gRPC client is connected to the server")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/echo", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		var message MessageRequest

		if err := render.DecodeJSON(r.Body, &message); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		fmt.Println("Title: ", message.Title)
		fmt.Println("Content: ", message.Content)

		response := echoMessage(ctx, client, message.Title, message.Content)

		w.Write([]byte(response))
	})

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
