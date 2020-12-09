学习笔记

基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
    ﻿

### 说明

1. 开启 HTTP Server，提供 HTTP 服务
2. 启动失败，自动退出
3. 如果有收到 ctrl+c 信号时候，会停止 HTTP 服务

### 编译运行

```shell

go run main.go

// 2020/12/09 21:00:23 start http server


// ctrl + c

^C2020/12/09 21:00:36 receive close: interrupt
2020/12/09 21:00:36 stop http server success
2020/12/09 21:00:36 exit http server
2020/12/09 21:00:36 server exit

```

### Code


```go

// http service

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


// 启动流程

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


```