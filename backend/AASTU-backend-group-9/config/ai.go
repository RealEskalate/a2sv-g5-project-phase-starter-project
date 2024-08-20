package config

import (
	"context"
	"log"
	"time"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)


func NewAIConfig(env *Env) *genai.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := genai.NewClient(ctx, option.WithAPIKey(env.AIAPIKey))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func CloseAIConnection(client *genai.Client) {
	if client == nil {
		return
	}
	client.Close()
	log.Println("Connection to AI closed.")
}