package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type keyDict string

const (
	contextKey keyDict = "first"
)

func main() {
	http.HandleFunc("/", handleEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handleEndpoint(w http.ResponseWriter, r *http.Request) {
	keyValue := r.Header.Get("key")
	ctx := context.WithValue(r.Context(), contextKey, keyValue)
	result := processingContext(ctx)
	w.Write([]byte(result))
}

func processingContext(ctx context.Context) string {
	select {
	case <-time.After(2 * time.Second):
		return fmt.Sprintf("task is done %d\n", ctx.Value(contextKey))
	case <-ctx.Done():
		log.Println("task is cancelled")
		return ""
	}
}
