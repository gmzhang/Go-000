package main

import (
	"net/http"
	"fmt"
	"log"
	"context"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type httpHandler struct {
}

func (h *httpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World\n")
}

type HttpServer struct {
	http *http.Server
}

func NewHttpServer() *HttpServer {
	mux := http.NewServeMux()
	mux.Handle("/", &httpHandler{})

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return &HttpServer{
		http: httpServer,
	}
}

func (h *HttpServer) StartSvc() error {
	log.Println("start http server")
	if err := h.http.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	log.Println("exit http server")
	return nil
}

func (h *HttpServer) StopSvc(ctx context.Context) error {
	if err := h.http.Shutdown(ctx); err != nil {
		log.Fatal("stop http server fail")
		return err
	}
	log.Println("stop http server success")
	return nil
}

func main() {
	eg, ctx := errgroup.WithContext(context.Background())

	httpServer := NewHttpServer()

	eg.Go(func() error {
		return httpServer.StartSvc()
	})

	eg.Go(func() error {
		sign := make(chan os.Signal, 1)
		signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

		select {
		case sig := <-sign:
			log.Printf("receive close: %s", sig.String())

		case <-ctx.Done():
			log.Println("http start fail")
		}

		return httpServer.StopSvc(ctx)

	})

	if err := eg.Wait(); err != nil {
		log.Printf("server exit error: %+v", err)
		return
	}

	log.Println("server exit")
}
