// @Title
// @Description
// @Author smirror
// @Update 2022-09-23

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(
			func(w http.ResponseWriter, request *http.Request) {
				fmt.Fprintf(w, "Hello, %s!", request.URL.Path[1:])
			}),
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// http.ErrorServerClosedはhttp.ErrorServerClosedが正常に終了したことを示す
		if err := s.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close:%v", err)
		}
		return nil
	})

	// チャネルからの終了通知を待機する
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	// Goメソッドで起動した別ゴルーチンの終了を待つ。
	return eg.Wait()
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Printf("faile to terminate server: %v", err)
	}
}
