package timeout

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for {
		select {
		case <- ctx.Done():
			log.Printf("timeout. finish job")
			return
		default:
			time.Sleep(1*time.Second)
			log.Print("a second elapsed")
		}
	}
}