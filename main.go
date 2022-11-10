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

	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context, l net.Listener) error {
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
	if len(os.Args) != 2 {
		log.Println("need port number")
		os.Exit(1)
	}
	p := os.Args[1]
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", p, err)
	}
	if err := run(context.Background(), l); err != nil {
		fmt.Printf("faile to terminate server: %v", err)
		os.Exit(1)
	}
}
