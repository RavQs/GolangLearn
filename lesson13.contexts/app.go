package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		err := CancelRequest(ctx)
		if err != nil {
			cancel()
		}
	}()

	DoRequest(ctx, "https://www.youtube.com/")
}

func CancelRequest(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return fmt.Errorf("fail request")
}

func DoRequest(ctx context.Context, requestStr string) {
	request, _ := http.NewRequest(http.MethodGet, requestStr, nil)
	request = request.WithContext(ctx)
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("response completed, status code = %d\n", res.StatusCode)
	case <-ctx.Done():
		fmt.Println("request takes too long")
	}

}
