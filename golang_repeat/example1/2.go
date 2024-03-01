package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// type key int

// const (
// 	userIdKey key = 0
// )

// func main() {
// 	ctx := context.WithValue(context.Background(), userIdKey, 1)
// 	ctx, cancel := context.WithCancel(ctx)
// 	go func() {
// 		fmt.Scanln()
// 		cancel()
// 	}()
// 	processContext(ctx)
// }

// func processContext(ctx context.Context) {
// 	select {
// 	case <-time.After(2 * time.Second):
// 		fmt.Println("task is done", ctx.Value(userIdKey))
// 	case <-ctx.Done():
// 		fmt.Println("task is canceleed")
// 	}
// }
