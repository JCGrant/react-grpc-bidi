package game

import (
	"math/rand"

	"github.com/JCGrant/react-grpc-bidi/server/protos"
)

var nextId = 0

var names = []string{
	"James",
	"Maya",
	"Colin",
	"Elizabeth",
}

func GenerateRandomPlayerUpdate() protos.PlayerUpdate {
	defer func() {
		nextId++
	}()
	return protos.PlayerUpdate{
		Id:   int64(nextId),
		Name: names[rand.Intn(len(names))],
		X:    int64(rand.Intn(800)),
		Y:    int64(rand.Intn(800)),
	}
}
