package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"golang.org/x/sync/errgroup"
)
type MyServerHandler struct {
	name string
}

func (h *MyServerHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	_, _ = w.Write([]byte(h.name))
}

//可触发http.Server Close
type MyCloseHandler struct {
	Server *http.Server
}

func (h *MyCloseHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	_, _ = w.Write([]byte("closing"))
    fmt.Print("close server address ",h.Server)
	err := h.Server.Close()
	if err != nil {
		fmt.Printf("service close err %+v \n", err)
	}
}

func main(){

	// 监听系统信号
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	ctx := context.Background()
	group, _ := errgroup.WithContext(ctx)

	//启动两个http服务
	mux1 := http.NewServeMux()
	mux1.Handle("/", &MyServerHandler{name: "my service 1"})
	close1 :=MyCloseHandler{}
	mux1.Handle("/close", &close1)

	mux2 := http.NewServeMux()
	mux2.Handle("/", &MyServerHandler{name: "my service 2"})
	close2 :=MyCloseHandler{}
	mux2.Handle("/close", &close2)

	server1 := http.Server{Addr: "127.0.0.1:8080", Handler: mux1}
	server2 := http.Server{Addr: "127.0.0.1:8081", Handler: mux2}
	s1Ch := make(chan error)
	s2Ch := make(chan error)

	close1.Server = &server1
	close2.Server = &server2

	group.Go(func() error {
		err := server1.ListenAndServe()
		close(s1Ch)
		fmt.Printf("server1 closed %+v \n", err)
		return nil
	})

	group.Go(func() error {
		err := server2.ListenAndServe()
		close(s2Ch)
		fmt.Printf("server1 closed %+v \n", err)
		return nil
	})
	// 收到关闭信号或者有服务挂掉，关闭所有服务服务
	group.Go(func() error {
		select {
		case <-ch:
			fmt.Println("receive close signal!")
		case <-s1Ch:
			fmt.Println("receive server1 close!")
		case <-s2Ch:
			fmt.Println("receive server2 close!")
		}

		signal.Stop(ch)

		err1 := server1.Close()
		fmt.Printf("server1 close %+v \n", err1)

		err2 := server2.Close()
		fmt.Printf("server2 close %+v \n", err2)

		return nil
	})

	group.Wait()
	fmt.Println("game over")
}
