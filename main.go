// @Title
// @Description
// @Author smirror
// @Update 2022-09-23

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"todo_app/config"

	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("server is running on %s", url)

	s := &http.Server{
		// 引数で受け取ったnet.Listnerをりようするので、
		// Addrフィールドは指定しない
		Handler: http.HandlerFunc(
			func(w http.ResponseWriter, request *http.Request) {
				fmt.Fprintf(w, "Hello, %s!", request.URL.Path[1:])
			}),
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.Serve(l); err != nil &&
			// http.ErrorServerClosedはhttp.ErrorServerClosedが正常に終了したことを示す
			err != http.ErrServerClosed {
			log.Printf("failed to close:%v", err)
			return err
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
		os.Exit(1)
	}
}
