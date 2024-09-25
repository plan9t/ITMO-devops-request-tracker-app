package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"testing"
)

func TestNATSConnection(t *testing.T) {
	clusterID := "test-cluster"
	clientID := "GO_TEST-client"
	natsURL := "nats://localhost:4222"

	// Подключаемся к NATS Streaming
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		t.Fatalf("Failed to connect to NATS Streaming: %v", err)
	}
	defer sc.Close() // Закрываем соед

	// соединение установлено????
	if !sc.NatsConn().IsConnected() {
		t.Error("Expected to be connected to NATS Streaming, but it is not")
	} else {
		log.Println("Successfully connected to NATS Streaming")
	}
}
