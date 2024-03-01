package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ctx := context.Background()
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
// 		fmt.Println("task is done")
// 	case <-ctx.Done():
// 		fmt.Println("task is canceleed")
// 	}
// }
