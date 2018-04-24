package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JCGrant/react-grpc-bidi/server/game"
	"github.com/JCGrant/react-grpc-bidi/server/protos"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type gameService struct{}

func newGameService() *gameService {
	return &gameService{}
}

func (s *gameService) StreamPlayerUpdates(req *protos.StreamPlayerUpdatesRequest, stream protos.GameService_StreamPlayerUpdatesServer) error {
	for {
		player := game.GenerateRandomPlayerUpdate()
		log.Printf("%+v", player)
		time.Sleep(time.Second / 10)
		stream.Send(&player)
	}
}

func main() {
	port := 9090

	grpcServer := grpc.NewServer()
	gameService := newGameService()
	protos.RegisterGameServiceServer(grpcServer, gameService)
	grpclog.SetLogger(log.New(os.Stdout, "game server: ", log.LstdFlags))

	websocketOriginFunc := grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
		return true
	})
	wrappedServer := grpcweb.WrapServer(grpcServer, grpcweb.WithWebsockets(true), websocketOriginFunc)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}

	grpclog.Printf("starting server at http://127.0.0.1:%d", port)

	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}
