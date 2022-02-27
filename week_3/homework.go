package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	ctxt, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctxt)

	s := http.Server{
		Addr: "127.0.0.1:8080",
	}
	group.Go(func() error {
		return serveApp(s)
	})

	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop")
		return s.Shutdown(errCtx)
	})

	sig := make(chan os.Signal, 10)

	signal.Notify(sig)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-sig:
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")
}

func serveApp(s http.Server) error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Hello, QCon!")
	})
	return s.ListenAndServe()
}
